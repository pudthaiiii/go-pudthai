package route

import (
	"github.com/pudthaiiii/golang-cms/src/app/controller"
	"github.com/pudthaiiii/golang-cms/src/app/middleware"
	technical "github.com/pudthaiiii/golang-cms/src/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/pudthaiiii/golang-cms/src/docs"
)

func InitializeRoute(app *fiber.App, c controller.AppController, m middleware.Middleware) *fiber.App {
	addDefaultRoute(app)

	addRoutesV1(app, c, m)

	return app
}

func addDefaultRoute(app *fiber.App) {
	// add default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Fiber!")
	})

	// add swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)
}

func addRoutesV1(app *fiber.App, c controller.AppController, m middleware.Middleware) {
	routes := technical.Routes{}

	// Append routes
	routes = append(routes, addPrototypeRoute(c.PrototypeController)...)
	routes = append(routes, addMerchantRoute(c.MerchantsController)...)

	// Prefix "/v1" for all routes
	apiV1 := app.Group("/v1")

	setupRoutes(apiV1, routes, m)
}

func setupRoutes(router fiber.Router, routes []technical.Route, m middleware.Middleware) {
	for _, route := range routes {
		routeHandler := func(c *fiber.Ctx) error {
			return route.HandlerFunc(c)
		}

		if route.Name == "Show" {
			routeHandler = m.Authenticate(routeHandler)
		}

		routeHandler = m.RequiredMerchant(routeHandler)

		registerRoute(router, route, routeHandler)
	}
}

func registerRoute(router fiber.Router, route technical.Route, handler fiber.Handler) {
	router.Add(route.Method, route.Pattern, handler)
}
