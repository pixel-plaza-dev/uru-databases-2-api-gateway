package rents

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestrents "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/stores/rents"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the stores rents module
// @Summary Shops Branches Router Group
// @Description Router group for shops stores rents-related endpoints
// @Tags v1 shops stores rents
// @Accept json
// @Produce json
// @Router /api/v1/shops/stores/rents [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new stores rents controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the stores rents controller
	route := baseRoute.Group(pbconfigrestrents.Base.String())

	// Create a new branches controller
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
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestrents.AddBranchRentMapper, c.addBranchRent))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestrents.GetBranchRentsMapper, c.getBranchRents))
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrents.UpdateBranchRentMapper,
			c.updateBranchRent,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrents.GetUnpaidBranchRentsMapper,
			c.getUnpaidBranchRents,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrents.GetBusinessUnpaidBranchRentsMapper,
			c.getBusinessUnpaidBranchRents,
		),
	)
}

// addBranchRent adds a new branch rent
// @Summary Add a new branch rent
// @Description Add a new branch rent
// @Tags v1 shops stores rents
// @Accept json
// @Produce json
// @Param request body pbshop.AddBranchRentRequest true "Add Branch Rent Request"
// @Success 201 {object} pbshop.AddBranchRentResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores/rents [post]
func (c *Controller) addBranchRent(ctx *gin.Context) {
	var request pbshop.AddBranchRentRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new branch rent
	response, err := c.client.AddBranchRent(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getBranchRents gets branch rents by branch ID
// @Summary Get branch rents by branch ID
// @Description Get branch rents by branch ID
// @Tags v1 shops stores rents
// @Accept json
// @Produce json
// @Param branchId path string true "Branch ID"
// @Success 200 {object} pbshop.GetBranchRentsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores/rents/{branchId} [get]
func (c *Controller) getBranchRents(ctx *gin.Context) {
	var request pbshop.GetBranchRentsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the branch ID from the path
	request.BranchId = ctx.Param(typesrest.BranchId.String())

	// Get branch rents by branch ID
	response, err := c.client.GetBranchRents(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// updateBranchRent updates a branch rent
// @Summary Update a branch rent
// @Description Update a branch rent
// @Tags v1 shops stores rents
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateBranchRentRequest true "Update Branch Rent Request"
// @Success 200 {object} pbshop.UpdateBranchRentResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores/rents [put]
func (c *Controller) updateBranchRent(ctx *gin.Context) {
	var request pbshop.UpdateBranchRentRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the branch rent
	response, err := c.client.UpdateBranchRent(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getUnpaidBranchRents gets unpaid branch rents by branch ID
// @Summary Get unpaid branch rents by branch ID
// @Description Get unpaid branch rents by branch ID
// @Tags v1 shops stores rents
// @Accept json
// @Produce json
// @Param branchId path string true "Branch ID"
// @Success 200 {object} pbshop.GetUnpaidBranchRentsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores/rents/branch-unpaid/{branchId} [get]
func (c *Controller) getUnpaidBranchRents(ctx *gin.Context) {
	var request pbshop.GetUnpaidBranchRentsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the branch ID from the path
	request.BranchId = ctx.Param(typesrest.BranchId.String())

	// Get unpaid branch rents by branch ID
	response, err := c.client.GetUnpaidBranchRents(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getBusinessUnpaidBranchRents gets unpaid branch rents by business ID
// @Summary Get unpaid branch rents by business ID
// @Description Get unpaid branch rents by business ID
// @Tags v1 shops stores rents
// @Accept json
// @Produce json
// @Param businessId path string true "Business ID"
// @Success 200 {object} pbshop.GetBusinessUnpaidBranchRentsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores/rents/business-unpaid/{businessId} [get]
func (c *Controller) getBusinessUnpaidBranchRents(ctx *gin.Context) {
	var request pbshop.GetBusinessUnpaidBranchRentsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the business ID from the path
	request.BusinessId = ctx.Param(typesrest.BusinessId.String())

	// Get unpaid branch rents by business ID
	response, err := c.client.GetBusinessUnpaidBranchRents(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
