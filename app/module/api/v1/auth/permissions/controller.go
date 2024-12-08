package permissions

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfigrestpermissions "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/permissions"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the permissions module
// @Summary Auth Permissions Router Group
// @Description Router group for auth permissions-related endpoints
// @Tags v1 auth permissions
// @Accept json
// @Produce json
// @Router /api/v1/auth/permissions [group]
type Controller struct {
	route           *gin.RouterGroup
	service         *appgrpcauth.Service
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new permissions controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcauth.Service,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the permissions controller
	route := baseRoute.Group(pbconfigrestpermissions.Base.String())

	// Create a new permissions controller
	return &Controller{
		route:           route,
		service:         service,
		routeHandler:    routeHandler,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestpermissions.AddPermissionMapper,
			c.addPermission,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestpermissions.GetPermissionsMapper,
			c.getPermissions,
		),
	)
	c.route.DELETE(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestpermissions.RevokePermissionMapper,
			c.revokePermission,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestpermissions.GetPermissionMapper,
			c.getPermission,
		),
	)
}

// addPermission adds a permission
// @Summary Add a permission
// @Description Add a new permission
// @Tags v1 auth permissions
// @Accept json
// @Produce json
// @Param request body pbauth.AddPermissionRequest true "Add Permission Request"
// @Success 201 {object} pbauth.AddPermissionResponse
// @Failure 400 {object} commongintypes.ErrorResponse
// @Failure 500 {object} commongintypes.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/permissions/ [post]
func (c *Controller) addPermission(ctx *gin.Context) {
	var request pbauth.AddPermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request, c.responseHandler)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewErrorResponse(err),
		)
		return
	}

	// Add a permission
	response, err := c.service.AddPermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// getPermissions gets all permissions
// @Summary Get all permissions
// @Description Get information about all permissions
// @Tags v1 auth permissions
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.GetPermissionsResponse
// @Failure 400 {object} commongintypes.ErrorResponse
// @Failure 500 {object} commongintypes.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/permissions/ [get]
func (c *Controller) getPermissions(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil, c.responseHandler)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewErrorResponse(err),
		)
		return
	}

	// Get all permissions
	response, err := c.service.GetPermissions(ctx, grpcCtx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// revokePermission revokes a permission
// @Summary Revoke a permission
// @Description Revoke a specific permission by its ID
// @Tags v1 auth permissions
// @Accept json
// @Produce json
// @Param permission-id path string true "Permission ID"
// @Success 200 {object} pbauth.RevokePermissionResponse
// @Failure 400 {object} commongintypes.ErrorResponse
// @Failure 500 {object} commongintypes.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/permissions/{permission-id} [delete]
func (c *Controller) revokePermission(ctx *gin.Context) {
	var request pbauth.RevokePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request, c.responseHandler)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewErrorResponse(err),
		)
		return
	}

	// Add the permission ID to the request
	request.PermissionId = ctx.Param(pbtypesrest.PermissionId.String())

	// Revoke a permission
	response, err := c.service.RevokePermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getPermission gets a permission
// @Summary Get a permission
// @Description Get information about a specific permission by its ID
// @Tags v1 auth permissions
// @Accept json
// @Produce json
// @Param permission-id path string true "Permission ID"
// @Success 200 {object} pbauth.GetPermissionResponse
// @Failure 400 {object} commongintypes.ErrorResponse
// @Failure 500 {object} commongintypes.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/permissions/{permission-id} [get]
func (c *Controller) getPermission(ctx *gin.Context) {
	var request pbauth.GetPermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request, c.responseHandler)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewErrorResponse(err),
		)
		return
	}

	// Add the permission ID to the request
	request.PermissionId = ctx.Param(pbtypesrest.PermissionId.String())

	// Get the permission
	response, err := c.service.GetPermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
