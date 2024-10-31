package user

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	apiRoute *gin.RouterGroup
	route    *gin.RouterGroup
	service  *Service
}

// NewController creates a new controller
func NewController(apiRoute *gin.RouterGroup, service *Service) *Controller {
	// Create a new route for the user service
	userRoute := apiRoute.Group("/user")

	// Create a new user controller
	instance := &Controller{apiRoute: apiRoute, route: userRoute, service: service}

	// Initialize the routes for the controller
	instance.initializeRoutes(userRoute)

	return instance
}

// initializeRoutes initializes the routes for the controller
func (c *Controller) initializeRoutes(route *gin.RouterGroup) {
	c.route = route
	c.route.POST("/sign-up", c.signUp)
}

// signUp creates a new user
func (c *Controller) signUp(context *gin.Context) {
	// Create a new user
	response, err := c.service.SignUp(context)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	context.JSON(201, gin.H{"message": response})
}
