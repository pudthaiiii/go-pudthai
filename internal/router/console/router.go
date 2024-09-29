package router

import (
	cc "go-pudthai/internal/adapter/console/controllers"
	cm "go-pudthai/internal/adapter/console/middleware"
	"go-pudthai/internal/model/technical"

	"github.com/gofiber/fiber/v2"
)

func InitializeConsoleRoute(app *fiber.App, c cc.ConsoleController, m cm.Middleware) *fiber.App {
	prefix := app.Group("/console")

	var routes = technical.Routes{}
	routes = append(routes, addDatabaseRoute(c.DatabaseController)...)

	for _, route := range routes {
		handler := route.HandlerFunc
		prefix.Name(route.Name)

		handler = m.CookieAuthenticate(handler, route.Action, route.Subject)

		prefix.Add(route.Method, route.Path, handler)
	}

	return app
}
