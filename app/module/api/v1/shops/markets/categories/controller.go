package categories

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestcategories "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/markets/categories"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the markets categories module
// @Summary Shops Markets Categories Router Group
// @Description Router group for shops markets categories-related endpoints
// @Tags v1 shops markets categories
// @Accept json
// @Produce json
// @Router /api/v1/shops/markets/categories [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new markets categories controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the markets categories controller
	route := baseRoute.Group(pbconfigrestcategories.Base.String())

	// Create a new markets categories controller
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
			pbconfigrestcategories.AddMarketCategoryMapper,
			c.addMarketCategory,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestcategories.GetMarketCategoryMapper,
			c.getMarketCategory,
		),
	)
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestcategories.UpdateMarketCategoryMapper,
			c.updateMarketCategory,
		),
	)
}

// addMarketCategory adds a new market category
// @Summary Add a new market category
// @Description Add a new market category
// @Tags v1 shops markets categories
// @Accept json
// @Produce json
// @Param request body pbshop.AddMarketCategoryRequest true "Add Market Category Request"
// @Success 201 {object} pbshop.AddMarketCategoryResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/markets/categories [post]
func (c *Controller) addMarketCategory(ctx *gin.Context) {
	var request pbshop.AddMarketCategoryRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new market category
	response, err := c.client.AddMarketCategory(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getMarketCategory gets a market category by ID
// @Summary Get a market category by ID
// @Description Get a market category by ID
// @Tags v1 shops markets categories
// @Accept json
// @Produce json
// @Param categoryId path string true "Category ID"
// @Success 200 {object} pbshop.GetMarketCategoryResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/markets/categories/{categoryId} [get]
func (c *Controller) getMarketCategory(ctx *gin.Context) {
	var request pbshop.GetMarketCategoryRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the category ID from the path
	request.MarketCategoryId = ctx.Param(typesrest.CategoryId.String())

	// Get the market category by ID
	response, err := c.client.GetMarketCategory(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// updateMarketCategory updates a market category
// @Summary Update a market category
// @Description Update a market category
// @Tags v1 shops markets categories
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateMarketCategoryRequest true "Update Market Category Request"
// @Success 200 {object} pbshop.UpdateMarketCategoryResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/markets/categories [put]
func (c *Controller) updateMarketCategory(ctx *gin.Context) {
	var request pbshop.UpdateMarketCategoryRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the market category
	response, err := c.client.UpdateMarketCategory(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
