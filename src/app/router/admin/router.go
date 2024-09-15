package route

import (
	controller "go-ibooking/src/app/http/admin"
	am "go-ibooking/src/app/middleware/admin"
	"go-ibooking/src/types"

	"github.com/gofiber/fiber/v2"
)

func InitializeAdminRoute(app *fiber.App, c controller.AdminController, m am.Middleware) *fiber.App {
	var routes = types.Routes{}

	routes = append(routes, addPrototypeRoute(c.PrototypeController)...)
	routes = append(routes, addRoleRoute(c.RoleController)...)
	routes = append(routes, addUsersRoute(c.UsersController)...)

	prefix := app.Group("/v1/admin")

	for _, route := range routes {
		handler := route.HandlerFunc

		// add middleware for admin here
		handler = m.Authenticate(handler)

		prefix.Add(route.Method, route.Pattern, handler)
		prefix.Name(route.Name)
	}

	return app
}
