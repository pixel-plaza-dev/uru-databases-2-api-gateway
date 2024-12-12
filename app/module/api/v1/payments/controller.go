package payments

import (
	"github.com/gin-gonic/gin"
	modulepaymentsaccounts "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/payments/accounts"
	modulepaymentsbranchrents "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/payments/branch-rents"
	modulepaymentsorders "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/payments/orders"
	apptypes "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/types"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commonhandler "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/route"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/payment"
	pbconfiggrpcpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/grpc/payment"
	pbconfigrestpayments "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1/payments"
)

// Controller struct for the payments module
// @Summary Payments Router Group
// @Description Router group for payments-related endpoints
// @Tags v1 payments
// @Accept json
// @Produce json
// @Router /api/v1/payments [group]
type Controller struct {
	route           *gin.RouterGroup
	client          pbpayment.PaymentClient
	authentication  authmiddleware.Authentication
	routeHandler    commonhandler.Handler
	responseHandler commonclientresponse.Handler
}

// NewController creates a new payments controller
func NewController(
	baseRoute *gin.RouterGroup,
	client pbpayment.PaymentClient,
	authentication authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the payments controller
	route := baseRoute.Group(pbconfigrestpayments.Base.String())

	// Create the route handler
	routeHandler := commonhandler.NewDefaultHandler(authentication, &pbconfiggrpcpayment.Interceptions)

	// Create a new payments controller
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
	accountsController := modulepaymentsaccounts.NewController(c.route, c.client, c.routeHandler, c.responseHandler)
	branchRentsController := modulepaymentsbranchrents.NewController(
		c.route,
		c.client,
		c.routeHandler,
		c.responseHandler,
	)
	ordersController := modulepaymentsorders.NewController(c.route, c.client, c.routeHandler, c.responseHandler)

	// Initialize the routes for the children controllers
	for _, controller := range []apptypes.Controller{
		accountsController,
		branchRentsController,
		ordersController,
	} {
		controller.Initialize()
	}
}
