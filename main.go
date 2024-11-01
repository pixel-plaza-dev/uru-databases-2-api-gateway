package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	apigrpc "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/listener"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/logger"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/user"
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/flag"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/grpc"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/listener"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// Load user service URI
	userUri, err := commongrpc.LoadUri(apigrpc.UserServiceHostKey, apigrpc.UserServicePortKey)
	if err != nil {
		panic(err)
	}
	logger.EnvironmentLogger.EnvironmentVariableLoaded(apigrpc.UserServiceHostKey, apigrpc.UserServicePortKey)

	// Connect to user service gRPC server
	conn, err := grpc.NewClient(userUri.Uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	// Create user client
	userClient := pbuser.NewUserClient(conn)

	// Gin router
	router := gin.Default()

	// Route group
	apiRoute := router.Group("/api/v1")

	// Create user controller
	userService := user.NewService(userClient)
	user.NewController(logger.UserServiceLogger, apiRoute, userService)

	// Run the server
	if err = router.Run(servicePort.FormattedPort); err != nil {
		panic(err)
	}
}
