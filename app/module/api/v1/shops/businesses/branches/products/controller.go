package products

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestproducts "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/businesses/branches/products"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the branches products module
// @Summary Shops Branches Products Router Group
// @Description Router group for shops branches products-related endpoints
// @Tags v1 shops businesses branches products
// @Accept json
// @Produce json
// @Router /api/v1/shops/businesses/branches/products [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new products controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the products controller
	route := baseRoute.Group(pbconfigrestproducts.Base.String())

	// Create a new products controller
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
			pbconfigrestproducts.AddBranchProductMapper,
			c.addBranchProduct,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.GetBranchProductMapper,
			c.getBranchProduct,
		),
	)
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.UpdateBranchProductMapper,
			c.updateBranchProduct,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.SearchBranchProductsMapper,
			c.searchBranchProducts,
		),
	)
}

// addBranchProduct adds a new branch product
// @Summary Add a new branch product
// @Description Add a new branch product
// @Tags v1 shops businesses branches products
// @Accept json
// @Produce json
// @Param request body pbshop.AddBranchProductRequest true "Add Branch Product Request"
// @Success 201 {object} pbshop.AddBranchProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/branches/products [post]
func (c *Controller) addBranchProduct(ctx *gin.Context) {
	var request pbshop.AddBranchProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new branch product
	response, err := c.client.AddBranchProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getBranchProduct gets a branch product by ID
// @Summary Get a branch product by ID
// @Description Get a branch product by ID
// @Tags v1 shops businesses branches products
// @Accept json
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} pbshop.GetBranchProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/businesses/branches/products/{productId} [get]
func (c *Controller) getBranchProduct(ctx *gin.Context) {
	var request pbshop.GetBranchProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the product ID from the path
	request.ProductId = ctx.Param(typesrest.ProductId.String())

	// Get the branch product by ID
	response, err := c.client.GetBranchProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// updateBranchProduct updates a branch product
// @Summary Update a branch product
// @Description Update a branch product
// @Tags v1 shops businesses branches products
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateBranchProductRequest true "Update Branch Product Request"
// @Success 200 {object} pbshop.UpdateBranchProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/branches/products [put]
func (c *Controller) updateBranchProduct(ctx *gin.Context) {
	var request pbshop.UpdateBranchProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the branch product
	response, err := c.client.UpdateBranchProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// searchBranchProducts searches for branch products
// @Summary Search for branch products
// @Description Search for branch products
// @Tags v1 shops businesses branches products
// @Accept json
// @Produce json
// @Param request body pbshop.SearchBranchProductsRequest true "Search Branch Products Request"
// @Success 200 {object} pbshop.SearchBranchProductsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/branches/products/search [post]
func (c *Controller) searchBranchProducts(ctx *gin.Context) {
	var request pbshop.SearchBranchProductsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// ADD PARAMETERS HERE

	// Search for branch products
	response, err := c.client.SearchBranchProducts(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
