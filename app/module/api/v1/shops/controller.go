package products

import (
	"github.com/gin-gonic/gin"
	moduleshopsbusinesses "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/businesses"
	moduleshopsmarkets "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/markets"
	moduleshopsproducts "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/products"
	moduleshopsstores "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops/stores"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbconfiggrpcshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/shop"
	pbconfigrestshops "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/shops"
)

// Controller struct for the shops module
// @Summary Shops Router Group
// @Description Router group for shops-related endpoints
// @Tags v1 shops
// @Accept json
// @Produce json
// @Router /api/v1/shops [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbshop.ShopClient
	authentication  authmiddleware.Authentication
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new shops controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbshop.ShopClient,
	authentication authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the shops controller
	route := baseRoute.Group(pbconfigrestshops.Base.String())

	// Create the route handler
	routeHandler := commonhandler.NewDefaultHandler(authentication, &pbconfiggrpcshop.Interceptions)

	// Create a new shops controller
	return &Controller{
		route:           route,
		client:          client,
		authentication:  authentication,
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
	marketsController := moduleshopsmarkets.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	businessesController := moduleshopsbusinesses.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	productsController := moduleshopsproducts.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	storesController := moduleshopsstores.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		marketsController,
		businessesController,
		productsController,
		storesController,
	} {
		controller.Initialize()
	}
}
