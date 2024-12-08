package users

import (
	"github.com/gin-gonic/gin"
	appgrpcuser "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/user"
	moduleusersemails "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/users/emails"
	moduleusersphonenumbers "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/users/phone-numbers"
	moduleusersprofiles "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/users/profiles"
	moduleusersusernames "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/users/usernames"
	apptypescontroller "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types/controller"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfiggrpcuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/user"
	pbconfigrestusers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct
// @Summary Users Router Group
// @Description Router group for users-related endpoints
// @Tags v1 users
// @Accept json
// @Produce json
// @Router /api/v1/users [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbuser.UserClient
	service         *appgrpcuser.Service
	authMiddleware  authmiddleware.Authentication
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new user controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbuser.UserClient,
	authMiddleware authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the user controller
	route := baseRoute.Group(pbconfigrestusers.Base.String())

	// Create the route handler
	routeHandler := commonhandler.NewDefaultHandler(authMiddleware, &pbconfiggrpcuser.Interceptions)

	// Create the user service
	service := appgrpcuser.NewService(client, responseHandler)

	// Create a new user controller
	return &Controller{
		route:           route,
		client:          client,
		service:         service,
		authMiddleware:  authMiddleware,
		routeHandler:    routeHandler,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.PATCH(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestusers.UpdateUserMapper, c.updateUser))
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestusers.SignUpMapper, c.signUp))
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestusers.GetUserIdByUsernameMapper,
			c.getUserIdByUsername,
		),
	)
	c.route.PATCH(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestusers.ChangePasswordMapper, c.changePassword))
	c.route.PATCH(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestusers.ChangeUsernameMapper, c.changeUsername))
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestusers.ForgotPasswordMapper, c.forgotPassword))
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestusers.ResetPasswordMapper, c.resetPassword))
	c.route.DELETE(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestusers.DeleteAccountMapper, c.deleteUser))

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	emailsController := moduleusersemails.NewController(c.route, c.service, c.routeHandler, c.responseHandler)
	phoneNumbersController := moduleusersphonenumbers.NewController(
		c.route,
		c.service,
		c.routeHandler,
		c.responseHandler,
	)
	profilesController := moduleusersprofiles.NewController(c.route, c.service, c.routeHandler, c.responseHandler)
	usernamesController := moduleusersusernames.NewController(c.route, c.service, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypescontroller.Controller{
		emailsController,
		phoneNumbersController,
		profilesController,
		usernamesController,
	} {
		controller.Initialize()
	}
}

// signUp creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Tags v1 users
// @Accept json
// @Produce json
// @Param request body pbuser.SignUpRequest true "Sign Up Request"
// @Success 201 {object} pbuser.SignUpResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/sign-up [post]
func (c *Controller) signUp(ctx *gin.Context) {
	var request pbuser.SignUpRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Create a new user
	response, err := c.service.SignUp(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// updateUser updates the user
// @Summary Update user
// @Description Update the user
// @Tags v1 users
// @Accept json
// @Produce json
// @Param request body pbuser.UpdateUserRequest true "Update User Request"
// @Success 200 {object} pbuser.UpdateUserResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users [patch]
func (c *Controller) updateUser(ctx *gin.Context) {
	var request pbuser.UpdateUserRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the user
	response, err := c.service.UpdateUser(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getUserIdByUsername gets the user's ID by username
// @Summary Get user ID by username
// @Description Get the user's ID by username
// @Tags v1 users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} pbuser.GetUserIdByUsernameResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/user-id/{username} [get]
func (c *Controller) getUserIdByUsername(ctx *gin.Context) {
	var request pbuser.GetUserIdByUsernameRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the username to the request
	request.Username = ctx.Param(pbtypesrest.Username.String())

	// Get the user's ID by username
	response, err := c.service.GetUserIdByUsername(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// changePassword changes the user's password
// @Summary Change user password
// @Description Change the user's password
// @Tags v1 users
// @Accept json
// @Produce json
// @Param request body pbuser.ChangePasswordRequest true "Change Password Request"
// @Success 200 {object} pbuser.ChangePasswordResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/password [put]
func (c *Controller) changePassword(ctx *gin.Context) {
	var request pbuser.ChangePasswordRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Change the user's password
	response, err := c.service.ChangePassword(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// changeUsername changes the user's username
// @Summary Change user username
// @Description Change the user's username
// @Tags v1 users
// @Accept json
// @Produce json
// @Param request body pbuser.ChangeUsernameRequest true "Change Username Request"
// @Success 200 {object} pbuser.ChangeUsernameResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/username [patch]
func (c *Controller) changeUsername(ctx *gin.Context) {
	var request pbuser.ChangeUsernameRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Change the user's username
	response, err := c.service.ChangeUsername(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// forgotPassword sends a reset password email to a user
// @Summary Send reset password email
// @Description Send a reset password email to a user
// @Tags v1 users
// @Accept json
// @Produce json
// @Param request body pbuser.ForgotPasswordRequest true "Forgot Password Request"
// @Success 200 {object} pbuser.ForgotPasswordResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/forgot-password [post]
func (c *Controller) forgotPassword(ctx *gin.Context) {
	var request pbuser.ForgotPasswordRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Send a reset password email
	response, err := c.service.ForgotPassword(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// resetPassword resets the user's password
// @Summary Reset user password
// @Description Reset the user's password
// @Tags v1 users
// @Accept json
// @Produce json
// @Param token path string true "Verification Token"
// @Param request body pbuser.ResetPasswordRequest true "Reset Password Request"
// @Success 200 {object} pbuser.ResetPasswordResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/reset-password/{token} [post]
func (c *Controller) resetPassword(ctx *gin.Context) {
	var request pbuser.ResetPasswordRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the token to the request
	request.Token = ctx.Param(pbtypesrest.Token.String())

	// Reset the user's password
	response, err := c.service.ResetPassword(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// deleteUser deletes the user's account
// @Summary Delete user account
// @Description Delete the user's account
// @Tags v1 users
// @Accept json
// @Produce json
// @Param request body pbuser.DeleteUserRequest true "Delete User Request"
// @Success 200 {object} pbuser.DeleteUserResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/delete-account [delete]
func (c *Controller) deleteUser(ctx *gin.Context) {
	var request pbuser.DeleteUserRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Delete the user's account
	response, err := c.service.DeleteUser(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
