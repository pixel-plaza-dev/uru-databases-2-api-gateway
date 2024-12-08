package emails

import (
	"github.com/gin-gonic/gin"
	appgrpcuser "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/user"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigrestemails "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users/emails"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the emails module
// @Summary Users Emails Router Group
// @Description Router group for users emails-related endpoints
// @Tags v1 users emails
// @Accept json
// @Produce json
// @Router /api/v1/users/emails [group]
type Controller struct {
	route           *gin.RouterGroup
	service         *appgrpcuser.Service
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new emails controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcuser.Service,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the emails controller
	route := baseRoute.Group(pbconfigrestemails.Base.String())

	// Create a new emails controller
	return &Controller{
		route:           route,
		service:         service,
		routeHandler:    routeHandler,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestemails.GetActiveEmailsMapper, c.getActiveEmails))
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestemails.AddEmailMapper, c.addEmail))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestemails.GetPrimaryEmailMapper, c.getPrimaryEmail))
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestemails.ChangePrimaryEmailMapper,
			c.changePrimaryEmail,
		),
	)
	c.route.DELETE(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestemails.DeleteEmailMapper, c.deleteEmail))
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestemails.SendVerificationEmailMapper,
			c.sendVerificationEmail,
		),
	)
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestemails.VerifyEmailMapper, c.verifyEmail))
}

// getActiveEmails gets the user's active emails
// @Summary Get active emails
// @Description Get the user's active emails
// @Tags v1 users emails
// @Accept json
// @Produce json
// @Success 200 {object} pbuser.GetActiveEmailsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/emails [get]
func (c *Controller) getActiveEmails(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the user's active emails
	response, err := c.service.GetActiveEmails(ctx, grpcCtx)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// addEmail adds an email to the user's account
// @Summary Add email
// @Description Add an email to the user's account
// @Tags v1 users emails
// @Accept json
// @Produce json
// @Param request body pbuser.AddEmailRequest true "Add Email Request"
// @Success 201 {object} pbuser.AddEmailResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/emails [post]
func (c *Controller) addEmail(ctx *gin.Context) {
	var request pbuser.AddEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add an email to the user's account
	response, err := c.service.AddEmail(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getPrimaryEmail gets the user's primary email
// @Summary Get primary email
// @Description Get the user's primary email
// @Tags v1 users emails
// @Accept json
// @Produce json
// @Success 200 {object} pbuser.GetPrimaryEmailResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/emails/primary [get]
func (c *Controller) getPrimaryEmail(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the user's primary email
	response, err := c.service.GetPrimaryEmail(ctx, grpcCtx)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// changePrimaryEmail changes the user's primary email
// @Summary Change primary email
// @Description Change the user's primary email
// @Tags v1 users emails
// @Accept json
// @Produce json
// @Param request body pbuser.ChangePrimaryEmailRequest true "Change Primary Email Request"
// @Success 200 {object} pbuser.ChangePrimaryEmailResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/emails/primary [put]
func (c *Controller) changePrimaryEmail(ctx *gin.Context) {
	var request pbuser.ChangePrimaryEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Change the user's primary email
	response, err := c.service.ChangePrimaryEmail(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// deleteEmail deletes an email from the user's account
// @Summary Delete email
// @Description Delete an email from the user's account
// @Tags v1 users emails
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Success 200 {object} pbuser.DeleteEmailResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/emails/{email} [delete]
func (c *Controller) deleteEmail(ctx *gin.Context) {
	var request pbuser.DeleteEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the email to the request
	request.Email = ctx.Param(pbtypesrest.Email.String())

	// Delete an email from the user's account
	response, err := c.service.DeleteEmail(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// sendVerificationEmail sends a verification email to a user
// @Summary Send verification email
// @Description Send a verification email to a user
// @Tags v1 users emails
// @Accept json
// @Produce json
// @Param request body pbuser.SendVerificationEmailRequest true "Send Verification Email Request"
// @Success 200 {object} pbuser.SendVerificationEmailResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/emails/send-verification [post]
func (c *Controller) sendVerificationEmail(ctx *gin.Context) {
	var request pbuser.SendVerificationEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Send a verification email
	response, err := c.service.SendVerificationEmail(
		ctx, grpcCtx,
		&request,
	)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// verifyEmail verifies the user's email
// @Summary Verify email
// @Description Verify the user's email
// @Tags v1 users emails
// @Accept json
// @Produce json
// @Param token path string true "Verification Token"
// @Success 200 {object} pbuser.VerifyEmailResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/emails/verify/{token} [post]
func (c *Controller) verifyEmail(ctx *gin.Context) {
	var request pbuser.VerifyEmailRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the token to the request
	request.Token = ctx.Param(pbtypesrest.Token.String())

	// Verify the user's email
	response, err := c.service.VerifyEmail(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
