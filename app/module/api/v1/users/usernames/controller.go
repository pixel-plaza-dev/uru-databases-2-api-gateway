package usernames

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigrestusernames "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/users/usernames"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the users usernames module
// @Summary Users Usernames Router Group
// @Description Router group for users usernames-related endpoints
// @Tags v1 users usernames
// @Accept json
// @Produce json
// @Router /api/v1/users/usernames [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbuser.UserClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new username controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbuser.UserClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the usernames controller
	route := baseRoute.Group(pbconfigrestusernames.Base.String())

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
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestusernames.UsernameExistsMapper,
			c.usernameExists,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestusernames.GetUsernameByUserIdMapper,
			c.getUsernameByUserId,
		),
	)
}

// usernameExists checks if a username exists
// @Summary Check if a username exists
// @Description Check if a username exists
// @Tags v1 users usernames
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} pbuser.UsernameExistsResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/usernames/exists/{username} [get]
func (c *Controller) usernameExists(ctx *gin.Context) {
	var request pbuser.UsernameExistsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the username to the request
	request.Username = ctx.Param(pbtypesrest.Username.String())

	// Check if the username exists
	response, err := c.client.UsernameExists(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getUsernameByUserId gets the username by user ID
// @Summary Get username by user ID
// @Description Get username by user ID
// @Tags v1 users usernames
// @Accept json
// @Produce json
// @Param user-id path string true "User ID"
// @Success 200 {object} pbuser.GetUsernameByUserIdResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/users/usernames/{user-id} [get]
func (c *Controller) getUsernameByUserId(ctx *gin.Context) {
	var request pbuser.GetUsernameByUserIdRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the user ID to the request
	request.UserId = ctx.Param(pbtypesrest.UserId.String())

	// Get the username by user ID
	response, err := c.client.GetUsernameByUserId(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
