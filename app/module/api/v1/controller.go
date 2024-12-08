package v1

import (
	"github.com/gin-gonic/gin"
	moduleauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth"
	moduleusers "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/users"
	apptypescontroller "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types/controller"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigrestv1 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1"
)

// Controller struct
// @Summary API Version 1 Router Group
// @Description Router group for API version 1-related endpoints
// @Tags v1
// @Accept json
// @Produce json
// @Router /api/v1 [group]
type Controller struct {
	route           *gin.RouterGroup
	userClient      pbuser.UserClient
	authClient      pbauth.AuthClient
	authMiddleware  authmiddleware.Authentication
	responseHandler commonclientresponse.Handler
}

// NewController creates a new controller
func NewController(
	baseRoute *gin.RouterGroup,
	userClient pbuser.UserClient,
	authClient pbauth.AuthClient,
	authMiddleware authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the API version 1 controller
	route := baseRoute.Group(pbconfigrestv1.Base.String())

	// Create a new  controller
	return &Controller{
		route:           route,
		userClient:      userClient,
		authClient:      authClient,
		authMiddleware:  authMiddleware,
		responseHandler: responseHandler,
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
	userController := moduleusers.NewController(c.route, c.userClient, c.authMiddleware, c.responseHandler)
	authController := moduleauth.NewController(c.route, c.authClient, c.authMiddleware, c.responseHandler)

	// Initialize the children controllers routes
	for _, controller := range []apptypescontroller.Controller{
		userController,
		authController,
	} {
		controller.Initialize()
	}
}
