package router

import (
	smm "go-pudthai/internal/adapter/shared/middleware"
	ac "go-pudthai/internal/adapter/v1/admin/controllers"
	t "go-pudthai/internal/model/technical"

	"github.com/gofiber/fiber/v2"
)

func InitializeAdminRoute(app *fiber.App, c ac.AdminController, sm smm.Middleware) *fiber.App {
	var routes = t.Routes{}
	prefix := app.Group("/v1/admin")

	routes = append(routes, addAuthRoute(c.AuthController)...)
	routes = append(routes, addUsersRoute(c.UsersController)...)

	for _, route := range routes {
		handler := route.HandlerFunc
		prefix.Name(route.Name)

		handler = sm.Authenticate(handler, route.Action, route.Subject)
		handler = sm.GoogleRecaptcha(handler, route.Action, route.Subject)

		prefix.Add(route.Method, route.Path, handler)
	}

	return app
}
