package revisions

import (
	"github.com/gin-gonic/gin"
	moduleshopsbusinesses "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/revisions/businesses"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfiggrpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	pbconfigrestrevisions "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/revisions"
	"net/http"
)

// Controller struct for the shops revisions module
// @Summary Shops Revisions Router Group
// @Description Router group for shops revisions-related endpoints
// @Tags v1 shops revisions
// @Accept json
// @Produce json
// @Router /api/v1/shops/revisions [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	authentication  authmiddleware.Authentication
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new revisions controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	authentication authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the revisions controller
	route := baseRoute.Group(pbconfigrestrevisions.Base.String())

	// Create the route handler
	routeHandler := commonhandler.NewDefaultHandler(authentication, &pbconfiggrpcshop.Interceptions)

	// Create a new revisions controller
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
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrevisions.UpdateAdminRevisionMapper,
			c.updateAdminRevision,
		),
	)
	c.route.DELETE(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrevisions.CloseAdminRevisionMapper,
			c.closeAdminRevision,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrevisions.OpenAdminRevisionToBranchMapper,
			c.openAdminRevisionToBranch,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrevisions.OpenAdminRevisionToProductMapper,
			c.openAdminRevisionToProduct,
		),
	)

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	businessesController := moduleshopsbusinesses.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		businessesController,
	} {
		controller.Initialize()
	}
}

// updateAdminRevision updates an admin revision
// @Summary Update an admin revision
// @Description Update an admin revision
// @Tags v1 shops revisions
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateAdminRevisionRequest true "Update Admin Revision Request"
// @Success 200 {object} pbshop.UpdateAdminRevisionResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/revisions [post]
func (c *Controller) updateAdminRevision(ctx *gin.Context) {
	var request pbshop.UpdateAdminRevisionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the admin revision
	response, err := c.client.UpdateAdminRevision(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// closeAdminRevision closes an admin revision
// @Summary Close an admin revision
// @Description Close an admin revision
// @Tags v1 shops revisions
// @Accept json
// @Produce json
// @Param request body pbshop.CloseAdminRevisionRequest true "Close Admin Revision Request"
// @Success 200 {object} pbshop.CloseAdminRevisionResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/revisions [delete]
func (c *Controller) closeAdminRevision(ctx *gin.Context) {
	var request pbshop.CloseAdminRevisionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Close the admin revision
	response, err := c.client.CloseAdminRevision(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// openAdminRevisionToBranch opens an admin revision to a branch
// @Summary Open an admin revision to a branch
// @Description Open an admin revision to a branch
// @Tags v1 shops revisions
// @Accept json
// @Produce json
// @Param request body pbshop.OpenAdminRevisionToBranchRequest true "Open Admin Revision To Branch Request"
// @Success 200 {object} pbshop.OpenAdminRevisionToBranchResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/revisions/branches [post]
func (c *Controller) openAdminRevisionToBranch(ctx *gin.Context) {
	var request pbshop.OpenAdminRevisionToBranchRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Open an admin revision to a branch
	response, err := c.client.OpenAdminRevisionToBranch(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// openAdminRevisionToProduct opens an admin revision to a product
// @Summary Open an admin revision to a product
// @Description Open an admin revision to a product
// @Tags v1 shops revisions
// @Accept json
// @Produce json
// @Param request body pbshop.OpenAdminRevisionToProductRequest true "Open Admin Revision To Product Request"
// @Success 200 {object} pbshop.OpenAdminRevisionToProductResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/revisions/products [post]
func (c *Controller) openAdminRevisionToProduct(ctx *gin.Context) {
	var request pbshop.OpenAdminRevisionToProductRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Open an admin revision to a product
	response, err := c.client.OpenAdminRevisionToProduct(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
