package router

import (
	controller "go-ibooking/internal/adapter/controllers/v1/console"
	am "go-ibooking/src/app/http/middleware/admin"
	"go-ibooking/src/types"

	"github.com/gofiber/fiber/v2"
)

func InitializeConsoleRoute(app *fiber.App, c controller.FeaturesController, m am.Middleware) *fiber.App {
	var routes = types.Routes{}

	routes = append(routes, addFeaturesRoute(c)...)

	prefix := app.Group("/v1/console")

	for _, route := range routes {
		handler := route.HandlerFunc
		prefix.Name(route.Name)

		// add middleware for admin here
		// handler = m.Authenticate(handler, route.Action, route.Subject)

		prefix.Add(route.Method, route.Path, handler)
	}

	return app
}
