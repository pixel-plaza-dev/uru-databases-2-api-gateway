package branches

import (
	"github.com/gin-gonic/gin"
	moduleshopsproducts "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/businesses/branches/products"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestbranches "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/businesses/branches"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the branches module
// @Summary Shops Branches Router Group
// @Description Router group for shops branches-related endpoints
// @Tags v1 shops businesses branches
// @Accept json
// @Produce json
// @Router /api/v1/shops/businesses/branches [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new branches controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the branches controller
	route := baseRoute.Group(pbconfigrestbranches.Base.String())

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
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestbranches.AddBranchMapper, c.addBranch))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestbranches.GetBranchMapper, c.getBranch))
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbranches.GetBusinessBranchesMapper,
			c.getBusinessBranches,
		),
	)
	c.route.PUT(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestbranches.UpdateBranchMapper, c.updateBranch))
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestbranches.SuspendBranchMapper, c.suspendBranch))
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbranches.ActivateBranchMapper,
			c.activateBranch,
		),
	)
	c.route.DELETE(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestbranches.DeleteBranchMapper, c.deleteBranch))

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	productsController := moduleshopsproducts.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		productsController,
	} {
		controller.Initialize()
	}
}

// addBranch adds a new branch
// @Summary Add a new branch
// @Description Add a new branch
// @Tags v1 shops businesses branches
// @Accept json
// @Produce json
// @Param request body pbshop.AddBranchRequest true "Add Branch Request"
// @Success 201 {object} pbshop.AddBranchResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/branches [post]
func (c *Controller) addBranch(ctx *gin.Context) {
	var request pbshop.AddBranchRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new branch
	response, err := c.client.AddBranch(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getBranch gets a branch by ID
// @Summary Get a branch by ID
// @Description Get a branch by ID
// @Tags v1 shops businesses branches
// @Accept json
// @Produce json
// @Param branchId path string true "Branch ID"
// @Success 200 {object} pbshop.GetBranchResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/businesses/branches/{branchId} [get]
func (c *Controller) getBranch(ctx *gin.Context) {
	var request pbshop.GetBranchRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the branch ID from the path
	request.BranchId = ctx.Param(typesrest.BranchId.String())

	// Get the branch by ID
	response, err := c.client.GetBranch(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getBusinessBranches gets all branches for a business
// @Summary Get all branches for a business
// @Description Get all branches for a business
// @Tags v1 shops businesses branches
// @Accept json
// @Produce json
// @Param businessId path string true "Business ID"
// @Success 200 {object} pbshop.GetBusinessBranchesResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/businesses/branches/business-id/{businessId} [get]
func (c *Controller) getBusinessBranches(ctx *gin.Context) {
	var request pbshop.GetBusinessBranchesRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the business ID from the path
	request.BusinessId = ctx.Param(typesrest.BusinessId.String())

	// Get all branches for the business
	response, err := c.client.GetBusinessBranches(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// updateBranch updates a branch
// @Summary Update a branch
// @Description Update a branch
// @Tags v1 shops businesses branches
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateBranchRequest true "Update Branch Request"
// @Success 200 {object} pbshop.UpdateBranchResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/branches [put]
func (c *Controller) updateBranch(ctx *gin.Context) {
	var request pbshop.UpdateBranchRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the branch
	response, err := c.client.UpdateBranch(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// suspendBranch suspends a branch
// @Summary Suspend a branch
// @Description Suspend a branch
// @Tags v1 shops businesses branches
// @Accept json
// @Produce json
// @Param branchId path string true "Branch ID"
// @Success 200 {object} pbshop.SuspendBranchResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/branches/suspend/{branchId} [post]
func (c *Controller) suspendBranch(ctx *gin.Context) {
	var request pbshop.SuspendBranchRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the branch ID from the path
	request.BranchId = ctx.Param(typesrest.BranchId.String())

	// Suspend the branch
	response, err := c.client.SuspendBranch(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// activateBranch activates a branch
// @Summary Activate a branch
// @Description Activate a branch
// @Tags v1 shops businesses branches
// @Accept json
// @Produce json
// @Param branchId path string true "Branch ID"
// @Success 200 {object} pbshop.ActivateBranchResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/branches/activate/{branchId} [post]
func (c *Controller) activateBranch(ctx *gin.Context) {
	var request pbshop.ActivateBranchRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the branch ID from the path
	request.BranchId = ctx.Param(typesrest.BranchId.String())

	// Activate the branch
	response, err := c.client.ActivateBranch(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// deleteBranch deletes a branch
// @Summary Delete a branch
// @Description Delete a branch
// @Tags v1 shops businesses branches
// @Accept json
// @Produce json
// @Param branchId path string true "Branch ID"
// @Success 200 {object} pbshop.DeleteBranchResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/branches/{branchId} [delete]
func (c *Controller) deleteBranch(ctx *gin.Context) {
	var request pbshop.DeleteBranchRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the branch ID from the path
	request.BranchId = ctx.Param(typesrest.BranchId.String())

	// Delete the branch
	response, err := c.client.DeleteBranch(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
