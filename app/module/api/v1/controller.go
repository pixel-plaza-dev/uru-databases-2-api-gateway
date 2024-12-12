package v1

import (
	"github.com/gin-gonic/gin"
	moduleauth "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/auth"
	moduleorders "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/orders"
	modulepayments "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/payments"
	moduleshops "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/shops"
	moduleusers "github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/api/v1/users"
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	_ "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	pborder "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/order"
	pbpayment "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/payment"
	pbshop "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/shop"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	pbconfigrestv1 "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest/api/v1"
)

// Controller struct for the API version 1 module
// @Summary API Version 1 Router Group
// @Description Router group for API version 1-related endpoints
// @Tags v1
// @Accept json
// @Produce json
// @Router /api/v1 [group]
type Controller struct {
	route              *gin.RouterGroup
	authentication     authmiddleware.Authentication
	responseHandler    commonclientresponse.Handler
	usersController    *moduleusers.Controller
	authController     *moduleauth.Controller
	paymentsController *modulepayments.Controller
	ordersController   *moduleorders.Controller
	shopsController    *moduleshops.Controller
}

// NewController creates a new controller
func NewController(
	baseRoute *gin.RouterGroup,
	authentication authmiddleware.Authentication,
	responseHandler commonclientresponse.Handler,
) *Controller {
	// Create a new route for the API version 1 controller
	route := baseRoute.Group(pbconfigrestv1.Base.String())

	// Create a new  controller
	return &Controller{
		route:           route,
		authentication:  authentication,
		responseHandler: responseHandler,
	}
}

// Initialize initializes the routes for the controller
func (c *Controller) Initialize() {}

// InitializeAuth initializes the routes for the API version 1 auth controller
func (c *Controller) InitializeAuth(authClient pbauth.AuthClient) *moduleauth.Controller {
	// Check if the API version 1 auth controller has already been initialized
	if c.authController != nil {
		return c.authController
	}

	// Initialize the API version 1 auth controller
	authController := moduleauth.NewController(c.route, authClient, c.authentication, c.responseHandler)
	authController.Initialize()

	// Store the API version 1 auth controller
	c.authController = authController

	return authController
}

// InitializeUsers initializes the routes for the API version 1 users controller
func (c *Controller) InitializeUsers(userClient pbuser.UserClient) *moduleusers.Controller {
	// Check if the API version 1 users controller has already been initialized
	if c.usersController != nil {
		return c.usersController
	}

	// Initialize the API version 1 users controller
	usersController := moduleusers.NewController(c.route, userClient, c.authentication, c.responseHandler)
	usersController.Initialize()

	// Store the API version 1 users controller
	c.usersController = usersController

	return usersController
}

// InitializeShops initializes the routes for the API version 1 shops controller
func (c *Controller) InitializeShops(shopClient pbshop.ShopClient) *moduleshops.Controller {
	// Check if the API version 1 shops controller has already been initialized
	if c.shopsController != nil {
		return c.shopsController
	}

	// Initialize the API version 1 shops controller
	shopsController := moduleshops.NewController(c.route, shopClient, c.authentication, c.responseHandler)
	shopsController.Initialize()

	// Store the API version 1 shops controller
	c.shopsController = shopsController

	return shopsController
}

// InitializeOrders initializes the routes for the API version 1 orders controller
func (c *Controller) InitializeOrders(orderClient pborder.OrderClient) *moduleorders.Controller {
	// Check if the API version 1 orders controller has already been initialized
	if c.ordersController != nil {
		return c.ordersController
	}

	// Initialize the API version 1 orders controller
	ordersController := moduleorders.NewController(c.route, orderClient, c.authentication, c.responseHandler)
	ordersController.Initialize()

	// Store the API version 1 orders controller
	c.ordersController = ordersController

	return ordersController
}

// InitializePayments initializes the routes for the API version 1 payments controller
func (c *Controller) InitializePayments(paymentClient pbpayment.PaymentClient) *modulepayments.Controller {
	// Check if the API version 1 payments controller has already been initialized
	if c.paymentsController != nil {
		return c.paymentsController
	}

	// Initialize the API version 1 payments controller
	paymentsController := modulepayments.NewController(c.route, paymentClient, c.authentication, c.responseHandler)
	paymentsController.Initialize()

	// Store the API version 1 payments controller
	c.paymentsController = paymentsController

	return paymentsController
}
