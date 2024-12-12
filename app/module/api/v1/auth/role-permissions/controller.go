package role_permissions

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfigrestrolepermissions "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/role-permissions"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the role-permissions module
// @Summary Auth Role Permissions Router Group
// @Description Router group for auth role-permissions-related endpoints
// @Tags v1 auth role-permissions
// @Accept json
// @Produce json
// @Router /api/v1/auth/role-permissions [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbauth.AuthClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new role-permissions controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbauth.AuthClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the role-permissions controller
	route := baseRoute.Group(pbconfigrestrolepermissions.Base.String())

	// Create a new role-permissions controller
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
	c.route.DELETE(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrolepermissions.RevokeRolePermissionMapper,
			c.revokeRolePermission,
		),
	)
}

// revokeRolePermission revokes a permission from a role
// @Summary Revoke a permission from a role
// @Description Revoke a specific permission from a role by its ID
// @Tags v1 auth role-permissions
// @Accept json
// @Produce json
// @Param role-id path string true "Role ID"
// @Success 200 {object} pbauth.RevokeRolePermissionResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/role-permissions/{role-id} [delete]
func (c *Controller) revokeRolePermission(ctx *gin.Context) {
	var request pbauth.RevokeRolePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypesrest.RoleId.String())

	// Revoke a permission from the role
	response, err := c.client.RevokeRolePermission(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
