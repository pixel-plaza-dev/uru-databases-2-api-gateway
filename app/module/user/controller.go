package user

import (
	"github.com/gin-gonic/gin"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/user"
	"net/http"
)

type Controller struct {
	apiRoute      *gin.RouterGroup
	route         *gin.RouterGroup
	service       *Service
	jwtMiddleware authmiddleware.Authentication
}

// NewController creates a new controller
func NewController(
	apiRoute *gin.RouterGroup, service *Service,
	jwtMiddleware authmiddleware.Authentication,
) *Controller {
	// Create a new route for the user service
	route := apiRoute.Group("/user")

	// Create a new user controller
	instance := &Controller{
		apiRoute: apiRoute, route: route, service: service,
		jwtMiddleware: jwtMiddleware,
	}

	// Initialize the routes for the controller
	instance.initializeRoutes()

	return instance
}

// initializeRoutes initializes the routes for the controller
func (c *Controller) initializeRoutes() {
	c.route.POST("/sign-up", c.signUp)
	c.route.PATCH(
		"/profile", c.jwtMiddleware.AuthenticateAccessToken(),
		c.updateProfile,
	)
	c.route.GET(
		"/profile", c.jwtMiddleware.AuthenticateAccessToken(),
		c.getProfile,
	)
	c.route.GET(
		"/full-profile", c.jwtMiddleware.AuthenticateAccessToken(),
		c.getFullProfile,
	)
	c.route.PATCH(
		"/password", c.jwtMiddleware.AuthenticateAccessToken(),
		c.changePassword,
	)
	c.route.PATCH(
		"/username", c.jwtMiddleware.AuthenticateAccessToken(),
		c.changeUsername,
	)
	c.route.PATCH(
		"/email", c.jwtMiddleware.AuthenticateAccessToken(),
		c.changePrimaryEmail,
	)
	c.route.POST(
		"/email", c.jwtMiddleware.AuthenticateAccessToken(),
		c.addEmail,
	)
	c.route.DELETE(
		"/email", c.jwtMiddleware.AuthenticateAccessToken(),
		c.deleteEmail,
	)
	c.route.GET(
		"/email", c.jwtMiddleware.AuthenticateAccessToken(),
		c.getPrimaryEmail,
	)
	c.route.GET(
		"/emails", c.jwtMiddleware.AuthenticateAccessToken(),
		c.getActiveEmails,
	)
	c.route.POST(
		"/send-verification-email", c.jwtMiddleware.AuthenticateAccessToken(),

		c.sendVerificationEmail,
	)
	c.route.POST(
		"/verify-email", c.jwtMiddleware.AuthenticateAccessToken(),
		c.verifyEmail,
	)
	c.route.PATCH(
		"/phone-number", c.jwtMiddleware.AuthenticateAccessToken(),
		c.changePhoneNumber,
	)
	c.route.GET(
		"/phone-number", c.jwtMiddleware.AuthenticateAccessToken(),
		c.getPhoneNumber,
	)
	c.route.POST(
		"/send-verification-phone-number", c.jwtMiddleware.AuthenticateAccessToken(),

		c.sendVerificationPhoneNumber,
	)
	c.route.POST(
		"/verify-phone-number", c.jwtMiddleware.AuthenticateAccessToken(),
		c.verifyPhoneNumber,
	)
	c.route.POST(
		"/forgot-password", c.forgotPassword,
	)
	c.route.POST(
		"/reset-password", c.resetPassword,
	)
	c.route.DELETE(
		"/close-account", c.jwtMiddleware.AuthenticateAccessToken(),
		c.deleteUser,
	)
}

