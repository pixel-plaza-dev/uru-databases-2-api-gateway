package businesses

import (
	"github.com/gin-gonic/gin"
	moduleshopsbranches "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/businesses/branches"
	moduleshopsclients "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/businesses/clients"
	moduleshopsmarkets "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/businesses/markets"
	moduleshopsowners "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/businesses/owners"
	moduleshopsproducts "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/businesses/products"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestbusinesses "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/businesses"
	typesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the businesses module
// @Summary Shops Businesses Router Group
// @Description Router group for shops businesses-related endpoints
// @Tags v1 shops businesses
// @Accept json
// @Produce json
// @Router /api/v1/shops/businesses [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	authMiddleware  authmiddleware.Authentication
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new businesses controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the businesses controller
	route := baseRoute.Group(pbconfigrestbusinesses.Base.String())

	// Create a new businesses controller
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
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestbusinesses.AddBusinessMapper, c.addBusiness))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestbusinesses.GetBusinessMapper, c.getBusiness))
	c.route.PUT(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbusinesses.UpdateBusinessMapper,
			c.updateBusiness,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbusinesses.SetBusinessProfilePictureMapper,
			c.setBusinessProfilePicture,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbusinesses.SuspendBusinessMapper,
			c.suspendBusiness,
		),
	)
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbusinesses.ActivateBusinessMapper,
			c.activateBusiness,
		),
	)
	c.route.DELETE(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestbusinesses.DeleteBusinessMapper,
			c.deleteBusiness,
		),
	)

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	marketsController := moduleshopsmarkets.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	productsController := moduleshopsproducts.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	branchesController := moduleshopsbranches.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	clientsController := moduleshopsclients.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	ownersController := moduleshopsowners.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		productsController,
		branchesController,
		clientsController,
		ownersController,
		marketsController,
	} {
		controller.Initialize()
	}
}

// addBusiness adds a new business
// @Summary Add a new business
// @Description Add a new business
// @Tags v1 shops businesses
// @Accept json
// @Produce json
// @Param request body pbshop.AddBusinessRequest true "Add Business Request"
// @Success 201 {object} pbshop.AddBusinessResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses [post]
func (c *Controller) addBusiness(ctx *gin.Context) {
	var request pbshop.AddBusinessRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a new business
	response, err := c.client.AddBusiness(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getBusiness gets a business by ID
// @Summary Get a business by ID
// @Description Get a business by ID
// @Tags v1 shops businesses
// @Accept json
// @Produce json
// @Param businessId path string true "Business ID"
// @Success 200 {object} pbshop.GetBusinessResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/shops/businesses/{businessId} [get]
func (c *Controller) getBusiness(ctx *gin.Context) {
	var request pbshop.GetBusinessRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the business ID from the path
	request.BusinessId = ctx.Param(typesrest.BusinessId.String())

	// Get the business by ID
	response, err := c.client.GetBusiness(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// updateBusiness updates a business
// @Summary Update a business
// @Description Update a business
// @Tags v1 shops businesses
// @Accept json
// @Produce json
// @Param request body pbshop.UpdateBusinessRequest true "Update Business Request"
// @Success 200 {object} pbshop.UpdateBusinessResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses [put]
func (c *Controller) updateBusiness(ctx *gin.Context) {
	var request pbshop.UpdateBusinessRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Update the business
	response, err := c.client.UpdateBusiness(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// setBusinessProfilePicture sets the profile picture of a business
// @Summary Set the profile picture of a business
// @Description Set the profile picture of a business
// @Tags v1 shops businesses
// @Accept json
// @Produce json
// @Param request body pbshop.SetBusinessProfilePictureRequest true "Set Business Profile Picture Request"
// @Success 200 {object} pbshop.SetBusinessProfilePictureResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/profile-picture [post]
func (c *Controller) setBusinessProfilePicture(ctx *gin.Context) {
	var request pbshop.SetBusinessProfilePictureRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Set the profile picture of the business
	response, err := c.client.SetBusinessProfilePicture(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// suspendBusiness suspends a business
// @Summary Suspend a business
// @Description Suspend a business
// @Tags v1 shops businesses
// @Accept json
// @Produce json
// @Param request body pbshop.SuspendBusinessRequest true "Suspend Business Request"
// @Success 200 {object} pbshop.SuspendBusinessResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/suspend [post]
func (c *Controller) suspendBusiness(ctx *gin.Context) {
	var request pbshop.SuspendBusinessRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Suspend the business
	response, err := c.client.SuspendBusiness(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// activateBusiness activates a business
// @Summary Activate a business
// @Description Activate a business
// @Tags v1 shops businesses
// @Accept json
// @Produce json
// @Param request body pbshop.ActivateBusinessRequest true "Activate Business Request"
// @Success 200 {object} pbshop.ActivateBusinessResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/activate [post]
func (c *Controller) activateBusiness(ctx *gin.Context) {
	var request pbshop.ActivateBusinessRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Activate the business
	response, err := c.client.ActivateBusiness(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// deleteBusiness deletes a business
// @Summary Delete a business
// @Description Delete a business
// @Tags v1 shops businesses
// @Accept json
// @Produce json
// @Param businessId path string true "Business ID"
// @Success 200 {object} pbshop.DeleteBusinessResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/shops/businesses/{businessId} [delete]
func (c *Controller) deleteBusiness(ctx *gin.Context) {
	var request pbshop.DeleteBusinessRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the business ID from the path
	request.BusinessId = ctx.Param(typesrest.BusinessId.String())

	// Delete the business
	response, err := c.client.DeleteBusiness(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
