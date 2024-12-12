package orders

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/payment"
	pbconfigrestorders "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/payments/orders"
	"net/http"
)

// Controller struct for the orders module
// @Summary Payments Orders Clients Router Group
// @Description Router group for payments orders-related endpoints
// @Tags v1 payments branch-rents
// @Accept json
// @Produce json
// @Router /api/v1/payments/branch-rents [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbpayment.PaymentClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new orders controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbpayment.PaymentClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the orders controller
	route := baseRoute.Group(pbconfigrestorders.Base.String())

	// Create a new orders controller
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
			pbconfigrestorders.AddOrderPaymentMapper,
			c.addOrderPayment,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestorders.GetOrderPaymentsMapper,
			c.getOrderPayments,
		),
	)
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestorders.PayForOrderMapper, c.payForOrder))
}

// addOrderPayment adds a new order payment
// @Summary Add a new order payment
// @Description Add a new order payment
// @Tags v1 payments orders
// @Accept json
// @Produce json
// @Param request body pbpayment.AddOrderPaymentRequest true "Add Order Payment Request"
// @Success 201 {object} pbpayment.AddOrderPaymentResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/orders [post]
func (c *Controller) addOrderPayment(ctx *gin.Context) {
	var request pbpayment.AddOrderPaymentRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new order payment
	response, err := c.client.AddOrderPayment(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getOrderPayments gets order payments
// @Summary Get order payments
// @Description Get order payments
// @Tags v1 payments orders
// @Accept json
// @Produce json
// @Success 200 {object} pbpayment.GetOrderPaymentsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/orders [get]
func (c *Controller) getOrderPayments(ctx *gin.Context) {
	var request pbpayment.GetOrderPaymentsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get order payments
	response, err := c.client.GetOrderPayments(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// payForOrder processes payment for an order
// @Summary Process payment for an order
// @Description Process payment for an order
// @Tags v1 payments orders
// @Accept json
// @Produce json
// @Param request body pbpayment.PayForOrderRequest true "Pay For Order Request"
// @Success 200 {object} pbpayment.PayForOrderResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/orders/pay [post]
func (c *Controller) payForOrder(ctx *gin.Context) {
	var request pbpayment.PayForOrderRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Process payment for the order
	response, err := c.client.PayForOrder(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
