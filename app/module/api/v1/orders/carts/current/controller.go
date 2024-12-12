package carts

import (
	"github.com/gin-gonic/gin"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pborder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/order"
	pbconfigrestcarts "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/orders/carts"
	pbconfigrestcurrentcart "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/orders/carts/current"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

// Controller struct for the orders current cart module
// @Summary Orders Current Cart Router Group
// @Description Router group for orders current cart-related endpoints
// @Tags v1 orders carts current-cart
// @Accept json
// @Produce json
// @Router /api/v1/orders/carts/current [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pborder.OrderClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new current cart controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pborder.OrderClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the current cart controller
	route := baseRoute.Group(pbconfigrestcurrentcart.Base.String())

	// Create a new current cart controller
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
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestcurrentcart.GetCurrentCart, c.getCurrentCart))
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestcurrentcart.AddProductToCartMapper,
			c.addProductToCart,
		),
	)
	c.route.DELETE(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestcurrentcart.RemoveProductFromCartMapper,
			c.removeProductFromCart,
		),
	)
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestcurrentcart.PlaceOrderMapper, c.placeOrder))
}

// getCurrentCart gets the current cart
// @Summary Get the current cart
// @Description Get the current cart
// @Tags v1 orders carts current-cart
// @Accept json
// @Produce json
// @Success 200 {object} pborder.GetCurrentCartResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders/carts/current [get]
func (c *Controller) getCurrentCart(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the current cart
	response, err := c.client.GetCurrentCart(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// addProductToCart adds a product to the current cart
// @Summary Add a product to the current cart
// @Description Add a product to the current cart
// @Tags v1 orders carts current-cart
// @Accept json
// @Produce json
// @Param request body pborder.AddProductToCartRequest true "Add Product To Cart Request"
// @Success 200 {object} pborder.AddProductToCartResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders/carts/current [post]
func (c *Controller) addProductToCart(ctx *gin.Context) {
	var request pborder.AddProductToCartRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add product to the current cart
	response, err := c.client.AddProductToCart(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// removeProductFromCart removes a product from the current cart
// @Summary Remove a product from the current cart
// @Description Remove a product from the current cart
// @Tags v1 orders carts current-cart
// @Accept json
// @Produce json
// @Param cartId path string true "Cart ID"
// @Param productId path string true "Product ID"
// @Success 200 {object} pborder.RemoveProductFromCartResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders/carts/current/{branchProductId} [delete]
func (c *Controller) removeProductFromCart(ctx *gin.Context) {
	var request pborder.RemoveProductFromCartRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the product ID from the path
	request.BranchProductId = ctx.Param(typesrest.ProductId.String())

	// Remove product from the current cart
	response, err := c.client.RemoveProductFromCart(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// placeOrder places an order for the current cart
// @Summary Place an order for the current cart
// @Description Place an order for the current cart
// @Tags v1 orders carts current-cart
// @Accept json
// @Produce json
// @Param request body pborder.PlaceOrderRequest true "Place Order Request"
// @Success 200 {object} pborder.PlaceOrderResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/orders/carts/current/checkout [post]
func (c *Controller) placeOrder(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Place order for the current cart
	response, err := c.client.PlaceOrder(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
