package api

import (
	"github.com/gin-gonic/gin"
	modulev1 "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1"
	apptypescontroller "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types/controller"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigrestapi "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api"
)

// Controller struct
// @Summary API Router Group
// @Description Router group for API-related endpoints
// @Tags api
// @Accept json
// @Produce json
// @Router /api [group]
type Controller struct {
	engine          *gin.Engine
	route           *gin.RouterGroup
	userClient      pbuser.UserClient
	authClient      pbauth.AuthClient
	authMiddleware  authmiddleware.Authentication
	responseHandler commonclientresponse.Handler
}

// NewController creates a new controller
func NewController(
	engine *gin.Engine,
	userClient pbuser.UserClient,
	authClient pbauth.AuthClient,
	authMiddleware authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the API controller
	route := engine.Group(pbconfigrestapi.Base.String())

	// Create a new  controller
	return &Controller{
		engine:          engine,
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
	v1Controller := modulev1.NewController(c.route, c.userClient, c.authClient, c.authMiddleware, c.responseHandler)

	// Initialize the children controllers routes
	for _, controller := range []apptypescontroller.Controller{
		v1Controller,
	} {
		controller.Initialize()
	}
}
