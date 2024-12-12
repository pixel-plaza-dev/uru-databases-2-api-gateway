package refresh_tokens

import (
	"github.com/gin-gonic/gin"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pbconfigrestrefreshtokens "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/auth/refresh-tokens"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

// Controller struct for the refresh tokens module
// @Summary Auth Refresh Tokens Router Group
// @Description Router group for auth refresh tokens-related endpoints
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Router /api/v1/auth/refresh-tokens [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbauth.AuthClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new refresh tokens controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbauth.AuthClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the refresh tokens controller
	route := baseRoute.Group(pbconfigrestrefreshtokens.Base.String())

	// Create a new refresh tokens controller
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
	c.route.POST(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrefreshtokens.RefreshTokenMapper,
			c.refreshToken,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrefreshtokens.GetRefreshTokensInformationMapper,
			c.getRefreshTokensInformation,
		),
	)
	c.route.DELETE(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrefreshtokens.RevokeRefreshTokensMapper,
			c.revokeRefreshTokens,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrefreshtokens.GetRefreshTokenInformationMapper,
			c.getRefreshTokenInformation,
		),
	)
	c.route.DELETE(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrefreshtokens.RevokeRefreshTokenMapper,
			c.revokeRefreshToken,
		),
	)
	c.route.GET(
		c.routeHandler.CreateAuthenticatedEndpoint(
			pbconfigrestrefreshtokens.IsRefreshTokenValidMapper,
			c.isRefreshTokenValid,
		),
	)
}

// isRefreshTokenValid checks if a refresh token is valid
// @Summary Check if a refresh token is valid
// @Description Check if a refresh token is valid by its JWT ID
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Param jwt-id path string true "JWT ID"
// @Success 200 {object} pbauth.IsRefreshTokenValidResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Router /api/v1/auth/refresh-tokens/valid/{jwt-id} [get]
func (c *Controller) isRefreshTokenValid(ctx *gin.Context) {
	var request pbauth.IsRefreshTokenValidRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the JWT Identifier to the request
	request.JwtId = ctx.Param(pbtypesrest.JwtId.String())

	// Check if the refresh token is valid
	response, err := c.client.IsRefreshTokenValid(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getRefreshTokensInformation gets all refresh tokens information
// @Summary Get all refresh tokens information
// @Description Get information about all refresh tokens for a user
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.GetRefreshTokensInformationResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/refresh-tokens [get]
func (c *Controller) getRefreshTokensInformation(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Get all user's refresh tokens information
	response, err := c.client.GetRefreshTokensInformation(
		grpcCtx,
		&emptypb.Empty{},
	)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// refreshToken refreshes a user's token
// @Summary Refresh a user's token
// @Description Refresh a user's token using a refresh token
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.RefreshTokenResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/refresh-tokens [post]
func (c *Controller) refreshToken(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Refresh the token
	response, err := c.client.RefreshToken(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// revokeRefreshTokens revokes all user's refresh tokens
// @Summary Revoke all refresh tokens
// @Description Revoke all refresh tokens for a user
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.RevokeRefreshTokensResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/refresh-tokens [delete]
func (c *Controller) revokeRefreshTokens(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Revoke all user's refresh tokens
	response, err := c.client.RevokeRefreshTokens(grpcCtx, &emptypb.Empty{})
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// getRefreshTokenInformation gets a refresh token information
// @Summary Get refresh token information
// @Description Get information about a specific refresh token by its JWT ID
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Param jwt-id path string true "JWT ID"
// @Success 200 {object} pbauth.GetRefreshTokenInformationResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/refresh-tokens/{jwt-id} [get]
func (c *Controller) getRefreshTokenInformation(ctx *gin.Context) {
	var request pbauth.GetRefreshTokenInformationRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the JWT Identifier to the request
	request.JwtId = ctx.Param(pbtypesrest.JwtId.String())

	// Get the refresh token information
	response, err := c.client.GetRefreshTokenInformation(
		grpcCtx,
		&request,
	)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}

// revokeRefreshToken revokes a user's refresh token
// @Summary Revoke a refresh token
// @Description Revoke a specific refresh token by its JWT ID
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Param jwt-id path string true "JWT ID"
// @Success 200 {object} pbauth.RevokeRefreshTokenResponse
// @Failure 400 {object} _.ErrorResponse
// @Failure 500 {object} _.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/auth/refresh-tokens/{jwt-id} [delete]
func (c *Controller) revokeRefreshToken(ctx *gin.Context) {
	var request pbauth.RevokeRefreshTokenRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		c.responseHandler.HandlePrepareCtxError(ctx, err)
		return
	}

	// Add the JWT ID to the request
	request.JwtId = ctx.Param(pbtypesrest.JwtId.String())

	// Revoke the given refresh token
	response, err := c.client.RevokeRefreshToken(grpcCtx, &request)
	c.responseHandler.HandleResponse(ctx, http.StatusOK, response, err)
}
