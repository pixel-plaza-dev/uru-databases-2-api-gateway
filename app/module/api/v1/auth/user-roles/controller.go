package user_roles

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfigrestuserroles "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/user-roles"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the user roles module
// @Summary Auth User Roles Router Group
// @Description Router group for auth user roles-related endpoints
// @Tags v1 auth user-roles
// @Accept json
// @Produce json
// @Router /api/v1/auth/user-roles [group]
type Controller struct {
	route           *gin.RouterGroup
	service         *appgrpcauth.Service
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new user roles controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcauth.Service,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the user roles controller
	route := baseRoute.Group(pbconfigrestuserroles.Base.String())

	// Create a new user roles controller
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
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestuserroles.AddUserRoleMapper, c.addUserRole))
	c.route.DELETE(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestuserroles.RevokeUserRoleMapper,
			c.revokeUserRole,
		),
	)
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestuserroles.GetUserRolesMapper, c.getUserRoles))
}

// addUserRole adds a role to a user
// @Summary Add a role to a user
// @Description Add a new role to a user by their ID
// @Tags v1 auth user-roles
// @Accept json
// @Produce json
// @Param user-id path string true "User ID"
// @Param request body pbauth.AddUserRoleRequest true "Add User Role Request"
// @Success 201 {object} pbauth.AddUserRoleResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/user-roles/{user-id} [post]
func (c *Controller) addUserRole(ctx *gin.Context) {
	var request pbauth.AddUserRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the user ID to the request
	request.UserId = ctx.Param(pbtypesrest.UserId.String())

	// Add a role to a user
	response, err := c.service.AddUserRole(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusCreated, response, err)
}

// revokeUserRole revokes a role from a user
// @Summary Revoke a role from a user
// @Description Revoke a specific role from a user by their ID
// @Tags v1 auth user-roles
// @Accept json
// @Produce json
// @Param user-id path string true "User ID"
// @Success 200 {object} pbauth.RevokeUserRoleResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/user-roles/{user-id} [delete]
func (c *Controller) revokeUserRole(ctx *gin.Context) {
	var request pbauth.RevokeUserRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the user ID to the request
	request.UserId = ctx.Param(pbtypesrest.UserId.String())

	// Revoke a role from the user
	response, err := c.service.RevokeUserRole(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getUserRoles gets all user's roles
// @Summary Get all user's roles
// @Description Get information about all roles for a specific user by their ID
// @Tags v1 auth user-roles
// @Accept json
// @Produce json
// @Param user-id path string true "User ID"
// @Success 200 {object} pbauth.GetUserRolesResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/user-roles/{user-id} [get]
func (c *Controller) getUserRoles(ctx *gin.Context) {
	var request pbauth.GetUserRolesRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the user ID to the request
	request.UserId = ctx.Param(pbtypesrest.UserId.String())

	// Get all user's roles
	response, err := c.service.GetUserRoles(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
