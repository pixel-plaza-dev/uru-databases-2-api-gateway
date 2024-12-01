package access_tokens

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
	pbconfigrestauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/auth"
	pbconfigrestaccesstokens "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/auth/access-tokens"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
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
	route   *gin.RouterGroup
	service *appgrpcauth.Service
}

// NewController creates a new access tokens controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcauth.Service,
) *Controller {
	// Create a new route for the access tokens controller
	route := baseRoute.Group(pbconfigrestauth.AccessTokens.String())

	// Create a new access tokens controller
	return &Controller{
		route:   route,
		service: service,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.GET(
		pbconfigrestaccesstokens.Valid.String(),
		c.isAccessTokenValid,
	)
}

// isAccessTokenValid checks if an access token is valid
// @Summary Check if an access token is valid
// @Description Check if an access token is valid by its JWT ID
// @Tags v1 auth access-tokens
// @Accept json
// @Produce json
// @Param jwtId path string true "JWT ID"
// @Success 200 {object} pbauth.IsAccessTokenValidResponse
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /api/v1/auth/access-tokens/valid/{jwtId} [get]
func (c *Controller) isAccessTokenValid(ctx *gin.Context) {
	var request pbauth.IsAccessTokenValidRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
		)
		return
	}

	// Add the JWT Identifier to the request
	request.JwtId = ctx.Param(pbtypesrest.JwtId.String())

	// Check if the access token is valid
	response, err := c.service.IsAccessTokenValid(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
