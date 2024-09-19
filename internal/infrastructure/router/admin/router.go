package router

import (
	controller "go-ibooking/src/app/http/admin"
	am "go-ibooking/src/app/http/middleware/admin"
	"go-ibooking/src/types"

	"github.com/gofiber/fiber/v2"
)

func InitializeAdminRoute(app *fiber.App, c controller.AdminController, m am.Middleware) *fiber.App {
	var routes = types.Routes{}

	routes = append(routes, addRolesRoute(c.RoleController)...)
	routes = append(routes, addUsersRoute(c.UsersController)...)
	routes = append(routes, addAuthRoute(c.AuthController)...)

	prefix := app.Group("/v1/admin")

	for _, route := range routes {
		handler := route.HandlerFunc
		prefix.Name(route.Name)

		// add middleware for admin here
		handler = m.Authenticate(handler, route.Action, route.Subject)

		prefix.Add(route.Method, route.Path, handler)

	}

	return app
}
