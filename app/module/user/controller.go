package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/user"
	pbtypes "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	pbuserapi "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/user/api"
	"net/http"
)

type Controller struct {
	apiRoute *gin.RouterGroup
	route    *gin.RouterGroup
	service  *Service
	restMap  *map[string]map[pbtypes.
			RESTMethod]pbtypes.GRPCMethod
	grpcInterceptions *map[pbtypes.GRPCMethod]pbtypes.Interception
	authMiddleware    authmiddleware.Authentication
}

// NewController creates a new controller
func NewController(
	apiRoute *gin.RouterGroup, service *Service,
	jwtValidator commonjwtvalidator.Validator,
	restMap *map[string]map[pbtypes.RESTMethod]pbtypes.
		GRPCMethod,
	grpcInterceptions *map[pbtypes.GRPCMethod]pbtypes.Interception,
	authLogger authmiddleware.Logger,
) (*Controller, error) {
	// Create a new route for the user service
	route := apiRoute.Group(RelativeUri)

	// Create the auth middleware
	authMiddleware, err := authmiddleware.NewMiddleware(
		route.BasePath(),
		jwtValidator,
		restMap,
		grpcInterceptions,
		authLogger,
	)
	if err != nil {
		return nil, err
	}

	// Add the auth middleware to the route
	route.Use(authMiddleware.Authenticate())

	// Create a new user controller
	instance := &Controller{
		apiRoute: apiRoute, route: route, service: service,
		authMiddleware: authMiddleware,
	}

	// Initialize the routes for the controller
	instance.initializeRoutes()

	return instance, nil
}

// initializeRoutes initializes the routes for the controller
func (c *Controller) initializeRoutes() {
	c.route.POST(pbuserapi.SignUp.String(), c.signUp)
	c.route.GET(pbuserapi.Profile.String(), c.getProfile)
	c.route.PATCH(pbuserapi.Profile.String(), c.updateProfile)
	c.route.GET(pbuserapi.FullProfile.String(), c.getFullProfile)
	c.route.GET(pbuserapi.IdByUsername.String(), c.getUserIdByUsername)
	c.route.PUT(pbuserapi.Password.String(), c.changePassword)
	c.route.GET(pbuserapi.UsernameExistsByUsername.String(), c.usernameExists)
	c.route.GET(pbuserapi.UsernameById.String(), c.getUsernameByUserId)
	c.route.PATCH(pbuserapi.Username.String(), c.changeUsername)
	c.route.GET(pbuserapi.Email.String(), c.getPrimaryEmail)
	c.route.PUT(pbuserapi.Email.String(), c.changePrimaryEmail)
	c.route.POST(pbuserapi.Email.String(), c.addEmail)
	c.route.DELETE(pbuserapi.EmailByEmail.String(), c.deleteEmail)
	c.route.GET(pbuserapi.Emails.String(), c.getActiveEmails)
	c.route.POST(
		pbuserapi.SendVerificationEmail.String(),
		c.sendVerificationEmail,
	)
	c.route.POST(pbuserapi.VerifyEmailByToken.String(), c.verifyEmail)
	c.route.GET(pbuserapi.PhoneNumber.String(), c.getPhoneNumber)
	c.route.PUT(pbuserapi.PhoneNumber.String(), c.changePhoneNumber)
	c.route.POST(pbuserapi.SendVerificationSMS.String(), c.sendVerificationSMS)
	c.route.POST(
		pbuserapi.VerifyPhoneNumberByToken.String(),
		c.verifyPhoneNumber,
	)
	c.route.POST(pbuserapi.ForgotPassword.String(), c.forgotPassword)
	c.route.POST(pbuserapi.ResetPasswordByToken.String(), c.resetPassword)
	c.route.DELETE(pbuserapi.DeleteAccount.String(), c.deleteUser)
}

