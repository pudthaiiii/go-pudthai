package route

import (
	mc "github.com/pudthaiiii/go-ibooking/src/app/http"
	am "github.com/pudthaiiii/go-ibooking/src/app/middleware/admin"

	"github.com/gofiber/fiber/v2"
)

func InitializeAdminRoute(app *fiber.App, c mc.AppController, m am.Middleware) *fiber.App {
	routes := addPrototypeRoute(c.AdminPrototype)

	// routes := append(
	// 	addPrototypeRoute(c.BackendPrototype),
	// 	// add route here
	// )

	prefix := app.Group("/v1/admin")

	for _, route := range routes {
		handler := route.HandlerFunc

		// add middleware for admin here
		handler = m.Authenticate(handler)

		prefix.Add(route.Method, route.Pattern, handler)
	}

	return app
}
