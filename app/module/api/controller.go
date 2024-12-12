package api

import (
	"github.com/gin-gonic/gin"
	modulev1 "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbconfigrestapi "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api"
)

// Controller struct for the API module
// @Summary API Router Group
// @Description Router group for API-related endpoints
// @Tags api
// @Accept json
// @Produce json
// @Router /api [group]
type Controller struct {
	engine          *gin.Engine
	route           *gin.RouterGroup
	authMiddleware  authmiddleware.Authentication
	responseHandler commonclientresponse.Handler
	v1Controller    *modulev1.Controller
}

// NewController creates a new controller
func NewController(
	engine *gin.Engine,
	authMiddleware authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the API controller
	route := engine.Group(pbconfigrestapi.Base.String())

	// Create a new  controller
	return &Controller{
		engine:          engine,
		route:           route,
		authMiddleware:  authMiddleware,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the API version 1 controller
func (c *Controller) Initialize() {}

// InitializeV1 initializes the routes for the API version 1 controller
func (c *Controller) InitializeV1() *modulev1.Controller {
	// Check if the API version 1 controller has already been initialized
	if c.v1Controller != nil {
		return c.v1Controller
	}

	// Initialize the API version 1 controller
	v1Controller := modulev1.NewController(c.route, c.authMiddleware, c.responseHandler)
	v1Controller.Initialize()

	// Store the API version 1 controller
	c.v1Controller = v1Controller

	return v1Controller
}
