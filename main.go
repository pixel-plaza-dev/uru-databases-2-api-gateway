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
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api"
	commonginmiddlewareauth "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonheader "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/security/header"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commonjwtvalidatorgrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator/grpc"
	clientauthinterceptor "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/interceptor/auth"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/listener"
	commontls "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/tls"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	_ "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/docs"
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

// @Title           Pixel Plaza REST API
// @Version         1.0
// @Description     The REST API Gateway to the Auth, User, Business, Payment and Order microservices

// @License.name  GPL-3.0
// @License.url   http://www.gnu.org/licenses/gpl-3.0.html

// @Host      uru-databases-2-api-gateway-246064477369.us-central1.run.app
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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
	for _, key := range []string{
		appgrpc.AuthServiceUriKey,
		appgrpc.UserServiceUriKey,
	} {
		uri, err := commonenv.LoadVariable(key)
		if err != nil {
			panic(err)
		}
		logger.EnvironmentLogger.EnvironmentVariableLoaded(key)
		uris[key] = uri
	}

	// Get the JWT public key
	jwtPublicKey, err := commonenv.LoadVariable(appjwt.PublicKey)
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
		tokenSources[appgrpc.AuthServiceUriKey], authClient, nil,
	)
	if err != nil {
		panic(err)
	}

	// Create JWT validator with ED25519 public key
	jwtValidator, err := commonjwtvalidator.NewEd25519Validator(
		[]byte(jwtPublicKey),
		tokenValidator,
		commonflag.Mode,
	)
	if err != nil {
		panic(err)
	}

	// Check if the mode is production
	if commonflag.Mode.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create the authentication middleware
	authMiddleware, err := commonginmiddlewareauth.NewMiddleware(
		jwtValidator,
		logger.AuthMiddlewareLogger,
		commonflag.Mode,
	)
	if err != nil {
		panic(err)
	}

	// Gin router
	router := gin.Default()

	// Added secure headers middleware
	router.Use(commonheader.SecurityHeaders())

	// Use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Create the response handler
	responseHandler, err := commonclientresponse.NewDefaultHandler(commonflag.Mode)
	if err != nil {
		panic(err)
	}

	// Create the module controller
	mainModule := api.NewController(
		router, userClient, authClient, authMiddleware, responseHandler,
	)

	// Initialize the module controllers
	mainModule.Initialize()

	// Run the server
	if err = router.Run(servicePort.FormattedPort); err != nil {
		panic(err)
	}
	logger.ListenerLogger.ServerStarted(servicePort.Port)
}
