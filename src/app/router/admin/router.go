package route

import (
	controller "go-ibooking/src/app/http/admin"
	am "go-ibooking/src/app/middleware/admin"

	"github.com/gofiber/fiber/v2"
)

func InitializeAdminRoute(app *fiber.App, c controller.AdminController, m am.Middleware) *fiber.App {
	routes := addPrototypeRoute(c.PrototypeController)

	routes = append(
		addRoleRoute(c.RoleController),
		// add route here
	)

	prefix := app.Group("/v1/admin")

	for _, route := range routes {
		handler := route.HandlerFunc

		// add middleware for admin here
		handler = m.Authenticate(handler)

		prefix.Add(route.Method, route.Pattern, handler)
	}

	return app
}
