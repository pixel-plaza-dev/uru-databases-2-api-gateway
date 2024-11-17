package auth

import (
	"github.com/gin-gonic/gin"
	jwtmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/jwt"
	mdmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/metadata"
)

type Controller struct {
	apiRoute      *gin.RouterGroup
	route         *gin.RouterGroup
	service       *Service
	jwtMiddleware jwtmiddleware.Authentication
	mdMiddleware  mdmiddleware.Authentication
}

// NewController creates a new controller
func NewController(
	apiRoute *gin.RouterGroup, service *Service,
	jwtMiddleware jwtmiddleware.Authentication,
	mdMiddleware mdmiddleware.Authentication,
) *Controller {
	// Create a new route for the auth service
	route := apiRoute.Group("/auth")

	// Create a new user controller
	instance := &Controller{
		apiRoute: apiRoute, route: route, service: service,
		jwtMiddleware: jwtMiddleware, mdMiddleware: mdMiddleware,
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
		c.mdMiddleware.Authenticate(), c.refreshToken,
	)
	c.route.POST(
		"/log-out", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.logOut,
	)
	c.route.GET(
		"/sessions", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.getSessions,
	)
	c.route.POST(
		"/close-sessions", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.closeSessions,
	)
	c.route.POST(
		"/permission", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.addPermission,
	)
	c.route.DELETE(
		"/permission", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.revokePermission,
	)
	c.route.GET(
		"/permissions", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.getPermissions,
	)
	c.route.GET(
		"/permission/:permission_id", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.getPermission,
	)
	c.route.POST(
		"/role-permission", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.addRolePermission,
	)
	c.route.DELETE(
		"/role-permission", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.revokeRolePermission,
	)
	c.route.GET(
		"/role-permissions/:role_id", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(),
		c.getRolePermissions,
	)
	c.route.POST(
		"/role", c.jwtMiddleware.Authenticate(), c.mdMiddleware.Authenticate(),
		c.addRole,
	)
	c.route.DELETE(
		"/role", c.jwtMiddleware.Authenticate(), c.mdMiddleware.Authenticate(),
		c.revokeRole,
	)
	c.route.GET(
		"/roles", c.jwtMiddleware.Authenticate(), c.mdMiddleware.Authenticate(),
		c.getRoles,
	)
	c.route.POST(
		"/user-role/:user_id", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.addUserRole,
	)
	c.route.DELETE(
		"/user-role/:user_id", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.revokeUserRole,
	)
	c.route.GET(
		"/user-roles/:user_id", c.jwtMiddleware.Authenticate(),
		c.mdMiddleware.Authenticate(), c.getUserRoles,
	)
}

// logIn logs in a user
func (c *Controller) logIn(ctx *gin.Context) {
	// Log in the user
	response, err := c.service.LogIn(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// refreshToken refreshes a user's token
func (c *Controller) refreshToken(ctx *gin.Context) {
	// Refresh the token
	response, err := c.service.RefreshToken(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// logOut logs out a user
func (c *Controller) logOut(ctx *gin.Context) {
	// Log out the user
	response, err := c.service.LogOut(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// closeSessions closes all sessions for a user
func (c *Controller) closeSessions(ctx *gin.Context) {
	// Close all sessions for the user
	response, err := c.service.CloseSessions(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// getSessions gets all sessions for a user
func (c *Controller) getSessions(ctx *gin.Context) {
	// Get all sessions for the user
	response, err := c.service.GetSessions(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// addPermission adds a permission to a user
func (c *Controller) addPermission(ctx *gin.Context) {
	// Add a permission to the user
	response, err := c.service.AddPermission(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, response)
}

// revokePermission revokes a permission from a user
func (c *Controller) revokePermission(ctx *gin.Context) {
	// Revoke a permission from the user
	response, err := c.service.RevokePermission(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// getPermission gets a permission
func (c *Controller) getPermission(ctx *gin.Context) {
	// Get the permission
	response, err := c.service.GetPermission(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// getPermissions gets all permissions
func (c *Controller) getPermissions(ctx *gin.Context) {
	// Get all permissions
	response, err := c.service.GetPermissions(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// addRole adds a role to a user
func (c *Controller) addRole(ctx *gin.Context) {
	// Add a role to the user
	response, err := c.service.AddRole(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, response)
}

// revokeRole revokes a role from a user
func (c *Controller) revokeRole(ctx *gin.Context) {
	// Revoke a role from the user
	response, err := c.service.RevokeRole(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// getRoles gets all roles
func (c *Controller) getRoles(ctx *gin.Context) {
	// Get all roles
	response, err := c.service.GetRoles(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// addRolePermission adds a permission to a role
func (c *Controller) addRolePermission(ctx *gin.Context) {
	// Add a permission to the role
	response, err := c.service.AddRolePermission(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, response)
}

// revokeRolePermission revokes a permission from a role
func (c *Controller) revokeRolePermission(ctx *gin.Context) {
	// Revoke a permission from the role
	response, err := c.service.RevokeRolePermission(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// getRolePermissions gets all permissions for a role
func (c *Controller) getRolePermissions(ctx *gin.Context) {
	// Get all permissions for the role
	response, err := c.service.GetRolePermissions(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// addUserRole adds a role to a user
func (c *Controller) addUserRole(ctx *gin.Context) {
	// Add a role to the user
	response, err := c.service.AddUserRole(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, response)
}

// revokeUserRole revokes a role from a user
func (c *Controller) revokeUserRole(ctx *gin.Context) {
	// Revoke a role from the user
	response, err := c.service.RevokeUserRole(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}

// getUserRoles gets all roles for a user
func (c *Controller) getUserRoles(ctx *gin.Context) {
	// Get all roles for the user
	response, err := c.service.GetUserRoles(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}