// signUp creates a new user
func (c *Controller) signUp(ctx *gin.Context) {
	var request pbuser.SignUpRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Create a new user
	response, err := c.service.SignUp(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": response})
}

// updateProfile updates the user's profile
func (c *Controller) updateProfile(ctx *gin.Context) {
	var request pbuser.UpdateProfileRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Update the user's profile
	response, err := c.service.UpdateProfile(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// getProfile gets the user's profile
func (c *Controller) getProfile(ctx *gin.Context) {
	var request pbuser.GetProfileRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get the user's profile
	response, err := c.service.GetProfile(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// getFullProfile gets the user's full profile
func (c *Controller) getFullProfile(ctx *gin.Context) {
	var request pbuser.GetFullProfileRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get the user's full profile
	response, err := c.service.GetFullProfile(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// changePassword changes the user's password
func (c *Controller) changePassword(ctx *gin.Context) {
	var request pbuser.ChangePasswordRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Change the user's password
	response, err := c.service.ChangePassword(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// changeUsername changes the user's username
func (c *Controller) changeUsername(ctx *gin.Context) {
	var request pbuser.ChangeUsernameRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Change the user's username
	response, err := c.service.ChangeUsername(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// changePrimaryEmail changes the user's primary email
func (c *Controller) changePrimaryEmail(ctx *gin.Context) {
	var request pbuser.ChangePrimaryEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Change the user's primary email
	response, err := c.service.ChangePrimaryEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// addEmail adds an email to the user's account
func (c *Controller) addEmail(ctx *gin.Context) {
	var request pbuser.AddEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Add an email to the user's account
	response, err := c.service.AddEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": response})
}

// deleteEmail deletes an email from the user's account
func (c *Controller) deleteEmail(ctx *gin.Context) {
	var request pbuser.DeleteEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Delete an email from the user's account
	response, err := c.service.DeleteEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// getPrimaryEmail gets the user's primary email
func (c *Controller) getPrimaryEmail(ctx *gin.Context) {
	var request pbuser.GetPrimaryEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get the user's primary email
	response, err := c.service.GetPrimaryEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// getActiveEmails gets the user's active emails
func (c *Controller) getActiveEmails(ctx *gin.Context) {
	var request pbuser.GetActiveEmailsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get the user's active emails
	response, err := c.service.GetActiveEmails(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// sendVerificationEmail sends a verification email
func (c *Controller) sendVerificationEmail(ctx *gin.Context) {
	var request pbuser.SendVerificationEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Send a verification email
	response, err := c.service.SendVerificationEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// verifyEmail verifies the user's email
func (c *Controller) verifyEmail(ctx *gin.Context) {
	var request pbuser.VerifyEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Verify the user's email
	response, err := c.service.VerifyEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// changePhoneNumber changes the user's phone number
func (c *Controller) changePhoneNumber(ctx *gin.Context) {
	var request pbuser.ChangePhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Change the user's phone number
	response, err := c.service.ChangePhoneNumber(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// getPhoneNumber gets the user's phone number
func (c *Controller) getPhoneNumber(ctx *gin.Context) {
	var request pbuser.GetPhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get the user's active phone numbers
	response, err := c.service.GetPhoneNumber(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// sendVerificationPhoneNumber sends a verification phone number
func (c *Controller) sendVerificationPhoneNumber(ctx *gin.Context) {
	var request pbuser.SendVerificationPhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Send a verification phone number
	response, err := c.service.SendVerificationPhoneNumber(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// verifyPhoneNumber verifies the user's phone number
func (c *Controller) verifyPhoneNumber(ctx *gin.Context) {
	var request pbuser.VerifyPhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Verify the user's phone number
	response, err := c.service.VerifyPhoneNumber(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// forgotPassword sends a reset password email
func (c *Controller) forgotPassword(ctx *gin.Context) {
	var request pbuser.ForgotPasswordRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Send a reset password email
	response, err := c.service.ForgotPassword(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// resetPassword resets the user's password
func (c *Controller) resetPassword(ctx *gin.Context) {
	var request pbuser.ResetPasswordRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Reset the user's password
	response, err := c.service.ResetPassword(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}

// deleteUser deletes the user's account
func (c *Controller) deleteUser(ctx *gin.Context) {
	var request pbuser.DeleteUserRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Delete the user's account
	response, err := c.service.DeleteUser(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})
}
