package businesses

import (
	"github.com/gin-gonic/gin"
	moduleshopscategories "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/markets/categories"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfigrestmarkets "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops/markets"
)

// Controller struct for the markets module
// @Summary Shops Markets Router Group
// @Description Router group for shops markets-related endpoints
// @Tags v1 shops markets
// @Accept json
// @Produce json
// @Router /api/v1/shops/markets [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new markets controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	routeHandler commonhandler.Handler,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the markets controller
	route := baseRoute.Group(pbconfigrestmarkets.Base.String())

	// Create a new markets controller
	return &Controller{
		route:           route,
		client:          client,
		routeHandler:    routeHandler,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {
	// Initialize the routes for the children controllers
	c.initializeChildren()
}

// initializeChildren initializes the routes for the children controllers
func (c *Controller) initializeChildren() {
	// Create the children controllers
	categoriesController := moduleshopscategories.NewController(
		c.route,
		c.client,
		c.routeHandler,
		c.responseHandler,
	)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		categoriesController,
	} {
		controller.Initialize()
	}
}
