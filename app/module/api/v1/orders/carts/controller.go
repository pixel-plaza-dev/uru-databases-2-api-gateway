package carts

import (
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"

	"github.com/gin-gonic/gin"
	moduleorderscurrent "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/orders/carts/current"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pborder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/order"
	pbconfigrestcarts "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/orders/carts"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the orders carts module
// @Summary Orders Carts Router Group
// @Description Router group for orders carts-related endpoints
// @Tags v1 orders carts
// @Accept json
// @Produce json
// @Router /api/v1/orders/carts [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pborder.OrderClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new carts controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pborder.OrderClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the carts controller
	route := baseRoute.Group(pbconfigrestcarts.Base.String())

	// Create a new carts controller
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
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestcarts.GetCartsMapper, c.getCarts))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestcarts.GetCartMapper, c.getCart))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestcarts.GetCartTotalMapper, c.getCartTotal))

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	currentController := moduleorderscurrent.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		currentController,
	} {
		controller.Initialize()
	}
}

// getCart gets a cart by ID
// @Summary Get a cart by ID
// @Description Get a cart by ID
// @Tags v1 orders carts
// @Accept json
// @Produce json
// @Param cartId path string true "Cart ID"
// @Success 200 {object} pborder.GetCartResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders/carts/{cartId} [get]
func (c *Controller) getCart(ctx *gin.Context) {
	var request pborder.GetCartRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the cart ID from the path
	request.CartId = ctx.Param(typesrest.CartId.String())

	// Get the cart by ID
	response, err := c.client.GetCart(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getCarts gets all carts
// @Summary Get all carts
// @Description Get all carts
// @Tags v1 orders carts
// @Accept json
// @Produce json
// @Success 200 {object} pborder.GetCartsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders/carts [get]
func (c *Controller) getCarts(ctx *gin.Context) {
	var request pborder.GetCartsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get all carts
	response, err := c.client.GetCarts(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getCartTotal gets the total of a cart by ID
// @Summary Get the total of a cart by ID
// @Description Get the total of a cart by ID
// @Tags v1 orders carts
// @Accept json
// @Produce json
// @Param cartId path string true "Cart ID"
// @Success 200 {object} pborder.GetCartTotalResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders/carts/{cartId}/total [get]
func (c *Controller) getCartTotal(ctx *gin.Context) {
	var request pborder.GetCartTotalRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the cart ID from the path
	request.CartId = ctx.Param(typesrest.CartId.String())

	// Get the total of the cart by ID
	response, err := c.client.GetCartTotal(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
