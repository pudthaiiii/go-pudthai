package router

import (
	_ "go-ibooking/docs"
	"go-ibooking/internal/adapter/v1/controllers"
	routerConsole "go-ibooking/internal/infrastructure/router/console"
	"go-ibooking/internal/model/technical"
	"go-ibooking/internal/registry"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func InitializeRoute(fiber *fiber.App, r registry.Registry) *fiber.App {
	fiber.Get("/swagger/*", swagger.HandlerDefault)

	routerConsole.InitializeConsoleRoute(fiber, r.NewConsoleController())

	initializeRoute(fiber, r.NewController())

	return fiber
}

func initializeRoute(app *fiber.App, c controllers.AppController) *fiber.App {
	var routes = technical.Routes{}
	prefix := app.Group("/v1")

	routes = append(routes, addUsersRoute(c.UsersController)...)
	routes = append(routes, addAuthRoute(c.AuthController)...)

	for _, route := range routes {
		handler := route.HandlerFunc
		// handler = m.Authenticate(handler, route.Action, route.Subject)

		prefix.Name(route.Name)
		prefix.Add(route.Method, route.Path, handler)
	}

	return app
}
