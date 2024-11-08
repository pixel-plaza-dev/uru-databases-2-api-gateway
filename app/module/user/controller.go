package user

import (
	"github.com/gin-gonic/gin"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/gin/middleware/auth"
)

type Controller struct {
	apiRoute      *gin.RouterGroup
	route         *gin.RouterGroup
	service       *Service
	jwtMiddleware authmiddleware.Authentication
}

// NewController creates a new controller
func NewController(apiRoute *gin.RouterGroup, service *Service, jwtMiddleware authmiddleware.Authentication) *Controller {
	// Create a new route for the user service
	route := apiRoute.Group("/user")

	// Create a new user controller
	instance := &Controller{apiRoute: apiRoute, route: route, service: service, jwtMiddleware: jwtMiddleware}

	// Initialize the routes for the controller
	instance.initializeRoutes()

	return instance
}

// initializeRoutes initializes the routes for the controller
func (c *Controller) initializeRoutes() {
	c.route.POST("/sign-up", c.signUp)
	c.route.PATCH("/profile", c.jwtMiddleware.Authenticate(), c.updateProfile)
	c.route.GET("/profile", c.jwtMiddleware.Authenticate(), c.getProfile)
	c.route.GET("/full-profile", c.jwtMiddleware.Authenticate(), c.getFullProfile)
	c.route.PATCH("/password", c.jwtMiddleware.Authenticate(), c.changePassword)
	c.route.PATCH("/username", c.jwtMiddleware.Authenticate(), c.changeUsername)
	c.route.PATCH("/email", c.jwtMiddleware.Authenticate(), c.changePrimaryEmail)
	c.route.POST("/email", c.jwtMiddleware.Authenticate(), c.addEmail)
	c.route.DELETE("/email", c.jwtMiddleware.Authenticate(), c.deleteEmail)
	c.route.GET("/email", c.jwtMiddleware.Authenticate(), c.getPrimaryEmail)
	c.route.GET("/emails", c.jwtMiddleware.Authenticate(), c.getActiveEmails)
	c.route.POST("/send-verification-email", c.jwtMiddleware.Authenticate(), c.sendVerificationEmail)
	c.route.POST("/verify-email", c.jwtMiddleware.Authenticate(), c.verifyEmail)
	c.route.PATCH("/phone-number", c.jwtMiddleware.Authenticate(), c.changePhoneNumber)
	c.route.GET("/phone-number", c.jwtMiddleware.Authenticate(), c.getPhoneNumber)
	c.route.POST("/send-verification-phone-number", c.jwtMiddleware.Authenticate(), c.sendVerificationPhoneNumber)
	c.route.POST("/verify-phone-number", c.jwtMiddleware.Authenticate(), c.verifyPhoneNumber)
	c.route.POST("/forgot-password", c.forgotPassword)
	c.route.POST("/reset-password", c.resetPassword)
	c.route.DELETE("/close-account", c.jwtMiddleware.Authenticate(), c.deleteUser)
}

// signUp creates a new user
func (c *Controller) signUp(ctx *gin.Context) {
	// Create a new user
	response, err := c.service.SignUp(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"message": response})
}

// updateProfile updates the user's profile
func (c *Controller) updateProfile(ctx *gin.Context) {
	// Update the user's profile
	response, err := c.service.UpdateProfile(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// getProfile gets the user's profile
func (c *Controller) getProfile(ctx *gin.Context) {
	// Get the user's profile
	response, err := c.service.GetProfile(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// getFullProfile gets the user's full profile
func (c *Controller) getFullProfile(ctx *gin.Context) {
	// Get the user's full profile
	response, err := c.service.GetFullProfile(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, response)
}

// changePassword changes the user's password
func (c *Controller) changePassword(ctx *gin.Context) {
	// Change the user's password
	response, err := c.service.ChangePassword(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// changeUsername changes the user's username
func (c *Controller) changeUsername(ctx *gin.Context) {
	// Change the user's username
	response, err := c.service.ChangeUsername(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// changePrimaryEmail changes the user's primary email
func (c *Controller) changePrimaryEmail(ctx *gin.Context) {
	// Change the user's primary email
	response, err := c.service.ChangePrimaryEmail(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// addEmail adds an email to the user's account
func (c *Controller) addEmail(ctx *gin.Context) {
	// Add an email to the user's account
	response, err := c.service.AddEmail(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"message": response})
}

// deleteEmail deletes an email from the user's account
func (c *Controller) deleteEmail(ctx *gin.Context) {
	// Delete an email from the user's account
	response, err := c.service.DeleteEmail(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": response})
}

// getPrimaryEmail gets the user's primary email
func (c *Controller) getPrimaryEmail(ctx *gin.Context) {
	// Get the user's primary email
	response, err := c.service.GetPrimaryEmail(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// getActiveEmails gets the user's active emails
func (c *Controller) getActiveEmails(ctx *gin.Context) {
	// Get the user's active emails
	response, err := c.service.GetActiveEmails(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// sendVerificationEmail sends a verification email
func (c *Controller) sendVerificationEmail(ctx *gin.Context) {
	// Send a verification email
	response, err := c.service.SendVerificationEmail(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// verifyEmail verifies the user's email
func (c *Controller) verifyEmail(ctx *gin.Context) {
	// Verify the user's email
	response, err := c.service.VerifyEmail(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// changePhoneNumber changes the user's phone number
func (c *Controller) changePhoneNumber(ctx *gin.Context) {
	// Change the user's phone number
	response, err := c.service.ChangePhoneNumber(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// getPhoneNumber gets the user's phone number
func (c *Controller) getPhoneNumber(ctx *gin.Context) {
	// Get the user's active phone numbers
	response, err := c.service.GetPhoneNumber(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// sendVerificationPhoneNumber sends a verification phone number
func (c *Controller) sendVerificationPhoneNumber(ctx *gin.Context) {
	// Send a verification phone number
	response, err := c.service.SendVerificationPhoneNumber(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// verifyPhoneNumber verifies the user's phone number
func (c *Controller) verifyPhoneNumber(ctx *gin.Context) {
	// Verify the user's phone number
	response, err := c.service.VerifyPhoneNumber(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// forgotPassword sends a reset password email
func (c *Controller) forgotPassword(ctx *gin.Context) {
	// Send a reset password email
	response, err := c.service.ForgotPassword(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// resetPassword resets the user's password
func (c *Controller) resetPassword(ctx *gin.Context) {
	// Reset the user's password
	response, err := c.service.ResetPassword(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}

// deleteUser deletes the user's account
func (c *Controller) deleteUser(ctx *gin.Context) {
	// Delete the user's account
	response, err := c.service.DeleteUser(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": response})
}
