package owners

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestowners "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/businesses/owners"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the shops businesses owners module
// @Summary Shops Businesses Owners Router Group
// @Description Router group for shops businesses owners-related endpoints
// @Tags v1 shops businesses owners
// @Accept json
// @Produce json
// @Router /api/v1/shops/businesses/owners [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new owners controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the owners controller
	route := baseRoute.Group(pbconfigrestowners.Base.String())

	// Create a new owners controller
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
			pbconfigrestowners.AddBusinessOwnerMapper,
			c.addBusinessOwner,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestowners.RemoveBusinessOwnerMapper,
			c.removeBusinessOwner,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestowners.GetBusinessOwnersMapper,
			c.getBusinessOwners,
		),
	)
}

// addBusinessOwner adds a new business owner
// @Summary Add a new business owner
// @Description Add a new business owner
// @Tags v1 shops businesses owners
// @Accept json
// @Produce json
// @Param request body pbshop.AddBusinessOwnerRequest true "Add Business Owner Request"
// @Success 201 {object} pbshop.AddBusinessOwnerResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/owners [post]
func (c *Controller) addBusinessOwner(ctx *gin.Context) {
	var request pbshop.AddBusinessOwnerRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new business owner
	response, err := c.client.AddBusinessOwner(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// removeBusinessOwner removes a business owner
// @Summary Remove a business owner
// @Description Remove a business owner
// @Tags v1 shops businesses owners
// @Accept json
// @Produce json
// @Param request body pbshop.RemoveBusinessOwnerRequest true "Remove Business Owner Request"
// @Success 200 {object} pbshop.RemoveBusinessOwnerResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/owners [delete]
func (c *Controller) removeBusinessOwner(ctx *gin.Context) {
	var request pbshop.RemoveBusinessOwnerRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Remove the business owner
	response, err := c.client.RemoveBusinessOwner(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getBusinessOwners gets all business owners
// @Summary Get all business owners
// @Description Get all business owners
// @Tags v1 shops businesses owners
// @Accept json
// @Produce json
// @Param businessId path string true "Business ID"
// @Success 200 {object} pbshop.GetBusinessOwnersResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/owners/{businessId} [get]
func (c *Controller) getBusinessOwners(ctx *gin.Context) {
	var request pbshop.GetBusinessOwnersRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the business ID from the path
	request.BusinessId = ctx.Param(typesrest.BusinessId.String())

	// Get all business owners
	response, err := c.client.GetBusinessOwners(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
