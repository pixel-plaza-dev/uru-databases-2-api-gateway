package phone_numbers

import (
	"github.com/gin-gonic/gin"
	appgrpcuser "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/user"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigrestphonenumbers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users/phone-numbers"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the users phone numbers module
// @Summary Users Phone Numbers Router Group
// @Description Router group for users phone numbers-related endpoints
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Router /api/v1/users/phone-numbers [group]
type Controller struct {
	route           *gin.RouterGroup
	service         *appgrpcuser.Service
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new phone numbers controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcuser.Service,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the phone numbers controller
	route := baseRoute.Group(pbconfigrestphonenumbers.Base.String())

	// Create a new user controller
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
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestphonenumbers.GetPhoneNumberMapper,
			c.getPhoneNumber,
		),
	)
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestphonenumbers.ChangePhoneNumberMapper,
			c.changePhoneNumber,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestphonenumbers.SendVerificationSMSMapper,
			c.sendVerificationSMS,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestphonenumbers.VerifyPhoneNumberMapper,
			c.verifyPhoneNumber,
		),
	)
}

// getPhoneNumber gets the user's phone number
// @Summary Get user phone number
// @Description Get the user's phone number
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Success 200 {object} pbuser.GetPhoneNumberResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/phone-numbers [get]
func (c *Controller) getPhoneNumber(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the user's active phone numbers
	response, err := c.service.GetPhoneNumber(ctx, grpcCtx)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// changePhoneNumber changes the user's phone number
// @Summary Change user phone number
// @Description Change the user's phone number
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Param request body pbuser.ChangePhoneNumberRequest true "Change Phone Number Request"
// @Success 200 {object} pbuser.ChangePhoneNumberResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/phone-numbers [put]
func (c *Controller) changePhoneNumber(ctx *gin.Context) {
	var request pbuser.ChangePhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Change the user's phone number
	response, err := c.service.ChangePhoneNumber(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// sendVerificationSMS sends a verification SMS to a user
// @Summary Send verification SMS
// @Description Send a verification SMS to a user
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Param request body pbuser.SendVerificationSMSRequest true "Send Verification SMS Request"
// @Success 200 {object} pbuser.SendVerificationSMSResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/phone-numbers/send-verification [post]
func (c *Controller) sendVerificationSMS(ctx *gin.Context) {
	var request pbuser.SendVerificationSMSRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Send a verification phone number
	response, err := c.service.SendVerificationSMS(
		ctx,
		grpcCtx,
		&request,
	)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// verifyEmail verifies the user's phone number
// @Summary Verify phone number
// @Description Verify the user's phone number
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Param token path string true "Verification Token"
// @Success 200 {object} pbuser.VerifyPhoneNumberResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/phone-numbers/verify/{token} [post]
func (c *Controller) verifyPhoneNumber(ctx *gin.Context) {
	var request pbuser.VerifyPhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the token to the request
	request.Token = ctx.Param(pbtypesrest.Token.String())

	// Verify the user's phone number
	response, err := c.service.VerifyPhoneNumber(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
