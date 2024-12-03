package permissions

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfigrestauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/v1/auth"
	pbconfigrestpermissions "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/v1/auth/permissions"
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
	route   *gin.RouterGroup
	service *appgrpcauth.Service
}

// NewController creates a new permissions controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcauth.Service,
) *Controller {
	// Create a new route for the permissions controller
	route := baseRoute.Group(pbconfigrestauth.Permissions.String())

	// Create a new permissions controller
	return &Controller{
		route:   route,
		service: service,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.POST(pbconfigrestpermissions.Relative.String(), c.addPermission)
	c.route.GET(pbconfigrestpermissions.Relative.String(), c.getPermissions)
	c.route.DELETE(pbconfigrestpermissions.ByPermissionId.String(), c.revokePermission)
	c.route.GET(pbconfigrestpermissions.ByPermissionId.String(), c.getPermission)
}

// addPermission adds a permission
// @Summary Add a permission
// @Description Add a new permission
// @Tags v1 auth permissions
// @Accept json
// @Produce json
// @Param request body pbauth.AddPermissionRequest true "Add Permission Request"
// @Success 201 {object} pbauth.AddPermissionResponse
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/permissions/ [post]
func (c *Controller) addPermission(ctx *gin.Context) {
	var request pbauth.AddPermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Add a permission
	response, err := c.service.AddPermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
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
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/permissions/ [get]
func (c *Controller) getPermissions(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Get all permissions
	response, err := c.service.GetPermissions(ctx, grpcCtx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
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
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/permissions/{permission-id} [delete]
func (c *Controller) revokePermission(ctx *gin.Context) {
	var request pbauth.RevokePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Add the permission ID to the request
	request.PermissionId = ctx.Param(pbtypesrest.PermissionId.String())

	// Revoke a permission
	response, err := c.service.RevokePermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
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
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/permissions/{permission-id} [get]
func (c *Controller) getPermission(ctx *gin.Context) {
	var request pbauth.GetPermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Add the permission ID to the request
	request.PermissionId = ctx.Param(pbtypesrest.PermissionId.String())

	// Get the permission
	response, err := c.service.GetPermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
