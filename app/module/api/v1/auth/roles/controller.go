package roles

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfigrestroles "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/roles"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"google.golang.org/protobuf/types/known/emptypb"
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
	route           *gin.RouterGroup
	client          pbauth.AuthClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new roles controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbauth.AuthClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the roles controller
	route := baseRoute.Group(pbconfigrestroles.Base.String())

	// Create a new roles controller
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
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestroles.AddRoleMapper, c.addRole))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestroles.GetRolesMapper, c.getRoles))
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestroles.AddRolePermissionMapper,
			c.addRolePermission,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestroles.GetRolePermissionsMapper,
			c.getRolePermissions,
		),
	)
	c.route.DELETE(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestroles.RevokeRoleMapper, c.revokeRole))
}

// addRole adds a role
// @Summary Add a role
// @Description Add a new role
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Param request body pbauth.AddRoleRequest true "Add Role Request"
// @Success 201 {object} pbauth.AddRoleResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/roles/ [post]
func (c *Controller) addRole(ctx *gin.Context) {
	var request pbauth.AddRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add a role
	response, err := c.client.AddRole(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getRoles gets all roles
// @Summary Get all roles
// @Description Get information about all roles
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.GetRolesResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/roles/ [get]
func (c *Controller) getRoles(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get all roles
	response, err := c.client.GetRoles(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
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
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/roles/{role-id} [post]
func (c *Controller) addRolePermission(ctx *gin.Context) {
	var request pbauth.AddRolePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypesrest.RoleId.String())

	// Add a permission to the role
	response, err := c.client.AddRolePermission(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// getRolePermissions gets all permissions for a role
// @Summary Get all permissions for a role
// @Description Get information about all permissions for a specific role by its ID
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Param role-id path string true "Role ID"
// @Success 200 {object} pbauth.GetRolePermissionsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/roles/{role-id} [get]
func (c *Controller) getRolePermissions(ctx *gin.Context) {
	var request pbauth.GetRolePermissionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypesrest.RoleId.String())

	// Get all permissions for the role
	response, err := c.client.GetRolePermissions(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// revokeRole revokes a role
// @Summary Revoke a role
// @Description Revoke a specific role by its ID
// @Tags v1 auth roles
// @Accept json
// @Produce json
// @Param role-id path string true "Role ID"
// @Success 200 {object} pbauth.RevokeRoleResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/roles/{role-id} [delete]
func (c *Controller) revokeRole(ctx *gin.Context) {
	var request pbauth.RevokeRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypesrest.RoleId.String())

	// Revoke a role
	response, err := c.client.RevokeRole(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
