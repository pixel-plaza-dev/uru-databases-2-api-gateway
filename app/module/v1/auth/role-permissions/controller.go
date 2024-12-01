package role_permissions

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
	pbconfigrestauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/auth"
	pbconfigrestrolepermissions "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/auth/role-permissions"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
	"net/http"
)

// Controller struct for the role-permissions module
// @Summary Auth Role Permissions Router Group
// @Description Router group for auth role-permissions-related endpoints
// @Tags v1 auth role-permissions
// @Accept json
// @Produce json
// @Router /role-permissions [group]
type Controller struct {
	route   *gin.RouterGroup
	service *appgrpcauth.Service
}

// NewController creates a new role-permissions controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcauth.Service,
) *Controller {
	// Create a new route for the role-permissions controller
	route := baseRoute.Group(pbconfigrestauth.RolePermissions.String())

	// Create a new role-permissions controller
	return &Controller{
		route:   route,
		service: service,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.DELETE(
		pbconfigrestrolepermissions.ById.String(),
		c.revokeRolePermission,
	)
}

// revokeRolePermission revokes a permission from a role
// @Summary Revoke a permission from a role
// @Description Revoke a specific permission from a role by its ID
// @Tags auth
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} pbauth.RevokeRolePermissionResponse
// @Failure 400 {object} module.BadRequest
// @Failure 500 {object} commongin.InternalServerError
// @Router /{id} [delete]
func (c *Controller) revokeRolePermission(ctx *gin.Context) {
	var request pbauth.RevokeRolePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
		)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypesrest.Id.String())

	// Revoke a permission from the role
	response, err := c.service.RevokeRolePermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
