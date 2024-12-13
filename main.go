package main

import (
	"context"
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app"
	appgrpc "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc"
	appjwt "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/jwt"
	applistener "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/listener"
	applogger "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/logger"
	appapi "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/docs"
	commonginmiddlewareauth "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonheader "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/security/header"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commonjwtvalidatorgrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator/grpc"
	clientauthinterceptor "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/interceptor/auth"
	commongrpcoutgoingctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/interceptor/outgoing-ctx"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/listener"
	commontls "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/tls"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pborder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/order"
	pbpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/payment"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	pbconfigorder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/order"
	pbconfigpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/payment"
	pbconfigshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	pbconfiguser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	pbtypesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func init() {
	// Declare flags and parse them
	commonflag.SetModeFlag()
	flag.Parse()
	applogger.FlagLogger.ModeFlagSet(commonflag.Mode)

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

// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Get the listener port
	servicePort, err := commonlistener.LoadServicePort(
		"0.0.0.0", applistener.PortKey,
	)
	if err != nil {
		panic(err)
	}
	applogger.EnvironmentLogger.EnvironmentVariableLoaded(applistener.PortKey)

	// Dynamically set the Swagger host
	docs.SwaggerInfo.Host = "localhost:" + servicePort.Port

	// Get the gRPC services URI
	var uriKeys = []string{
		appgrpc.AuthServiceUriKey,
		appgrpc.UserServiceUriKey,
		appgrpc.ShopServiceUriKey,
		appgrpc.PaymentServiceUriKey,
		appgrpc.OrderServiceUriKey,
	}
	var uris = make(map[string]string)
	for _, uriKey := range uriKeys {
		uri, err := commonenv.LoadVariable(uriKey)
		if err != nil {
			panic(err)
		}
		applogger.EnvironmentLogger.EnvironmentVariableLoaded(uriKey)
		uris[uriKey] = uri
	}

	// Get the JWT public key
	jwtPublicKey, err := commonenv.LoadVariable(appjwt.PublicKey)
	if err != nil {
		panic(err)
	}
	applogger.EnvironmentLogger.EnvironmentVariableLoaded(appjwt.PublicKey)

	// Load Google Cloud service account credentials
	googleCredentials, err := commongcloud.LoadGoogleCredentials(context.Background())
	if err != nil {
		panic(err)
	}

	// Get the service account token source for each gRPC server URI
	var tokenSources = make(map[string]*oauth.TokenSource)
	for _, uriKey := range uriKeys {
		tokenSource, err := commongcloud.LoadServiceAccountCredentials(
			context.Background(), "https://"+uris[uriKey], googleCredentials,
		)
		if err != nil {
			panic(err)
		}
		tokenSources[uriKey] = tokenSource
		// applogger.GCloudLogger.LoadedTokenSource(tokenSource)
	}

	// Load transport credentials
	var transportCredentials credentials.TransportCredentials

	if commonflag.Mode.IsDev() {
		// Load server TLS credentials
		transportCredentials, err = credentials.NewServerTLSFromFile(
			app.ServerCertPath, app.ServerKeyPath,
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

	// Create gRPC interceptions map
	var grpcInterceptions = map[string]*map[pbtypesgrpc.Method]pbtypesgrpc.Interception{
		appgrpc.UserServiceUriKey:    &pbconfiguser.Interceptions,
		appgrpc.AuthServiceUriKey:    &pbconfigauth.Interceptions,
		appgrpc.ShopServiceUriKey:    &pbconfigshop.Interceptions,
		appgrpc.OrderServiceUriKey:   &pbconfigorder.Interceptions,
		appgrpc.PaymentServiceUriKey: &pbconfigpayment.Interceptions,
	}

	// Create client authentication interceptors
	var clientAuthInterceptors = make(map[string]*clientauthinterceptor.Interceptor)
	for _, uriKey := range uriKeys {
		clientAuthInterceptor, err := clientauthinterceptor.NewInterceptor(
			tokenSources[uriKey],
			grpcInterceptions[uriKey],
		)
		if err != nil {
			panic(err)
		}
		clientAuthInterceptors[uriKey] = clientAuthInterceptor
	}

	// Create common gRPC client interceptors after authentication
	var commonInterceptorsAfterAuth []grpc.UnaryClientInterceptor
	if commonflag.Mode.IsDev() {
		outgoingCtx, err := commongrpcoutgoingctx.NewInterceptor(applogger.OutgoingCtxLogger)
		if err != nil {
			panic(err)
		}

		// Add the outgoing context interceptor
		commonInterceptorsAfterAuth = append(commonInterceptorsAfterAuth, outgoingCtx.PrintOutgoingCtx())
	}

	// Create gRPC connections
	var conns = make(map[string]*grpc.ClientConn)
	for _, uriKey := range uriKeys {
		conn, err := grpc.NewClient(
			uris[uriKey], grpc.WithTransportCredentials(transportCredentials),
			grpc.WithChainUnaryInterceptor(
				append(
					[]grpc.UnaryClientInterceptor{
						clientAuthInterceptors[uriKey].Authenticate(),
					}, commonInterceptorsAfterAuth...,
				)...,
			),
		)
		if err != nil {
			panic(err)
		}
		conns[uriKey] = conn
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
	shopClient := pbshop.NewShopClient(conns[appgrpc.ShopServiceUriKey])
	paymentClient := pbpayment.NewPaymentClient(conns[appgrpc.PaymentServiceUriKey])
	orderClient := pborder.NewOrderClient(conns[appgrpc.OrderServiceUriKey])

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

	// Create the response handler
	responseHandler, err := commonclientresponse.NewDefaultHandler(commonflag.Mode)
	if err != nil {
		panic(err)
	}

	// Create the authentication middleware
	authMiddleware, err := commonginmiddlewareauth.NewMiddleware(
		jwtValidator,
		applogger.AuthMiddlewareLogger,
		responseHandler,
	)
	if err != nil {
		panic(err)
	}

	// Gin router
	router := gin.Default()

	// Set up CORS middleware
	router.Use(cors.Default())

	// Added secure headers middleware
	router.Use(commonheader.SecurityHeaders())

	// Use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Create the API controller
	mainController := appapi.NewController(
		router, authMiddleware, responseHandler,
	)

	// Initialize the API version 1 controller
	v1Controller := mainController.InitializeV1()

	// Initialize the API version 1 children controllers
	v1Controller.InitializeAuth(authClient)
	v1Controller.InitializeUsers(userClient)
	v1Controller.InitializePayments(paymentClient)
	v1Controller.InitializeShops(shopClient)
	v1Controller.InitializeOrders(orderClient)

	// Run the server
	if err = router.Run(servicePort.FormattedPort); err != nil {
		panic(err)
	}
	applogger.ListenerLogger.ServerStarted(servicePort.Port)

	/*
		// Run the server
		if commonflag.Mode.IsProd() {
			if err = router.Run(servicePort.FormattedPort); err != nil {
				panic(err)
			}
			applogger.ListenerLogger.ServerStarted(servicePort.Port)
		} else {
			// Start the server with HTTPS
			if err = router.RunTLS(servicePort.FormattedPort, app.ServerCertPath, app.ServerKeyPath); err != nil {
				panic(err)
			}
		}
	*/
}
