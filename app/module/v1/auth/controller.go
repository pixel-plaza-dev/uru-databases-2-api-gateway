package auth

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	moduleauthaccesstokens "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/v1/auth/access-tokens"
	moduleauthpermissions "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/v1/auth/permissions"
	moduleauthrefreshtokens "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/v1/auth/refresh-tokens"
	moduleauthrolepermissions "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/v1/auth/role-permissions"
	moduleauthroles "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/v1/auth/roles"
	moduleauthuserroles "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/v1/auth/user-roles"
	apptypescontroller "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types/controller"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientrequest "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/request"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfiggrpcauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/auth"
	pbconfigrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/v1"
	pbconfigrestauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/v1/auth"
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
	route          *gin.RouterGroup
	client         pbauth.AuthClient
	service        *appgrpcauth.Service
	authMiddleware authmiddleware.Authentication
	requestHandler commonclientrequest.Handler
}

// NewController creates a new auth controller
func NewController(
	apiRoute *gin.RouterGroup,
	client pbauth.AuthClient,
	authMiddleware authmiddleware.Authentication,
	requestHandler commonclientrequest.Handler,
) *Controller {
	// Create a new route for the auth controller
	route := apiRoute.Group(pbconfigrest.Auth.String())

	// Create the auth service
	service := appgrpcauth.NewService(client, requestHandler)

	// Add the auth middleware to the route
	route.Use(
		authMiddleware.Authenticate(
			route.BasePath(),
			pbconfigrestauth.Map,
			&pbconfiggrpcauth.Interceptions,
		),
	)

	// Create a new auth controller
	return &Controller{
		route:          route,
		client:         client,
		service:        service,
		authMiddleware: authMiddleware,
		requestHandler: requestHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.POST(pbconfigrestauth.LogIn.String(), c.logIn)
	c.route.POST(pbconfigrestauth.LogOut.String(), c.logOut)

	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	accessTokensController := moduleauthaccesstokens.NewController(c.route, c.service)
	refreshTokensController := moduleauthrefreshtokens.NewController(c.route, c.service)
	permissionsController := moduleauthpermissions.NewController(c.route, c.service)
	rolePermissionsController := moduleauthrolepermissions.NewController(c.route, c.service)
	rolesController := moduleauthroles.NewController(c.route, c.service)
	userRolesController := moduleauthuserroles.NewController(c.route, c.service)

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
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/log-in [post]
func (c *Controller) logIn(ctx *gin.Context) {
	var request pbauth.LogInRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Log in the user
	response, err := c.service.LogIn(
		ctx, grpcCtx, &request,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
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
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/auth/log-out [post]
func (c *Controller) logOut(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Log out the user
	response, err := c.service.LogOut(ctx, grpcCtx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