// signUp creates a new user
func (c *Controller) signUp(ctx *gin.Context) {
	var request pbuser.SignUpRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Create a new user
	response, err := c.service.SignUp(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// updateProfile updates the user's profile
func (c *Controller) updateProfile(ctx *gin.Context) {
	var request pbuser.UpdateProfileRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Update the user's profile
	response, err := c.service.UpdateProfile(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getProfile gets the user's profile
func (c *Controller) getProfile(ctx *gin.Context) {
	var request pbuser.GetProfileRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Get the user's profile
	response, err := c.service.GetProfile(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getFullProfile gets the user's full profile
func (c *Controller) getFullProfile(ctx *gin.Context) {
	var request pbuser.GetFullProfileRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
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

// getUserIdByUsername gets the user's ID by username
func (c *Controller) getUserIdByUsername(ctx *gin.Context) {
	var request pbuser.GetUserIdByUsernameRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Add the username to the request
	request.Username = ctx.Param(pbtypes.Username.String())

	// Get the user's ID by username
	response, err := c.service.GetUserIdByUsername(ctx, grpcCtx, &request)
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
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Change the user's password
	response, err := c.service.ChangePassword(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// usernameExists checks if a username exists
func (c *Controller) usernameExists(ctx *gin.Context) {
	var request pbuser.UsernameExistsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Check if the username exists
	response, err := c.service.UsernameExists(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getUsernameByUserId gets the username by user ID
func (c *Controller) getUsernameByUserId(ctx *gin.Context) {
	var request pbuser.GetUsernameByUserIdRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Get the username by user ID
	response, err := c.service.GetUsernameByUserId(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// changeUsername changes the user's username
func (c *Controller) changeUsername(ctx *gin.Context) {
	var request pbuser.ChangeUsernameRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Change the user's username
	response, err := c.service.ChangeUsername(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// changePrimaryEmail changes the user's primary email
func (c *Controller) changePrimaryEmail(ctx *gin.Context) {
	var request pbuser.ChangePrimaryEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Change the user's primary email
	response, err := c.service.ChangePrimaryEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// addEmail adds an email to the user's account
func (c *Controller) addEmail(ctx *gin.Context) {
	var request pbuser.AddEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Add an email to the user's account
	response, err := c.service.AddEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// deleteEmail deletes an email from the user's account
func (c *Controller) deleteEmail(ctx *gin.Context) {
	var request pbuser.DeleteEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Add the email to the request
	request.Email = ctx.Param(pbtypes.Email.String())

	// Delete an email from the user's account
	response, err := c.service.DeleteEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getPrimaryEmail gets the user's primary email
func (c *Controller) getPrimaryEmail(ctx *gin.Context) {
	var request pbuser.GetPrimaryEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Get the user's primary email
	response, err := c.service.GetPrimaryEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getActiveEmails gets the user's active emails
func (c *Controller) getActiveEmails(ctx *gin.Context) {
	var request pbuser.GetActiveEmailsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Get the user's active emails
	response, err := c.service.GetActiveEmails(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// sendVerificationEmail sends a verification email to a user
func (c *Controller) sendVerificationEmail(ctx *gin.Context) {
	var request pbuser.SendVerificationEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Send a verification email
	response, err := c.service.SendVerificationEmail(
		ctx, grpcCtx,
		&request,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// verifyEmail verifies the user's email
func (c *Controller) verifyEmail(ctx *gin.Context) {
	var request pbuser.VerifyEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Add the token to the request
	request.Token = ctx.Param(pbtypes.Token.String())

	// Verify the user's email
	response, err := c.service.VerifyEmail(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// changePhoneNumber changes the user's phone number
func (c *Controller) changePhoneNumber(ctx *gin.Context) {
	var request pbuser.ChangePhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Change the user's phone number
	response, err := c.service.ChangePhoneNumber(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getPhoneNumber gets the user's phone number
func (c *Controller) getPhoneNumber(ctx *gin.Context) {
	var request pbuser.GetPhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Get the user's active phone numbers
	response, err := c.service.GetPhoneNumber(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// sendVerificationSMS sends a verification SMS to a user
func (c *Controller) sendVerificationSMS(ctx *gin.Context) {
	var request pbuser.SendVerificationSMSRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Send a verification phone number
	response, err := c.service.SendVerificationSMS(
		ctx,
		grpcCtx,
		&request,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// verifyPhoneNumber verifies the user's phone number
func (c *Controller) verifyPhoneNumber(ctx *gin.Context) {
	var request pbuser.VerifyPhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Add the token to the request
	request.Token = ctx.Param(pbtypes.Token.String())

	// Verify the user's phone number
	response, err := c.service.VerifyPhoneNumber(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// forgotPassword sends a reset password email to a user
func (c *Controller) forgotPassword(ctx *gin.Context) {
	var request pbuser.ForgotPasswordRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Send a reset password email
	response, err := c.service.ForgotPassword(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// resetPassword resets the user's password
func (c *Controller) resetPassword(ctx *gin.Context) {
	var request pbuser.ResetPasswordRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Add the token to the request
	request.Token = ctx.Param(pbtypes.Token.String())

	// Reset the user's password
	response, err := c.service.ResetPassword(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// deleteUser deletes the user's account
func (c *Controller) deleteUser(ctx *gin.Context) {
	var request pbuser.DeleteUserRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": module.InternalServerError},
		)
		return
	}

	// Delete the user's account
	response, err := c.service.DeleteUser(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
