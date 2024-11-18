package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	appgrpc "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc"
	appjwt "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/jwt"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/listener"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/logger"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/auth"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/user"
	jwtmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/jwt"
	commonheader "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/security/header"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc"
	clientauth "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc/client/interceptor/auth"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/listener"
	commontls "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/tls"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/auth"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
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

	// Get the API gateway URI
	apiUri, err := commongrpc.LoadServiceURI(appgrpc.ApiGatewayUriKey)
	if err != nil {
		panic(err)
	}
	logger.EnvironmentLogger.EnvironmentVariableLoaded(appgrpc.ApiGatewayUriKey)

	// Get the auth service URI
	authUri, err := commongrpc.LoadServiceURI(appgrpc.AuthServiceUriKey)
	if err != nil {
		panic(err)
	}
	logger.EnvironmentLogger.EnvironmentVariableLoaded(appgrpc.AuthServiceUriKey)

	// Get the user service URI
	userUri, err := commongrpc.LoadServiceURI(appgrpc.UserServiceUriKey)
	if err != nil {
		panic(err)
	}
	logger.EnvironmentLogger.EnvironmentVariableLoaded(appgrpc.UserServiceUriKey)

	// Get the JWT public key
	jwtPublicKey, err := commonjwt.LoadJwtKey(appjwt.PublicKey)
	if err != nil {
		panic(err)
	}
	logger.EnvironmentLogger.EnvironmentVariableLoaded(appjwt.PublicKey)

	// gRPC servers URI
	var uris = []string{authUri, userUri}

	// Get the API gateway account token source
	tokenSource, err := commongcloud.LoadServiceAccountCredentials(
		context.Background(), "https://"+apiUri,
	)
	if err != nil {
		panic(err)
	}

	// Load transport credentials
	var transportCredentials credentials.TransportCredentials

	if commonflag.Mode.IsDev() {
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
	clientAuthInterceptor, err := clientauth.NewInterceptor(tokenSource)
	if err != nil {
		panic(err)
	}

	// Create gRPC connections
	var conns = make(map[string]*grpc.ClientConn)
	for _, uri := range uris {
		conn, err := grpc.NewClient(
			uri, grpc.WithTransportCredentials(transportCredentials),
			grpc.WithChainUnaryInterceptor(clientAuthInterceptor.Authenticate()),
		)
		if err != nil {
			panic(err)
		}
		conns[uri] = conn
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
	userClient := pbuser.NewUserClient(conns[userUri])
	authClient := pbauth.NewAuthClient(conns[authUri])

	// Create JWT validator
	jwtValidator, err := commonjwtvalidator.NewDefaultValidator(
		[]byte(jwtPublicKey),
		func(claims *jwt.MapClaims) (*jwt.MapClaims, error) {
			// Get the expiration time
			exp, err := claims.GetExpirationTime()
			if err != nil {
				return nil, commonjwt.InvalidClaimsError
			}

			// Check if the token is expired
			if exp.Before(time.Now()) {
				return nil, commonjwt.TokenExpiredError
			}
			return claims, nil
		},
	)
	if err != nil {
		panic(err)
	}

	// Create JWT middleware
	jwtMiddleware := jwtmiddleware.NewMiddleware(jwtValidator)

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
