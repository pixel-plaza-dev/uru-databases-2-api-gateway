package branch_rents

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/payment"
	pbconfigrestbranchrents "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/payments/branch-rents"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the branch rents module
// @Summary Payments Branch Rents Clients Router Group
// @Description Router group for payments branch rents-related endpoints
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

// NewController creates a new branch rents controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbpayment.PaymentClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the branch rents controller
	route := baseRoute.Group(pbconfigrestbranchrents.Base.String())

	// Create a new branch rents controller
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
			pbconfigrestbranchrents.AddBranchRentPaymentMapper,
			c.addBranchRentPayment,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbranchrents.GetBranchRentsPaymentsMapper,
			c.getBranchRentsPayments,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbranchrents.GetBranchRentPaymentsMapper,
			c.getBranchRentPayments,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbranchrents.PayForBranchRentMapper,
			c.payForBranchRent,
		),
	)
}

// addBranchRentPayment adds a new branch rent payment
// @Summary Add a new branch rent payment
// @Description Add a new branch rent payment
// @Tags v1 payments branch-rents
// @Accept json
// @Produce json
// @Param request body pbpayment.AddBranchRentPaymentRequest true "Add Branch Rent Payment Request"
// @Success 201 {object} pbpayment.AddBranchRentPaymentResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/branch-rents [post]
func (c *Controller) addBranchRentPayment(ctx *gin.Context) {
	var request pbpayment.AddBranchRentPaymentRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new branch rent payment
	response, err := c.client.AddBranchRentPayment(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getBranchRentsPayments gets branch rents payments
// @Summary Get branch rents payments
// @Description Get branch rents payments
// @Tags v1 payments branch-rents
// @Accept json
// @Produce json
// @Success 200 {object} pbpayment.GetBranchRentsPaymentsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/branch-rents [get]
func (c *Controller) getBranchRentsPayments(ctx *gin.Context) {
	var request pbpayment.GetBranchRentsPaymentsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get branch rents payments
	response, err := c.client.GetBranchRentsPayments(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getBranchRentPayments gets branch rent payments by branch rent ID
// @Summary Get branch rent payments by branch rent ID
// @Description Get branch rent payments by branch rent ID
// @Tags v1 payments branch-rents
// @Accept json
// @Produce json
// @Param branchRentId path string true "Branch Rent ID"
// @Success 200 {object} pbpayment.GetBranchRentPaymentsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/branch-rents/{branchRentId} [get]
func (c *Controller) getBranchRentPayments(ctx *gin.Context) {
	var request pbpayment.GetBranchRentPaymentsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the branch rent ID from the path
	request.BranchRentId = ctx.Param(typesrest.BranchRentId.String())

	// Get branch rent payments by branch rent ID
	response, err := c.client.GetBranchRentPayments(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// payForBranchRent processes payment for a branch rent
// @Summary Process payment for a branch rent
// @Description Process payment for a branch rent
// @Tags v1 payments branch-rents
// @Accept json
// @Produce json
// @Param request body pbpayment.PayForBranchRentRequest true "Pay For Branch Rent Request"
// @Success 200 {object} pbpayment.PayForBranchRentResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/payments/branch-rents/pay [post]
func (c *Controller) payForBranchRent(ctx *gin.Context) {
	var request pbpayment.PayForBranchRentRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Process payment for the branch rent
	response, err := c.client.PayForBranchRent(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
