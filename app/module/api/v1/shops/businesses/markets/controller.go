package markets

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestmarkets "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/businesses/markets"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the businesses markets module
// @Summary Shops Businesses Markets Router Group
// @Description Router group for shops businesses markets-related endpoints
// @Tags v1 shops businesses markets
// @Accept json
// @Produce json
// @Router /api/v1/shops/businesses/markets [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new markets controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the markets controller
	route := baseRoute.Group(pbconfigrestmarkets.Base.String())

	// Create a new markets controller
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
			pbconfigrestmarkets.AddBusinessMarketCategoryMapper,
			c.addBusinessMarketCategory,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestmarkets.GetBusinessMarketCategoriesMapper,
			c.getBusinessMarketCategories,
		),
	)
}

// addBusinessMarketCategory adds a new business market category
// @Summary Add a new business market category
// @Description Add a new business market category
// @Tags v1 shops businesses markets
// @Accept json
// @Produce json
// @Param request body pbshop.AddBusinessMarketCategoryRequest true "Add Business Market Category Request"
// @Success 201 {object} pbshop.AddBusinessMarketCategoryResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/markets/categories [post]
func (c *Controller) addBusinessMarketCategory(ctx *gin.Context) {
	var request pbshop.AddBusinessMarketCategoryRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new business market category
	response, err := c.client.AddBusinessMarketCategory(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getBusinessMarketCategories gets all business market categories
// @Summary Get all business market categories
// @Description Get all business market categories
// @Tags v1 shops businesses markets
// @Accept json
// @Produce json
// @Param businessId path string true "Business ID"
// @Success 200 {object} pbshop.GetBusinessMarketCategoriesResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/businesses/markets/categories/{businessId} [get]
func (c *Controller) getBusinessMarketCategories(ctx *gin.Context) {
	var request pbshop.GetBusinessMarketCategoriesRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the business ID from the path
	request.BusinessId = ctx.Param(typesrest.BusinessId.String())

	// Get all business market categories
	response, err := c.client.GetBusinessMarketCategories(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
