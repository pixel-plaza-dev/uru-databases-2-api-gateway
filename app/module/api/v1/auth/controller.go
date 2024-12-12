package auth

import (
	"github.com/gin-gonic/gin"
	moduleauthaccesstokens "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/access-tokens"
	moduleauthpermissions "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/permissions"
	moduleauthrefreshtokens "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/refresh-tokens"
	moduleauthrolepermissions "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/role-permissions"
	moduleauthroles "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/roles"
	moduleauthuserroles "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/user-roles"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfiggrpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	pbconfigrestauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

// Controller struct for the auth module
// @Summary Auth Router Group
// @Description Router group for authentication-related endpoints
// @Tags v1 auth
// @Accept json
// @Produce json
// @Router /api/v1/auth [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbauth.AuthClient
	authentication  authmiddleware.Authentication
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new auth controller
func NewController(
	apiRoute *gin.RouterGroup,
	client pbauth.AuthClient,
	authentication authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the auth controller
	route := apiRoute.Group(pbconfigrestauth.Base.String())

	// Create the route handler
	routeHandler := commonhandler.NewDefaultHandler(authentication, &pbconfiggrpcauth.Interceptions)

	// Create a new auth controller
	return &Controller{
		route:           route,
		client:          client,
		authentication:  authentication,
		routeHandler:    routeHandler,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestauth.LogInMapper, c.logIn))
	c.route.POST(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestauth.LogOutMapper, c.logOut))

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	accessTokensController := moduleauthaccesstokens.NewController(
		c.route,
		c.client,
		c.routeHandler,
		c.responseHandler,
	)
	refreshTokensController := moduleauthrefreshtokens.NewController(
		c.route,
		c.client,
		c.routeHandler,
		c.responseHandler,
	)
	permissionsController := moduleauthpermissions.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	rolePermissionsController := moduleauthrolepermissions.NewController(
		c.route,
		c.client,
		c.routeHandler,
		c.responseHandler,
	)
	rolesController := moduleauthroles.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	userRolesController := moduleauthuserroles.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		accessTokensController,
		refreshTokensController,
		permissionsController,
		rolePermissionsController,
		rolesController,
		userRolesController,
	} {
		controller.Initialize()
	}
}

// logIn logs in a user
// @Summary Log in a user
// @Description Log in a user with their credentials
// @Tags v1 auth
// @Accept json
// @Produce json
// @Param request body LogInRequest true "Log In Request"
// @Success 200 {object} LogInResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/auth/log-in [post]
func (c *Controller) logIn(ctx *gin.Context) {
	var request pbauth.LogInRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Log in the user
	response, err := c.client.LogIn(
		grpcCtx, &request,
	)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// logOut logs out a user
// @Summary Log out a user
// @Description Log out a user by invalidating their access token
// @Tags v1 auth
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.LogOutResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/log-out [post]
func (c *Controller) logOut(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Log out the user
	response, err := c.client.LogOut(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
