package main

import (
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
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/gin/middleware/auth"
	commonheader "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/gin/middleware/security/header"
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/flag"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/grpc"
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt/validator"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/listener"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/auth"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/user"
	"google.golang.org/grpc"
	"time"
)

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		panic(commonenv.FailedToLoadEnvironmentVariablesError)
	}
}

func main() {
	// Declare flags and parse them
	commonflag.SetModeFlag()
	flag.Parse()
	logger.FlagLogger.ModeFlagSet(commonflag.Mode)

	// Get the listener port
	servicePort, err := commonlistener.LoadServicePort(listener.ApiGatewayPortKey)
	if err != nil {
		panic(err)
	}
	logger.EnvironmentLogger.EnvironmentVariableLoaded(listener.ApiGatewayPortKey)

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

	// Load the CA certificate for the Pixel Plaza's services
	CACredentials, err := commongrpc.LoadTLSCredentials(appgrpc.CACertificatePath)
	if err != nil {
		panic(err)
	}

	// Connect to user service gRPC server
	userConn, err := grpc.NewClient(userUri, grpc.WithTransportCredentials(CACredentials))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}(userConn)

	// Connect to auth service gRPC server
	authConn, err := grpc.NewClient(authUri, grpc.WithTransportCredentials(CACredentials))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}(authConn)

	// Create user client
	userClient := pbuser.NewUserClient(userConn)

	// Create auth client
	authClient := pbauth.NewAuthClient(authConn)

	// Read the JWT public key
	jwtFile, err := commonjwt.ReadJwtKey(appjwt.PublicKeyPath)
	if err != nil {
		panic(err)
	}

	// Create JWT validator
	jwtValidator, err := commonjwtvalidator.NewDefaultValidator(jwtFile, func(claims *jwt.MapClaims) (*jwt.MapClaims, error) {
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
	})
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
