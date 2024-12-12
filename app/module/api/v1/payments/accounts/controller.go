package accounts

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/payment"
	pbconfigrestaccounts "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/payments/accounts"
	"net/http"
)

// Controller struct for the payments accounts module
// @Summary Payments Accounts Router Group
// @Description Router group for payments accounts-related endpoints
// @Tags v1 payments accounts
// @Accept json
// @Produce json
// @Router /api/v1/payments/accounts [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbpayment.PaymentClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new accounts controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbpayment.PaymentClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the accounts controller
	route := baseRoute.Group(pbconfigrestaccounts.Base.String())

	// Create a new accounts controller
	return &Controller{
		route:           route,
		client:          client,
		routeHandler:    routeHandler,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestaccounts.AddPaymentAccountMapper,
			c.addPaymentAccount,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestaccounts.GetPaymentAccountsMapper,
			c.getPaymentAccounts,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestaccounts.GetActivePaymentAccountsMapper,
			c.getActivePaymentAccounts,
		),
	)
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestaccounts.ActivatePaymentAccountMapper,
			c.activatePaymentAccount,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestaccounts.GetSuspendedPaymentAccountsMapper,
			c.getSuspendedPaymentAccounts,
		),
	)
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestaccounts.SuspendPaymentAccountMapper,
			c.suspendPaymentAccount,
		),
	)
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestaccounts.VerifyPaymentMapper, c.verifyPayment))
}

// addPaymentAccount adds a new payment account
// @Summary Add a new payment account
// @Description Add a new payment account
// @Tags v1 payments accounts
// @Accept json
// @Produce json
// @Param request body pbpayment.AddPaymentAccountRequest true "Add Payment Account Request"
// @Success 201 {object} pbpayment.AddPaymentAccountResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/accounts [post]
func (c *Controller) addPaymentAccount(ctx *gin.Context) {
	var request pbpayment.AddPaymentAccountRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new payment account
	response, err := c.client.AddPaymentAccount(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getPaymentAccounts gets payment accounts
// @Summary Get payment accounts
// @Description Get payment accounts
// @Tags v1 payments accounts
// @Accept json
// @Produce json
// @Success 200 {object} pbpayment.GetPaymentAccountsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/accounts [get]
func (c *Controller) getPaymentAccounts(ctx *gin.Context) {
	var request pbpayment.GetPaymentAccountsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get payment accounts
	response, err := c.client.GetPaymentAccounts(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getActivePaymentAccounts gets active payment accounts
// @Summary Get active payment accounts
// @Description Get active payment accounts
// @Tags v1 payments accounts
// @Accept json
// @Produce json
// @Success 200 {object} pbpayment.GetActivePaymentAccountsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/accounts/active [get]
func (c *Controller) getActivePaymentAccounts(ctx *gin.Context) {
	var request pbpayment.GetActivePaymentAccountsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get active payment accounts
	response, err := c.client.GetActivePaymentAccounts(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// activatePaymentAccount activates a payment account
// @Summary Activate a payment account
// @Description Activate a payment account
// @Tags v1 payments accounts
// @Accept json
// @Produce json
// @Param request body pbpayment.ActivatePaymentAccountRequest true "Activate Payment Account Request"
// @Success 200 {object} pbpayment.ActivatePaymentAccountResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/accounts/activate [put]
func (c *Controller) activatePaymentAccount(ctx *gin.Context) {
	var request pbpayment.ActivatePaymentAccountRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Activate the payment account
	response, err := c.client.ActivatePaymentAccount(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getSuspendedPaymentAccounts gets suspended payment accounts
// @Summary Get suspended payment accounts
// @Description Get suspended payment accounts
// @Tags v1 payments accounts
// @Accept json
// @Produce json
// @Success 200 {object} pbpayment.GetSuspendedPaymentAccountsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/accounts/suspended [get]
func (c *Controller) getSuspendedPaymentAccounts(ctx *gin.Context) {
	var request pbpayment.GetSuspendedPaymentAccountsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get suspended payment accounts
	response, err := c.client.GetSuspendedPaymentAccounts(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// suspendPaymentAccount suspends a payment account
// @Summary Suspend a payment account
// @Description Suspend a payment account
// @Tags v1 payments accounts
// @Accept json
// @Produce json
// @Param request body pbpayment.SuspendPaymentAccountRequest true "Suspend Payment Account Request"
// @Success 200 {object} pbpayment.SuspendPaymentAccountResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/accounts/suspend [put]
func (c *Controller) suspendPaymentAccount(ctx *gin.Context) {
	var request pbpayment.SuspendPaymentAccountRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Suspend the payment account
	response, err := c.client.SuspendPaymentAccount(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// verifyPayment verifies a payment
// @Summary Verify a payment
// @Description Verify a payment
// @Tags v1 payments accounts
// @Accept json
// @Produce json
// @Param request body pbpayment.VerifyPaymentRequest true "Verify Payment Request"
// @Success 200 {object} pbpayment.VerifyPaymentResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/accounts/verify [post]
func (c *Controller) verifyPayment(ctx *gin.Context) {
	var request pbpayment.VerifyPaymentRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Verify the payment
	response, err := c.client.VerifyPayment(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
