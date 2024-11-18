package auth

import (
	"github.com/gin-gonic/gin"
	jwtmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/jwt"
	commongrpcctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/grpc/context"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/auth"
	"net/http"
)

type Controller struct {
	apiRoute      *gin.RouterGroup
	route         *gin.RouterGroup
	service       *Service
	jwtMiddleware jwtmiddleware.Authentication
}

// NewController creates a new controller
func NewController(
	apiRoute *gin.RouterGroup, service *Service,
	jwtMiddleware jwtmiddleware.Authentication,
) *Controller {
	// Create a new route for the auth service
	route := apiRoute.Group("/auth")

	// Create a new user controller
	instance := &Controller{
		apiRoute: apiRoute, route: route, service: service,
		jwtMiddleware: jwtMiddleware,
	}

	// Initialize the routes for the controller
	instance.initializeRoutes()

	return instance
}

// initializeRoutes initializes the routes for the controller
func (c *Controller) initializeRoutes() {
	c.route.POST("/log-in", c.logIn)
	c.route.POST(
		"/refresh-token", c.jwtMiddleware.Authenticate(),
		c.refreshToken,
	)
	c.route.POST(
		"/log-out", c.jwtMiddleware.Authenticate(),
		c.logOut,
	)
	c.route.GET(
		"/sessions", c.jwtMiddleware.Authenticate(),
		c.getSessions,
	)
	c.route.POST(
		"/close-sessions", c.jwtMiddleware.Authenticate(),
		c.closeSessions,
	)
	c.route.POST(
		"/permission", c.jwtMiddleware.Authenticate(),
		c.addPermission,
	)
	c.route.DELETE(
		"/permission", c.jwtMiddleware.Authenticate(),
		c.revokePermission,
	)
	c.route.GET(
		"/permissions", c.jwtMiddleware.Authenticate(),
		c.getPermissions,
	)
	c.route.GET(
		"/permission/:permission_id", c.jwtMiddleware.Authenticate(),
		c.getPermission,
	)
	c.route.POST(
		"/role-permission", c.jwtMiddleware.Authenticate(),
		c.addRolePermission,
	)
	c.route.DELETE(
		"/role-permission", c.jwtMiddleware.Authenticate(),
		c.revokeRolePermission,
	)
	c.route.GET(
		"/role-permissions/:role_id", c.jwtMiddleware.Authenticate(),

		c.getRolePermissions,
	)
	c.route.POST(
		"/role", c.jwtMiddleware.Authenticate(),
		c.addRole,
	)
	c.route.DELETE(
		"/role", c.jwtMiddleware.Authenticate(),
		c.revokeRole,
	)
	c.route.GET(
		"/roles", c.jwtMiddleware.Authenticate(),
		c.getRoles,
	)
	c.route.POST(
		"/user-role/:user_id", c.jwtMiddleware.Authenticate(),
		c.addUserRole,
	)
	c.route.DELETE(
		"/user-role/:user_id", c.jwtMiddleware.Authenticate(),
		c.revokeUserRole,
	)
	c.route.GET(
		"/user-roles/:user_id", c.jwtMiddleware.Authenticate(),
		c.getUserRoles,
	)
}

// logIn logs in a user
func (c *Controller) logIn(ctx *gin.Context) {
	var request pbauth.LogInRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Log in the user
	response, err := c.service.LogIn(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// refreshToken refreshes a user's token
func (c *Controller) refreshToken(ctx *gin.Context) {
	var request pbauth.RefreshTokenRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Refresh the token
	response, err := c.service.RefreshToken(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// logOut logs out a user
func (c *Controller) logOut(ctx *gin.Context) {
	var request pbauth.LogOutRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Log out the user
	response, err := c.service.LogOut(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// closeSessions closes all sessions for a user
func (c *Controller) closeSessions(ctx *gin.Context) {
	var request pbauth.CloseSessionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Close all sessions for the user
	response, err := c.service.CloseSessions(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getSessions gets all sessions for a user
func (c *Controller) getSessions(ctx *gin.Context) {
	var request pbauth.GetSessionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get all sessions for the user
	response, err := c.service.GetSessions(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// addPermission adds a permission to a user
func (c *Controller) addPermission(ctx *gin.Context) {
	var request pbauth.AddPermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Add a permission to the user
	response, err := c.service.AddPermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// revokePermission revokes a permission from a user
func (c *Controller) revokePermission(ctx *gin.Context) {
	var request pbauth.RevokePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Revoke a permission from the user
	response, err := c.service.RevokePermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getPermission gets a permission
func (c *Controller) getPermission(ctx *gin.Context) {
	var request pbauth.GetPermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get the permission
	response, err := c.service.GetPermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getPermissions gets all permissions
func (c *Controller) getPermissions(ctx *gin.Context) {
	var request pbauth.GetPermissionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get all permissions
	response, err := c.service.GetPermissions(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// addRole adds a role to a user
func (c *Controller) addRole(ctx *gin.Context) {
	var request pbauth.AddRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Add a role to the user
	response, err := c.service.AddRole(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// revokeRole revokes a role from a user
func (c *Controller) revokeRole(ctx *gin.Context) {
	var request pbauth.RevokeRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Revoke a role from the user
	response, err := c.service.RevokeRole(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getRoles gets all roles
func (c *Controller) getRoles(ctx *gin.Context) {
	var request pbauth.GetRolesRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get all roles
	response, err := c.service.GetRoles(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// addRolePermission adds a permission to a role
func (c *Controller) addRolePermission(ctx *gin.Context) {
	var request pbauth.AddRolePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Add a permission to the role
	response, err := c.service.AddRolePermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// revokeRolePermission revokes a permission from a role
func (c *Controller) revokeRolePermission(ctx *gin.Context) {
	var request pbauth.RevokeRolePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Revoke a permission from the role
	response, err := c.service.RevokeRolePermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getRolePermissions gets all permissions for a role
func (c *Controller) getRolePermissions(ctx *gin.Context) {
	var request pbauth.GetRolePermissionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get all permissions for the role
	response, err := c.service.GetRolePermissions(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// addUserRole adds a role to a user
func (c *Controller) addUserRole(ctx *gin.Context) {
	var request pbauth.AddUserRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Add a role to the user
	response, err := c.service.AddUserRole(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// revokeUserRole revokes a role from a user
func (c *Controller) revokeUserRole(ctx *gin.Context) {
	var request pbauth.RevokeUserRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Revoke a role from the user
	response, err := c.service.RevokeUserRole(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getUserRoles gets all roles for a user
func (c *Controller) getUserRoles(ctx *gin.Context) {
	var request pbauth.GetUserRolesRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": commongrpc.InternalServerError.Error()})
		return
	}

	// Get all roles for the user
	response, err := c.service.GetUserRoles(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
