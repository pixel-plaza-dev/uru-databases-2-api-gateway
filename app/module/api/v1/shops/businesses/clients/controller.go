package clients

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestclients "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/businesses/clients"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the businesses clients module
// @Summary Shops Businesses Clients Router Group
// @Description Router group for shops businesses clients-related endpoints
// @Tags v1 shops businesses clients
// @Accept json
// @Produce json
// @Router /api/v1/shops/businesses/clients [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new clients controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the clients controller
	route := baseRoute.Group(pbconfigrestclients.Base.String())

	// Create a new clients controller
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
			pbconfigrestclients.AddBusinessClientMapper,
			c.addBusinessClient,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestclients.IsBusinessClientMapper,
			c.isBusinessClient,
		),
	)
}

// addBusinessClient adds a new business client
// @Summary Add a new business client
// @Description Add a new business client
// @Tags v1 shops businesses clients
// @Accept json
// @Produce json
// @Param request body pbshop.AddBusinessClientRequest true "Add Business Client Request"
// @Success 201 {object} pbshop.AddBusinessClientResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/clients [post]
func (c *Controller) addBusinessClient(ctx *gin.Context) {
	var request pbshop.AddBusinessClientRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new business client
	response, err := c.client.AddBusinessClient(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// isBusinessClient checks if a business is a client
// @Summary Check if a business is a client
// @Description Check if a business is a client
// @Tags v1 shops businesses clients
// @Accept json
// @Produce json
// @Param businessId path string true "Business ID"
// @Success 200 {object} pbshop.IsBusinessClientResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/clients/{businessId} [get]
func (c *Controller) isBusinessClient(ctx *gin.Context) {
	var request pbshop.IsBusinessClientRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the business ID from the path
	request.BusinessId = ctx.Param(typesrest.BusinessId.String())

	// Check if the business is a client
	response, err := c.client.IsBusinessClient(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
