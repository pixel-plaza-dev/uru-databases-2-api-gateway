package refresh_tokens

import (
	"github.com/gin-gonic/gin"
	appgrpcauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/auth"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
	pbconfigrestauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/auth"
	pbconfigrestrefreshtokens "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/auth/refresh-tokens"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
	"net/http"
)

// Controller struct for the refresh tokens module
// @Summary Auth Refresh Tokens Router Group
// @Description Router group for auth refresh tokens-related endpoints
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Router /refresh-tokens [group]
type Controller struct {
	route   *gin.RouterGroup
	service *appgrpcauth.Service
}

// NewController creates a new refresh tokens controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcauth.Service,
) *Controller {
	// Create a new route for the refresh tokens controller
	route := baseRoute.Group(pbconfigrestauth.RefreshTokens.String())

	// Create a new refresh tokens controller
	return &Controller{
		route:   route,
		service: service,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.POST(pbconfigrestrefreshtokens.Relative.String(), c.refreshToken)
	c.route.GET(pbconfigrestrefreshtokens.Relative.String(), c.getRefreshTokensInformation)
	c.route.DELETE(pbconfigrestrefreshtokens.Relative.String(), c.revokeRefreshTokens)
	c.route.GET(
		pbconfigrestrefreshtokens.ByJwtId.String(),
		c.getRefreshTokenInformation,
	)
	c.route.DELETE(pbconfigrestrefreshtokens.ByJwtId.String(), c.revokeRefreshToken)
	c.route.GET(
		pbconfigrestrefreshtokens.Valid.String(),
		c.isRefreshTokenValid,
	)
}

// isRefreshTokenValid checks if a refresh token is valid
// @Summary Check if a refresh token is valid
// @Description Check if a refresh token is valid by its JWT ID
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Param jwtId path string true "JWT ID"
// @Success 200 {object} pbauth.IsRefreshTokenValidResponse
// @Failure 400 {object} module.BadRequest
// @Failure 500 {object} commongin.InternalServerError
// @Router /valid/{jwtId} [get]
func (c *Controller) isRefreshTokenValid(ctx *gin.Context) {
	var request pbauth.IsRefreshTokenValidRequest

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

	// Check if the refresh token is valid
	response, err := c.service.IsRefreshTokenValid(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getRefreshTokensInformation gets all refresh tokens information
// @Summary Get all refresh tokens information
// @Description Get information about all refresh tokens for a user
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.GetRefreshTokensInformationResponse
// @Failure 400 {object} module.BadRequest
// @Failure 500 {object} commongin.InternalServerError
// @Router / [get]
func (c *Controller) getRefreshTokensInformation(ctx *gin.Context) {
	var request pbauth.GetRefreshTokensInformationRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
		)
		return
	}

	// Get all user' sessions
	response, err := c.service.GetRefreshTokensInformation(
		ctx,
		grpcCtx,
		&request,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// refreshToken refreshes a user's token
// @Summary Refresh a user's token
// @Description Refresh a user's token using a refresh token
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.RefreshTokenResponse
// @Failure 400 {object} module.BadRequest
// @Failure 500 {object} commongin.InternalServerError
// @Router / [post]
func (c *Controller) refreshToken(ctx *gin.Context) {
	var request pbauth.RefreshTokenRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
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

// revokeRefreshTokens revokes all user's refresh tokens
// @Summary Revoke all refresh tokens
// @Description Revoke all refresh tokens for a user
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Success 200 {object} pbauth.RevokeRefreshTokensResponse
// @Failure 400 {object} module.BadRequest
// @Failure 500 {object} commongin.InternalServerError
// @Router / [delete]
func (c *Controller) revokeRefreshTokens(ctx *gin.Context) {
	var request pbauth.RevokeRefreshTokensRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
		)
		return
	}

	// Revoke all user's refresh tokens
	response, err := c.service.RevokeRefreshTokens(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getRefreshTokenInformation gets a refresh token information
// @Summary Get refresh token information
// @Description Get information about a specific refresh token by its JWT ID
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Param jwtId path string true "JWT ID"
// @Success 200 {object} pbauth.GetRefreshTokenInformationResponse
// @Failure 400 {object} module.BadRequest
// @Failure 500 {object} commongin.InternalServerError
// @Router /{jwtId} [get]
func (c *Controller) getRefreshTokenInformation(ctx *gin.Context) {
	var request pbauth.GetRefreshTokenInformationRequest

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

	// Get the refresh token information
	response, err := c.service.GetRefreshTokenInformation(
		ctx,
		grpcCtx,
		&request,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// revokeRefreshToken revokes a user's refresh token
// @Summary Revoke a refresh token
// @Description Revoke a specific refresh token by its JWT ID
// @Tags v1 auth refresh-tokens
// @Accept json
// @Produce json
// @Param jwtId path string true "JWT ID"
// @Success 200 {object} pbauth.RevokeRefreshTokenResponse
// @Failure 400 {object} module.BadRequest
// @Failure 500 {object} commongin.InternalServerError
// @Router /{jwtId} [delete]
func (c *Controller) revokeRefreshToken(ctx *gin.Context) {
	var request pbauth.RevokeRefreshTokenRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
		)
		return
	}

	// Add the JWT ID to the request
	request.JwtId = ctx.Param(pbtypesrest.JwtId.String())

	// Revoke the given refresh token
	response, err := c.service.RevokeRefreshToken(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
