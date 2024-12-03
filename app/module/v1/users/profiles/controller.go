package profiles

import (
	"github.com/gin-gonic/gin"
	appgrpcuser "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/user"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigrestusers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/v1/users"
	pbconfigrestprofiles "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/v1/users/profiles"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
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
	route   *gin.RouterGroup
	service *appgrpcuser.Service
}

// NewController creates a new profiles controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcuser.Service,
) *Controller {
	// Create a new route for the profiles controller
	route := baseRoute.Group(pbconfigrestusers.Profiles.String())

	// Create a new user controller
	return &Controller{
		route:   route,
		service: service,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.GET(pbconfigrestprofiles.Relative.String(), c.getMyProfile)
	c.route.GET(pbconfigrestprofiles.ByUsername.String(), c.getProfile)
}

// getMyProfile gets the user's profile
// @Summary Get the user's profile
// @Description Get the profile information of the authenticated user
// @Tags v1 users profiles
// @Accept json
// @Produce json
// @Success 200 {object} pbuser.GetMyProfileResponse
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/users/profiles [get]
func (c *Controller) getMyProfile(ctx *gin.Context) {
	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, nil)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Get the user's profile
	response, err := c.service.GetMyProfile(ctx, grpcCtx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getProfile gets the user's profile
// @Summary Get a user's profile
// @Description Get the profile information of a user by their username
// @Tags v1 users profiles
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} pbuser.GetProfileResponse
// @Failure 400 {object} commongintypes.BadRequest
// @Failure 500 {object} commongintypes.InternalServerError
// @Router /api/v1/users/profiles/{username} [get]
func (c *Controller) getProfile(ctx *gin.Context) {
	var request pbuser.GetProfileRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongintypes.NewInternalServerError(),
		)
		return
	}

	// Add the username to the request
	request.Username = ctx.Param(pbtypesrest.Username.String())

	// Get the user's profile
	response, err := c.service.GetProfile(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, commongintypes.NewBadRequest(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
