package phone_numbers

import (
	"github.com/gin-gonic/gin"
	appgrpcuser "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/user"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/user"
	pbconfigrestusers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/users"
	pbconfigrestphonenumbers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/users/phone-numbers"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
	"net/http"
)

// Controller struct for the users phone numbers module
// @Summary Users Phone Numbers Router Group
// @Description Router group for users phone numbers-related endpoints
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Router /phone-numbers [group]
type Controller struct {
	route   *gin.RouterGroup
	service *appgrpcuser.Service
}

// NewController creates a new phone numbers controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcuser.Service,
) *Controller {
	// Create a new route for the phone numbers controller
	route := baseRoute.Group(pbconfigrestusers.PhoneNumbers.String())

	// Create a new user controller
	return &Controller{
		route:   route,
		service: service,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.GET(pbconfigrestphonenumbers.Relative.String(), c.getPhoneNumber)
	c.route.PUT(pbconfigrestphonenumbers.Relative.String(), c.changePhoneNumber)
	c.route.POST(pbconfigrestphonenumbers.SendVerification.String(), c.sendVerificationSMS)
	c.route.POST(
		pbconfigrestphonenumbers.VerifyByToken.String(),
		c.verifyPhoneNumber,
	)
}

// getPhoneNumber gets the user's phone number
// @Summary Get user phone number
// @Description Get the user's phone number
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Success 200 {object} pbuser.GetPhoneNumberResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} commongin.InternalServerError
// @Router / [get]
func (c *Controller) getPhoneNumber(ctx *gin.Context) {
	var request pbuser.GetPhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
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

// changePhoneNumber changes the user's phone number
// @Summary Change user phone number
// @Description Change the user's phone number
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Param request body pbuser.ChangePhoneNumberRequest true "Change Phone Number Request"
// @Success 200 {object} pbuser.ChangePhoneNumberResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} commongin.InternalServerError
// @Router /phone-numbers [put]
func (c *Controller) changePhoneNumber(ctx *gin.Context) {
	var request pbuser.ChangePhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
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

// sendVerificationSMS sends a verification SMS to a user
// @Summary Send verification SMS
// @Description Send a verification SMS to a user
// @Tags v1 users phone-numbers
// @Accept json
// @Produce json
// @Param request body pbuser.SendVerificationSMSRequest true "Send Verification SMS Request"
// @Success 200 {object} pbuser.SendVerificationSMSResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} commongin.InternalServerError
// @Router /send-verification [post]
func (c *Controller) sendVerificationSMS(ctx *gin.Context) {
	var request pbuser.SendVerificationSMSRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
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

// verifyEmail verifies the user's email
// @Summary Verify email
// @Description Verify the user's email
// @Tags v1 profiles
// @Accept json
// @Produce json
// @Param token path string true "Verification Token"
// @Success 200 {object} pbuser.VerifyEmailResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} commongin.InternalServerError
// @Router /verify/{token} [post]
func (c *Controller) verifyPhoneNumber(ctx *gin.Context) {
	var request pbuser.VerifyPhoneNumberRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
		)
		return
	}

	// Add the token to the request
	request.Token = ctx.Param(pbtypesrest.Token.String())

	// Verify the user's phone number
	response, err := c.service.VerifyPhoneNumber(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}