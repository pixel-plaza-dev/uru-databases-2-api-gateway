package auth

import (
	"github.com/gin-gonic/gin"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
	pbauthapi "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/auth/api"
	pbtypes "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	"net/http"
)

type Controller struct {
	apiRoute *gin.RouterGroup
	route    *gin.RouterGroup
	service  *Service
	restMap  *map[string]map[pbtypes.
			RESTMethod]pbtypes.GRPCMethod
	grpcInterceptions *map[pbtypes.GRPCMethod]pbtypes.Interception
	authMiddleware    authmiddleware.Authentication
}

// NewController creates a new controller
func NewController(
	apiRoute *gin.RouterGroup, service *Service,
	jwtValidator commonjwtvalidator.Validator,
	restMap *map[string]map[pbtypes.RESTMethod]pbtypes.
		GRPCMethod,
	grpcInterceptions *map[pbtypes.GRPCMethod]pbtypes.Interception,
	authLogger authmiddleware.Logger,
) (*Controller, error) {
	// Create a new route for the user service
	route := apiRoute.Group(RelativeUri)

	// Create the auth middleware
	authMiddleware, err := authmiddleware.NewMiddleware(
		route.BasePath(),
		jwtValidator,
		restMap,
		grpcInterceptions,
		authLogger,
	)
	if err != nil {
		return nil, err
	}

	// Add the auth middleware to the route
	route.Use(authMiddleware.Authenticate())

	// Create a new user controller
	instance := &Controller{
		apiRoute: apiRoute, route: route, service: service,
		authMiddleware: authMiddleware,
	}

	// Initialize the routes for the controller
	instance.initializeRoutes()

	return instance, nil
}

// initializeRoutes initializes the routes for the controller
func (c *Controller) initializeRoutes() {
	c.route.POST(pbauthapi.LogIn.String(), c.logIn)
	c.route.GET(pbauthapi.AccessTokenByToken.String(), c.isAccessTokenValid)
	c.route.GET(pbauthapi.RefreshTokenByToken.String(), c.isRefreshTokenValid)
	c.route.POST(pbauthapi.RefreshToken.String(), c.refreshToken)
	c.route.POST(pbauthapi.LogOut.String(), c.logOut)
	c.route.GET(pbauthapi.Sessions.String(), c.getSessions)
	c.route.DELETE(pbauthapi.Sessions.String(), c.closeSessions)
	c.route.DELETE(pbauthapi.SessionByToken.String(), c.closeSession)
	c.route.POST(pbauthapi.Permission.String(), c.addPermission)
	c.route.DELETE(pbauthapi.PermissionById.String(), c.revokePermission)
	c.route.GET(pbauthapi.PermissionById.String(), c.getPermission)
	c.route.GET(pbauthapi.Permissions.String(), c.getPermissions)
	c.route.DELETE(
		pbauthapi.RolePermissionById.String(),
		c.revokeRolePermission,
	)
	c.route.POST(pbauthapi.Role.String(), c.addRole)
	c.route.POST(pbauthapi.RoleById.String(), c.addRolePermission)
	c.route.GET(pbauthapi.RoleById.String(), c.getRolePermissions)
	c.route.DELETE(pbauthapi.RoleById.String(), c.revokeRole)
	c.route.GET(pbauthapi.Roles.String(), c.getRoles)
	c.route.POST(pbauthapi.UserRoleByUserId.String(), c.addUserRole)
	c.route.DELETE(pbauthapi.UserRoleByUserId.String(), c.revokeUserRole)
	c.route.GET(pbauthapi.UserRolesByUserId.String(), c.getUserRoles)
}

// logIn logs in a user
func (c *Controller) logIn(ctx *gin.Context) {
	var request pbauth.LogInRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
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

// isAccessTokenValid checks if an access token is valid
func (c *Controller) isAccessTokenValid(ctx *gin.Context) {
	var request pbauth.IsAccessTokenValidRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the access token to the request
	request.AccessToken = ctx.Param(pbtypes.Token.String())

	// Check if the access token is valid
	response, err := c.service.IsAccessTokenValid(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// isRefreshTokenValid checks if a refresh token is valid
func (c *Controller) isRefreshTokenValid(ctx *gin.Context) {
	var request pbauth.IsRefreshTokenValidRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the refresh token to the request
	request.RefreshToken = ctx.Param(pbtypes.Token.String())

	// Check if the refresh token is valid
	response, err := c.service.IsRefreshTokenValid(ctx, grpcCtx, &request)
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
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
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
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
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

// getSessions gets all user' sessions
func (c *Controller) getSessions(ctx *gin.Context) {
	var request pbauth.GetSessionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Get all user' sessions
	response, err := c.service.GetSessions(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// closeSession closes a given user' session
func (c *Controller) closeSession(ctx *gin.Context) {
	var request pbauth.CloseSessionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the refresh token to the request
	request.RefreshToken = ctx.Param(pbtypes.Token.String())

	// Close a given user' session
	response, err := c.service.CloseSession(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// closeSessions closes all user' sessions
func (c *Controller) closeSessions(ctx *gin.Context) {
	var request pbauth.CloseSessionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Close all user' sessions
	response, err := c.service.CloseSessions(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// addPermission adds a permission
func (c *Controller) addPermission(ctx *gin.Context) {
	var request pbauth.AddPermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add a permission
	response, err := c.service.AddPermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// revokePermission revokes a permission
func (c *Controller) revokePermission(ctx *gin.Context) {
	var request pbauth.RevokePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the permission ID to the request
	request.PermissionId = ctx.Param(pbtypes.Id.String())

	// Revoke a permission
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
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the permission ID to the request
	request.PermissionId = ctx.Param(pbtypes.Id.String())

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
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
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

// revokeRolePermission revokes a permission from a role
func (c *Controller) revokeRolePermission(ctx *gin.Context) {
	var request pbauth.RevokeRolePermissionRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypes.Id.String())

	// Revoke a permission from the role
	response, err := c.service.RevokeRolePermission(ctx, grpcCtx, &request)
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
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypes.Id.String())

	// Add a permission to the role
	response, err := c.service.AddRolePermission(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// getRolePermissions gets all permissions for a role
func (c *Controller) getRolePermissions(ctx *gin.Context) {
	var request pbauth.GetRolePermissionsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypes.Id.String())

	// Get all permissions for the role
	response, err := c.service.GetRolePermissions(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// addRole adds a role
func (c *Controller) addRole(ctx *gin.Context) {
	var request pbauth.AddRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add a role
	response, err := c.service.AddRole(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// revokeRole revokes a role
func (c *Controller) revokeRole(ctx *gin.Context) {
	var request pbauth.RevokeRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the role ID to the request
	request.RoleId = ctx.Param(pbtypes.Id.String())

	// Revoke a role
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
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
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

// addUserRole adds a role to a user
func (c *Controller) addUserRole(ctx *gin.Context) {
	var request pbauth.AddUserRoleRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the user ID to the request
	request.UserId = ctx.Param(pbtypes.UserId.String())

	// Add a role to a user
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
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the user ID to the request
	request.UserId = ctx.Param(pbtypes.UserId.String())

	// Revoke a role from the user
	response, err := c.service.RevokeUserRole(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getUserRoles gets all user's roles
func (c *Controller) getUserRoles(ctx *gin.Context) {
	var request pbauth.GetUserRolesRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": commongrpc.InternalServerError.Error()},
		)
		return
	}

	// Add the user ID to the request
	request.UserId = ctx.Param(pbtypes.UserId.String())

	// Get all user's roles
	response, err := c.service.GetUserRoles(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
