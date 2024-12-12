package businesses

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestbusinesses "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/revisions/businesses"
	"net/http"
)

// Controller struct for the businesses revisions module
// @Summary Shops Businesses Revisions Router Group
// @Description Router group for shops businesses revisions-related endpoints
// @Tags v1 shops revisions businesses
// @Accept json
// @Produce json
// @Router /api/v1/shops/revisions/businesses [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new businesses revisions controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the businesses revisions controller
	route := baseRoute.Group(pbconfigrestbusinesses.Base.String())

	// Create a new businesses revisions controller
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
			pbconfigrestbusinesses.OpenAdminRevisionToBusinessMapper,
			c.openAdminRevisionToBusiness,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbusinesses.OpenAdminRevisionToBusinessProductMapper,
			c.openAdminRevisionToBusinessProduct,
		),
	)
}

// openAdminRevisionToBusiness opens an admin revision to a business
// @Summary Open an admin revision to a business
// @Description Open an admin revision to a business
// @Tags v1 shops revisions businesses
// @Accept json
// @Produce json
// @Param request body pbshop.OpenAdminRevisionToBusinessRequest true "Open Admin Revision To Business Request"
// @Success 200 {object} pbshop.OpenAdminRevisionToBusinessResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/revisions/businesses [post]
func (c *Controller) openAdminRevisionToBusiness(ctx *gin.Context) {
	var request pbshop.OpenAdminRevisionToBusinessRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Open an admin revision to a business
	response, err := c.client.OpenAdminRevisionToBusiness(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// openAdminRevisionToBusinessProduct opens an admin revision to a business product
// @Summary Open an admin revision to a business product
// @Description Open an admin revision to a business product
// @Tags v1 shops revisions businesses
// @Accept json
// @Produce json
// @Param request body pbshop.OpenAdminRevisionToBusinessProductRequest true "Open Admin Revision To Business Product Request"
// @Success 200 {object} pbshop.OpenAdminRevisionToBusinessProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/revisions/businesses/products [get]
func (c *Controller) openAdminRevisionToBusinessProduct(ctx *gin.Context) {
	var request pbshop.OpenAdminRevisionToBusinessProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Open an admin revision to a business product
	response, err := c.client.OpenAdminRevisionToBusinessProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
