package v1

import (
	"github.com/gin-gonic/gin"
	moduleauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/v1/auth"
	moduleusers "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/v1/users"
	apptypescontroller "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types/controller"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonclientrequest "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/request"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/user"
	pbconfigrestv1 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1"
)

// Controller struct
// @Summary API Version 1 Router Group
// @Description Router group for API version 1-related endpoints
// @Tags v1
// @Accept json
// @Produce json
// @Router /api/v1 [group]
type Controller struct {
	engine         *gin.Engine
	route          *gin.RouterGroup
	userClient     pbuser.UserClient
	authClient     pbauth.AuthClient
	authMiddleware authmiddleware.Authentication
	requestHandler commonclientrequest.Handler
}

// NewController creates a new controller
func NewController(
	engine *gin.Engine,
	userClient pbuser.UserClient,
	authClient pbauth.AuthClient,
	authMiddleware authmiddleware.Authentication,
	requestHandler commonclientrequest.Handler,
) *Controller {
	// Create a new route
	route := engine.Group(pbconfigrestv1.BaseURI)

	// Create a new  controller
	return &Controller{
		engine:         engine,
		route:          route,
		userClient:     userClient,
		authClient:     authClient,
		authMiddleware: authMiddleware,
		requestHandler: requestHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	userController := moduleusers.NewController(c.route, c.userClient, c.authMiddleware, c.requestHandler)
	authController := moduleauth.NewController(c.route, c.authClient, c.authMiddleware, c.requestHandler)

	// Initialize the children controllers routes
	for _, controller := range []apptypescontroller.Controller{
		userController,
		authController,
	} {
		controller.Initialize()
	}
}
