package roles

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfigrestauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/v1/auth"
	pbconfigrestroles "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/v1/auth/roles"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the roles module
// @Summary Auth Roles Router Group
// @Description Router group for auth roles-related endpoints
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Router /api/v1/auth/roles [group]
type Controller struct {
	route   *gin.RouterGroup
	service *appgrpcauth.Service
}

// NewController creates a new roles controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcauth.Service,
) *Controller {
	// Create a new route for the roles controller
	route := baseRoute.Group(pbconfigrestauth.Roles.String())

	// Create a new roles controller
	return &Controller{
		route:   route,
		service: service,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.POST(pbconfigrestroles.Relative.String(), c.addRole)
	c.route.GET(pbconfigrestroles.Relative.String(), c.getRoles)
	c.route.POST(pbconfigrestroles.ByRoleId.String(), c.addRolePermission)
	c.route.GET(pbconfigrestroles.ByRoleId.String(), c.getRolePermissions)
	c.route.DELETE(pbconfigrestroles.ByRoleId.String(), c.revokeRole)
}

// addRole adds a role
// @Summary Add a role
// @Description Add a new role
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Param request body pbauth.AddRoleRequest true "Add Role Request"
// @Success 201 {object} pbauth.AddRoleResponse
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/roles/ [post]
func (c *Controller) addRole(ctx *gin.Context) {
	var request pbauth.AddRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Add a role
	response, err := c.service.AddRole(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// getRoles gets all roles
// @Summary Get all roles
// @Description Get information about all roles
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.GetRolesResponse
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/roles/ [get]
func (c *Controller) getRoles(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Get all roles
	response, err := c.service.GetRoles(ctx, grpcCtx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// addRolePermission adds a permission to a role
// @Summary Add a permission to a role
// @Description Add a new permission to a role by its ID
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Param role-id path string true "Role ID"
// @Param request body pbauth.AddRolePermissionRequest true "Add Role Permission Request"
// @Success 201 {object} pbauth.AddRolePermissionResponse
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/roles/{role-id} [post]
func (c *Controller) addRolePermission(ctx *gin.Context) {
	var request pbauth.AddRolePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypesrest.RoleId.String())

	// Add a permission to the role
	response, err := c.service.AddRolePermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// getRolePermissions gets all permissions for a role
// @Summary Get all permissions for a role
// @Description Get information about all permissions for a specific role by its ID
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Param role-id path string true "Role ID"
// @Success 200 {object} pbauth.GetRolePermissionsResponse
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/roles/{role-id} [get]
func (c *Controller) getRolePermissions(ctx *gin.Context) {
	var request pbauth.GetRolePermissionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypesrest.RoleId.String())

	// Get all permissions for the role
	response, err := c.service.GetRolePermissions(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// revokeRole revokes a role
// @Summary Revoke a role
// @Description Revoke a specific role by its ID
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Param role-id path string true "Role ID"
// @Success 200 {object} pbauth.RevokeRoleResponse
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/roles/{role-id} [delete]
func (c *Controller) revokeRole(ctx *gin.Context) {
	var request pbauth.RevokeRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypesrest.RoleId.String())

	// Revoke a role
	response, err := c.service.RevokeRole(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
