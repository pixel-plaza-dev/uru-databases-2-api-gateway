package products

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestproducts "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/businesses/products"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the businesses products module
// @Summary Shops Businesses Products Router Group
// @Description Router group for shops businesses products-related endpoints
// @Tags v1 shops businesses products
// @Accept json
// @Produce json
// @Router /api/v1/shops/businesses/products [group]
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
			pbconfigrestproducts.AddBusinessProductMapper,
			c.addBusinessProduct,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.GetBusinessProductMapper,
			c.getBusinessProduct,
		),
	)
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.UpdateBusinessProductMapper,
			c.updateBusinessProduct,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.SearchBusinessProductsMapper,
			c.searchBusinessProducts,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.ActivateBusinessProductMapper,
			c.activateBusinessProduct,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.SuspendBusinessProductMapper,
			c.suspendBusinessProduct,
		),
	)
}

// addBusinessProduct adds a new business product
// @Summary Add a new business product
// @Description Add a new business product
// @Tags v1 shops businesses products
// @Accept json
// @Produce json
// @Param request body pbshop.AddBusinessProductRequest true "Add Business Product Request"
// @Success 201 {object} pbshop.AddBusinessProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/products [post]
func (c *Controller) addBusinessProduct(ctx *gin.Context) {
	var request pbshop.AddBusinessProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new business product
	response, err := c.client.AddBusinessProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getBusinessProduct gets a business product by ID
// @Summary Get a business product by ID
// @Description Get a business product by ID
// @Tags v1 shops businesses products
// @Accept json
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} pbshop.GetBusinessProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/businesses/products/{productId} [get]
func (c *Controller) getBusinessProduct(ctx *gin.Context) {
	var request pbshop.GetBusinessProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the product ID from the path
	request.ProductId = ctx.Param(typesrest.ProductId.String())

	// Get the business product by ID
	response, err := c.client.GetBusinessProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// updateBusinessProduct updates a business product
// @Summary Update a business product
// @Description Update a business product
// @Tags v1 shops businesses products
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateBusinessProductRequest true "Update Business Product Request"
// @Success 200 {object} pbshop.UpdateBusinessProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/products [put]
func (c *Controller) updateBusinessProduct(ctx *gin.Context) {
	var request pbshop.UpdateBusinessProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the business product
	response, err := c.client.UpdateBusinessProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// searchBusinessProducts searches for business products
// @Summary Search for business products
// @Description Search for business products
// @Tags v1 shops businesses products
// @Accept json
// @Produce json
// @Param request body pbshop.SearchBusinessProductsRequest true "Search Business Products Request"
// @Success 200 {object} pbshop.SearchBusinessProductsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/businesses/products/search [post]
func (c *Controller) searchBusinessProducts(ctx *gin.Context) {
	var request pbshop.SearchBusinessProductsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// ADD PARAMETERS HERE

	// Search for business products
	response, err := c.client.SearchBusinessProducts(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// activateBusinessProduct activates a business product
// @Summary Activate a business product
// @Description Activate a business product
// @Tags v1 shops businesses products
// @Accept json
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} pbshop.ActivateBusinessProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/products/activate/{productId} [post]
func (c *Controller) activateBusinessProduct(ctx *gin.Context) {
	var request pbshop.ActivateBusinessProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the product ID from the path
	request.ProductId = ctx.Param(typesrest.ProductId.String())

	// Activate the business product
	response, err := c.client.ActivateBusinessProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// suspendBusinessProduct suspends a business product
// @Summary Suspend a business product
// @Description Suspend a business product
// @Tags v1 shops businesses products
// @Accept json
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} pbshop.SuspendBusinessProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/products/suspend/{productId} [post]
func (c *Controller) suspendBusinessProduct(ctx *gin.Context) {
	var request pbshop.SuspendBusinessProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the product ID from the path
	request.ProductId = ctx.Param(typesrest.ProductId.String())

	// Suspend the business product
	response, err := c.client.SuspendBusinessProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
