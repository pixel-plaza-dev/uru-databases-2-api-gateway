package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	appgrpc "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc"
	appjwt "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/jwt"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/listener"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/logger"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/user"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonheader "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/security/header"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commonjwtvalidatorgrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator/grpc"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	clientauthinterceptor "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/interceptor/auth"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/listener"
	commontls "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/tls"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func init() {
	// Declare flags and parse them
	commonflag.SetModeFlag()
	flag.Parse()
	logger.FlagLogger.ModeFlagSet(commonflag.Mode)

	// Check if the environment is production
	if commonflag.Mode.IsProd() {
		return
	}

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		panic(commonenv.FailedToLoadEnvironmentVariablesError)
	}
}

func main() {
	// Get the listener port
	servicePort, err := commonlistener.LoadServicePort(
		"0.0.0.0", listener.PortKey,
	)
	if err != nil {
		panic(err)
	}
	logger.EnvironmentLogger.EnvironmentVariableLoaded(listener.PortKey)

	// Get the gRPC services URI
	var uris = make(map[string]string)
	for _, key := range []string{appgrpc.AuthServiceUriKey, appgrpc.UserServiceUriKey} {
		uri, err := commongrpc.LoadServiceURI(key)
		if err != nil {
			panic(err)
		}
		logger.EnvironmentLogger.EnvironmentVariableLoaded(key)
		uris[key] = uri
	}

	// Get the JWT public key
	jwtPublicKey, err := commonjwt.LoadJwtKey(appjwt.PublicKey)
	if err != nil {
		panic(err)
	}
	logger.EnvironmentLogger.EnvironmentVariableLoaded(appjwt.PublicKey)

	// Load Google Cloud service account credentials
	googleCredentials, err := commongcloud.LoadGoogleCredentials(context.Background())
	if err != nil {
		panic(err)
	}

	// Get the service account token source for each gRPC server URI
	var tokenSources = make(map[string]*oauth.TokenSource)
	for key, uri := range uris {
		tokenSource, err := commongcloud.LoadServiceAccountCredentials(
			context.Background(), "https://"+uri, googleCredentials,
		)
		if err != nil {
			panic(err)
		}
		tokenSources[key] = tokenSource
		// logger.GCloudLogger.LoadedTokenSource(tokenSource)
	}

	// Load transport credentials
	var transportCredentials credentials.TransportCredentials

	if commonflag.Mode.IsDev() {
		// Load server TLS credentials
		transportCredentials, err = credentials.NewServerTLSFromFile(
			appgrpc.ServerCertPath, appgrpc.ServerKeyPath,
		)
		if err != nil {
			panic(err)
		}
	} else {
		// Load system certificates pool
		transportCredentials, err = commontls.LoadSystemCredentials()
		if err != nil {
			panic(err)
		}
	}

	// Create client authentication interceptors
	var clientAuthInterceptors = make(map[string]*clientauthinterceptor.Interceptor)
	for key, tokenSource := range tokenSources {
		clientAuthInterceptor, err := clientauthinterceptor.NewInterceptor(tokenSource)
		if err != nil {
			panic(err)
		}
		clientAuthInterceptors[key] = clientAuthInterceptor
	}

	// Create gRPC connections
	var conns = make(map[string]*grpc.ClientConn)
	for key, uri := range uris {
		conn, err := grpc.NewClient(
			uri, grpc.WithTransportCredentials(transportCredentials),
			grpc.WithChainUnaryInterceptor(clientAuthInterceptors[key].Authenticate()),
		)
		if err != nil {
			panic(err)
		}
		conns[key] = conn
	}
	defer func(conns map[string]*grpc.ClientConn) {
		for _, conn := range conns {
			err = conn.Close()
			if err != nil {
				panic(err)
			}
		}
	}(conns)

	// Create gRPC server clients
	userClient := pbuser.NewUserClient(conns[appgrpc.UserServiceUriKey])
	authClient := pbauth.NewAuthClient(conns[appgrpc.AuthServiceUriKey])

	// Create token validator
	tokenValidator, err := commonjwtvalidatorgrpc.NewDefaultTokenValidator(
		tokenSources[appgrpc.AuthServiceUriKey], &authClient)
	if err != nil {
		panic(err)
	}

	// Create JWT validator
	jwtValidator, err := commonjwtvalidator.NewDefaultValidator(
		[]byte(jwtPublicKey),
		tokenValidator)
	if err != nil {
		panic(err)
	}

	// Create JWT middleware
	jwtMiddleware := authmiddleware.NewMiddleware(jwtValidator)

	// Check if the mode is production
	if commonflag.Mode.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Gin router
	router := gin.Default()

	// Added secure headers middleware
	router.Use(commonheader.SecurityHeaders())

	// Route group
	apiRoute := router.Group("/api/v1")

	// Create user controller
	userService := user.NewService(commonflag.Mode, userClient)
	user.NewController(apiRoute, userService, jwtMiddleware)

	// Create auth controller
	authService := auth.NewService(commonflag.Mode, authClient)
	auth.NewController(apiRoute, authService, jwtMiddleware)

	// Run the server
	if err = router.Run(servicePort.FormattedPort); err != nil {
		panic(err)
	}
}
