package user

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	apiRoute *gin.RouterGroup
	route    *gin.RouterGroup
	logger   Logger
	service  *Service
}

// NewController creates a new controller
func NewController(logger Logger, apiRoute *gin.RouterGroup, service *Service) *Controller {
	// Create a new route for the user service
	userRoute := apiRoute.Group("/user")

	// Create a new user controller
	instance := &Controller{apiRoute: apiRoute, route: userRoute, service: service, logger: logger}

	// Initialize the routes for the controller
	instance.initializeRoutes()

	return instance
}

// initializeRoutes initializes the routes for the controller
func (c *Controller) initializeRoutes() {
	c.route.POST("/sign-up", c.signUp)
}

// signUp creates a new user
func (c *Controller) signUp(context *gin.Context) {
	// Log the received sign up request
	// c.logger.ReceivedSignUpRequest()

	// Create a new user
	response, err := c.service.SignUp(context)
	if err != nil {
		// Log the sign-up failure and return the error
		c.logger.SignUpFailed(err)
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Log the sign-up success and return the response
	c.logger.SignUpSuccess()
	context.JSON(201, gin.H{"message": response})
}
