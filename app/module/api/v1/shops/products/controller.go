package products

import (
	"github.com/gin-gonic/gin"
	moduleshopscategories "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/markets/categories"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestproducts "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/products"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the products module
// @Summary Shops Products Router Group
// @Description Router group for shops products-related endpoints
// @Tags v1 shops products
// @Accept json
// @Produce json
// @Router /api/v1/shops/products [group]
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
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestproducts.AddProductMapper, c.addProduct))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestproducts.GetProductMapper, c.getProduct))
	c.route.PUT(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestproducts.UpdateProductMapper, c.updateProduct))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestproducts.SearchProductsMapper, c.searchProducts))
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.SuspendProductMapper,
			c.suspendProduct,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestproducts.ActivateProductMapper,
			c.activateProduct,
		),
	)

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	categoriesController := moduleshopscategories.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		categoriesController,
	} {
		controller.Initialize()
	}
}

// addProduct adds a new product
// @Summary Add a new product
// @Description Add a new product
// @Tags v1 shops products
// @Accept json
// @Produce json
// @Param request body pbshop.AddProductRequest true "Add Product Request"
// @Success 201 {object} pbshop.AddProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/products [post]
func (c *Controller) addProduct(ctx *gin.Context) {
	var request pbshop.AddProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new product
	response, err := c.client.AddProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getProduct gets a product by ID
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags v1 shops products
// @Accept json
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} pbshop.GetProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/products/{productId} [get]
func (c *Controller) getProduct(ctx *gin.Context) {
	var request pbshop.GetProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the product ID from the path
	request.ProductId = ctx.Param(typesrest.ProductId.String())

	// Get the product by ID
	response, err := c.client.GetProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// updateProduct updates a product
// @Summary Update a product
// @Description Update a product
// @Tags v1 shops products
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateProductRequest true "Update Product Request"
// @Success 200 {object} pbshop.UpdateProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/products [put]
func (c *Controller) updateProduct(ctx *gin.Context) {
	var request pbshop.UpdateProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the product
	response, err := c.client.UpdateProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// searchProducts searches for products
// @Summary Search for products
// @Description Search for products
// @Tags v1 shops products
// @Accept json
// @Produce json
// @Param request body pbshop.SearchProductsRequest true "Search Products Request"
// @Success 200 {object} pbshop.SearchProductsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/products/search [post]
func (c *Controller) searchProducts(ctx *gin.Context) {
	var request pbshop.SearchProductsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// ADD PARAMETERS HERE

	// Search for products
	response, err := c.client.SearchProducts(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// suspendProduct suspends a product
// @Summary Suspend a product
// @Description Suspend a product
// @Tags v1 shops products
// @Accept json
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} pbshop.SuspendProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/products/suspend/{productId} [post]
func (c *Controller) suspendProduct(ctx *gin.Context) {
	var request pbshop.SuspendProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the product ID from the path
	request.ProductId = ctx.Param(typesrest.ProductId.String())

	// Suspend the product
	response, err := c.client.SuspendProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// activateProduct activates a product
// @Summary Activate a product
// @Description Activate a product
// @Tags v1 shops products
// @Accept json
// @Produce json
// @Param productId path string true "Product ID"
// @Success 200 {object} pbshop.ActivateProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/products/activate/{productId} [post]
func (c *Controller) activateProduct(ctx *gin.Context) {
	var request pbshop.ActivateProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the product ID from the path
	request.ProductId = ctx.Param(typesrest.ProductId.String())

	// Activate the product
	response, err := c.client.ActivateProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
