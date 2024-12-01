package usernames

import (
	"github.com/gin-gonic/gin"
	appgrpcuser "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/grpc/user"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commongrpcclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/context"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/user"
	pbconfigrestusers "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/users"
	pbconfigrestusernames "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/config/rest/v1/users/usernames"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/types/rest"
	"net/http"
)

// Controller struct for the users usernames module
// @Summary Users Usernames Router Group
// @Description Router group for users usernames-related endpoints
// @Tags v1 users usernames
// @Accept json
// @Produce json
// @Router /usernames [group]
type Controller struct {
	route   *gin.RouterGroup
	service *appgrpcuser.Service
}

// NewController creates a new username controller
func NewController(
	baseRoute *gin.RouterGroup,
	service *appgrpcuser.Service,
) *Controller {
	// Create a new route for the usernames controller
	route := baseRoute.Group(pbconfigrestusers.Usernames.String())

	// Create a new user controller
	return &Controller{
		route:   route,
		service: service,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes
	c.route.GET(pbconfigrestusernames.ExistsByUsername.String(), c.usernameExists)
	c.route.GET(pbconfigrestusernames.ByUserId.String(), c.getUsernameByUserId)
}

// usernameExists checks if a username exists
// @Summary Check if a username exists
// @Description Check if a username exists
// @Tags v1 users usernames
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} pbuser.UsernameExistsResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} commongin.InternalServerError
// @Router /exists/{username} [get]
func (c *Controller) usernameExists(ctx *gin.Context) {
	var request pbuser.UsernameExistsRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
		)
		return
	}

	// Add the username to the request
	request.Username = ctx.Param(pbtypesrest.Username.String())

	// Check if the username exists
	response, err := c.service.UsernameExists(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// getUsernameByUserId gets the username by user ID
// @Summary Get username by user ID
// @Description Get username by user ID
// @Tags v1 users usernames
// @Accept json
// @Produce json
// @Param user-id path string true "User ID"
// @Success 200 {object} pbuser.GetUsernameByUserIdResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} commongin.InternalServerError
// @Router /{user-id} [get]
func (c *Controller) getUsernameByUserId(ctx *gin.Context) {
	var request pbuser.GetUsernameByUserIdRequest

	// Prepare the gRPC context
	grpcCtx, err := commongrpcclientctx.PrepareCtx(ctx, &request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			commongin.InternalServerError,
		)
		return
	}

	// Add the user ID to the request
	request.UserId = ctx.Param(pbtypesrest.UserId.String())

	// Get the username by user ID
	response, err := c.service.GetUsernameByUserId(ctx, grpcCtx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
