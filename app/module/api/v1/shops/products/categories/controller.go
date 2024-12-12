package categories

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestcategories "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/products/categories"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the products categories module
// @Summary Shops Markets Categories Router Group
// @Description Router group for shops products categories-related endpoints
// @Tags v1 shops products categories
// @Accept json
// @Produce json
// @Router /api/v1/shops/products/categories [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new products categories controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the products categories controller
	route := baseRoute.Group(pbconfigrestcategories.Base.String())

	// Create a new products categories controller
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
			pbconfigrestcategories.AddProductCategoryMapper,
			c.addProductCategory,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestcategories.GetProductCategoryMapper,
			c.getProductCategory,
		),
	)
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestcategories.UpdateProductCategoryMapper,
			c.updateProductCategory,
		),
	)
}

// addProductCategory adds a new product category
// @Summary Add a new product category
// @Description Add a new product category
// @Tags v1 shops products categories
// @Accept json
// @Produce json
// @Param request body pbshop.AddProductCategoryRequest true "Add Product Category Request"
// @Success 201 {object} pbshop.AddProductCategoryResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/products/categories [post]
func (c *Controller) addProductCategory(ctx *gin.Context) {
	var request pbshop.AddProductCategoryRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new product category
	response, err := c.client.AddProductCategory(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getProductCategory gets a product category by ID
// @Summary Get a product category by ID
// @Description Get a product category by ID
// @Tags v1 shops products categories
// @Accept json
// @Produce json
// @Param categoryId path string true "Category ID"
// @Success 200 {object} pbshop.GetProductCategoryResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/products/categories/{categoryId} [get]
func (c *Controller) getProductCategory(ctx *gin.Context) {
	var request pbshop.GetProductCategoryRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the category ID from the path
	request.ProductCategoryId = ctx.Param(typesrest.CategoryId.String())

	// Get the product category by ID
	response, err := c.client.GetProductCategory(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// updateProductCategory updates a product category
// @Summary Update a product category
// @Description Update a product category
// @Tags v1 shops products categories
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateProductCategoryRequest true "Update Product Category Request"
// @Success 200 {object} pbshop.UpdateProductCategoryResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/products/categories [put]
func (c *Controller) updateProductCategory(ctx *gin.Context) {
	var request pbshop.UpdateProductCategoryRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the product category
	response, err := c.client.UpdateProductCategory(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
