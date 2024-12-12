package stores

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigreststores "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/stores"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

// Controller struct for the stores module
// @Summary Shops Stores Router Group
// @Description Router group for shops stores-related endpoints
// @Tags v1 shops stores
// @Accept json
// @Produce json
// @Router /api/v1/shops/stores [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new stores controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the stores controller
	route := baseRoute.Group(pbconfigreststores.Base.String())

	// Create a new stores controller
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
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigreststores.AddStoreMapper, c.addStore))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigreststores.GetStoreMapper, c.getStore))
	c.route.DELETE(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigreststores.DeleteStoreMapper, c.deleteStore))
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigreststores.GetUnoccupiedStoresMapper,
			c.getUnoccupiedStores,
		),
	)
}

// addStore adds a new store
// @Summary Add a new store
// @Description Add a new store
// @Tags v1 shops stores
// @Accept json
// @Produce json
// @Param request body pbshop.AddStoreRequest true "Add Store Request"
// @Success 201 {object} pbshop.AddStoreResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores [post]
func (c *Controller) addStore(ctx *gin.Context) {
	var request pbshop.AddStoreRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new store
	response, err := c.client.AddStore(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getStore gets a store by ID
// @Summary Get a store by ID
// @Description Get a store by ID
// @Tags v1 shops stores
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} pbshop.GetStoreResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores/{storeId} [get]
func (c *Controller) getStore(ctx *gin.Context) {
	var request pbshop.GetStoreRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the store ID from the path
	request.StoreId = ctx.Param(typesrest.StoreId.String())

	// Get the store by ID
	response, err := c.client.GetStore(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// deleteStore deletes a store
// @Summary Delete a store
// @Description Delete a store
// @Tags v1 shops stores
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} pbshop.DeleteStoreResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores/{storeId} [delete]
func (c *Controller) deleteStore(ctx *gin.Context) {
	var request pbshop.DeleteStoreRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the store ID from the path
	request.StoreId = ctx.Param(typesrest.StoreId.String())

	// Delete the store
	response, err := c.client.DeleteStore(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getUnoccupiedStores gets unoccupied stores
// @Summary Get unoccupied stores
// @Description Get unoccupied stores
// @Tags v1 shops stores
// @Accept json
// @Produce json
// @Success 200 {object} pbshop.GetUnoccupiedStoresResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/stores/unoccupied [get]
func (c *Controller) getUnoccupiedStores(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get unoccupied stores
	response, err := c.client.GetUnoccupiedStores(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
