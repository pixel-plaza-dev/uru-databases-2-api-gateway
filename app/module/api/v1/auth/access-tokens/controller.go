package access_tokens

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfigrestaccesstokens "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/access-tokens"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"net/http"
)

// Controller struct for the access tokens module
// @Summary Auth Access Tokens Router Group
// @Description Router group for auth access tokens-related endpoints
// @Tags v1 auth access-tokens
// @Accept json
// @Produce json
// @Router /api/v1/auth/access-tokens [group]
type Controller struct {
	route           *gin.RouterGroup
	service         *appgrpcauth.Service
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new access tokens controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcauth.Service,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the access tokens controller
	route := baseRoute.Group(pbconfigrestaccesstokens.Base.String())

	// Create a new access tokens controller
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
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestaccesstokens.IsAccessTokenValidMapper,
			c.isAccessTokenValid,
		),
	)
}

// isAccessTokenValid checks if an access token is valid
// @Summary Check if an access token is valid
// @Description Check if an access token is valid by its JWT ID
// @Tags v1 auth access-tokens
// @Accept json
// @Produce json
// @Param jwt-id path string true "JWT ID"
// @Success 200 {object} pbauth.IsAccessTokenValidResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/auth/access-tokens/valid/{jwt-id} [get]
func (c *Controller) isAccessTokenValid(ctx *gin.Context) {
	var request pbauth.IsAccessTokenValidRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the JWT Identifier to the request
	request.JwtId = ctx.Param(pbtypesrest.JwtId.String())

	// Check if the access token is valid
	response, err := c.service.IsAccessTokenValid(ctx, grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
