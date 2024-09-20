package router

import (
	controller "go-ibooking/internal/adapter/v1/controllers/console"
	"go-ibooking/internal/model/technical"

	"github.com/gofiber/fiber/v2"
)

func InitializeConsoleRoute(app *fiber.App, c controller.FeaturesController) *fiber.App {
	var routes = technical.Routes{}
	prefix := app.Group("/v1/console")

	routes = append(routes, addFeaturesRoute(c)...)

	for _, route := range routes {
		handler := route.HandlerFunc
		prefix.Name(route.Name)

		prefix.Add(route.Method, route.Path, handler)
	}

	return app
}
