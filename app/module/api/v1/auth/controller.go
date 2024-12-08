package auth

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	moduleauthaccesstokens "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/access-tokens"
	moduleauthpermissions "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/permissions"
	moduleauthrefreshtokens "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/refresh-tokens"
	moduleauthrolepermissions "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/role-permissions"
	moduleauthroles "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/roles"
	moduleauthuserroles "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth/user-roles"
	apptypescontroller "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types/controller"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfiggrpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	pbconfigrestauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth"
	"net/http"
)

// Controller struct for the auth module
// @Summary Auth Router Group
// @Description Router group for authentication-related endpoints
// @Tags auth
// @Accept json
// @Produce json
// @Router /api/v1/auth [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbauth.AuthClient
	service         *appgrpcauth.Service
	authMiddleware  authmiddleware.Authentication
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new auth controller
func NewController(
	apiRoute *gin.RouterGroup,
	client pbauth.AuthClient,
	authMiddleware authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the auth controller
	route := apiRoute.Group(pbconfigrestauth.Base.String())

	// Create the route handler
	routeHandler := commonhandler.NewDefaultHandler(authMiddleware, &pbconfiggrpcauth.Interceptions)

	// Create the auth service
	service := appgrpcauth.NewService(client, responseHandler)

	// Create a new auth controller
	return &Controller{
		route:           route,
		client:          client,
		service:         service,
		authMiddleware:  authMiddleware,
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
		c.service,
		c.routeHandler,
		c.responseHandler,
	)
	refreshTokensController := moduleauthrefreshtokens.NewController(
		c.route,
		c.service,
		c.routeHandler,
		c.responseHandler,
	)
	permissionsController := moduleauthpermissions.NewController(c.route, c.service, c.routeHandler, c.responseHandler)
	rolePermissionsController := moduleauthrolepermissions.NewController(
		c.route,
		c.service,
		c.routeHandler,
		c.responseHandler,
	)
	rolesController := moduleauthroles.NewController(c.route, c.service, c.routeHandler, c.responseHandler)
	userRolesController := moduleauthuserroles.NewController(c.route, c.service, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypescontroller.Controller{
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
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LogInRequest true "Log In Request"
// @Success 200 {object} LogInResponse
// @Failure 400 {object} commongintypes.ErrorResponse
// @Failure 500 {object} commongintypes.ErrorResponse
// @Router /api/v1/auth/log-in [post]
func (c *Controller) logIn(ctx *gin.Context) {
	var request pbauth.LogInRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request, c.responseHandler)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewErrorResponse(err),
		)
		return
	}

	// Log in the user
	response, err := c.service.LogIn(
		ctx, grpcCtx, &request,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// logOut logs out a user
// @Summary Log out a user
// @Description Log out a user by invalidating their access token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.LogOutResponse
// @Failure 400 {object} commongintypes.ErrorResponse
// @Failure 500 {object} commongintypes.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/log-out [post]
func (c *Controller) logOut(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil, c.responseHandler)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewErrorResponse(err),
		)
		return
	}

	// Log out the user
	response, err := c.service.LogOut(ctx, grpcCtx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
