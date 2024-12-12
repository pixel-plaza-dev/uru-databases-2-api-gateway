package orders

import (
	"github.com/gin-gonic/gin"
	moduleorderscarts "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/orders/carts"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pborder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/order"
	pbconfiggrpcorder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/order"
	pbconfigrestorders "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/orders"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

// Controller struct for the orders module
// @Summary Orders Router Group
// @Description Router group for orders-related endpoints
// @Tags v1 orders
// @Accept json
// @Produce json
// @Router /api/v1/orders [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pborder.OrderClient
	authentication  authmiddleware.Authentication
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new orders controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pborder.OrderClient,
	authentication authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the orders controller
	route := baseRoute.Group(pbconfigrestorders.Base.String())

	// Create the route handler
	routeHandler := commonhandler.NewDefaultHandler(authentication, &pbconfiggrpcorder.Interceptions)

	// Create a new orders controller
	return &Controller{
		route:           route,
		client:          client,
		authentication:  authentication,
		routeHandler:    routeHandler,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestorders.GetOrderMapper, c.getOrder))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestorders.GetOrdersMapper, c.getOrders))

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	cartsController := moduleorderscarts.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		cartsController,
	} {
		controller.Initialize()
	}
}

// getOrder gets an order by ID
// @Summary Get an order by ID
// @Description Get an order by ID
// @Tags v1 orders
// @Accept json
// @Produce json
// @Param orderId path string true "Order ID"
// @Success 200 {object} pborder.GetOrderResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders/{orderId} [get]
func (c *Controller) getOrder(ctx *gin.Context) {
	var request pborder.GetOrderRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the order ID from the path
	request.OrderId = ctx.Param(typesrest.OrderId.String())

	// Get the order by ID
	response, err := c.client.GetOrder(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getOrders gets all orders
// @Summary Get all orders
// @Description Get all orders
// @Tags v1 orders
// @Accept json
// @Produce json
// @Success 200 {object} pborder.GetOrdersResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders [get]
func (c *Controller) getOrders(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get all orders
	response, err := c.client.GetOrders(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
