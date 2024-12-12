package profiles

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigrestprofiles "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users/profiles"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

// Controller struct for the users profiles module
// @Summary Users Profiles Router Group
// @Description Router group for users profiles-related endpoints
// @Tags v1 users profiles
// @Accept json
// @Produce json
// @Router /api/v1/users/profiles [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbuser.UserClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new profiles controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbuser.UserClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the profiles controller
	route := baseRoute.Group(pbconfigrestprofiles.Base.String())

	// Create a new user controller
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
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestprofiles.GetMyProfileMapper, c.getMyProfile))
	c.route.GET(c.routeHandler.CreateAuthenticatedEndpoint(pbconfigrestprofiles.GetProfileMapper, c.getProfile))
}

// getMyProfile gets the user's profile
// @Summary Get the user's profile
// @Description Get the profile information of the authenticated user
// @Tags v1 users profiles
// @Accept json
// @Produce json
// @Success 200 {object} pbuser.GetMyProfileResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/users/profiles [get]
func (c *Controller) getMyProfile(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get the user's profile
	response, err := c.client.GetMyProfile(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getProfile gets the user's profile
// @Summary Get a user's profile
// @Description Get the profile information of a user by their username
// @Tags v1 users profiles
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} pbuser.GetProfileResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/profiles/{username} [get]
func (c *Controller) getProfile(ctx *gin.Context) {
	var request pbuser.GetProfileRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the username to the request
	request.Username = ctx.Param(pbtypesrest.Username.String())

	// Get the user's profile
	response, err := c.client.GetProfile(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
