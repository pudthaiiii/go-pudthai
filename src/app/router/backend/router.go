package route

import (
	controller "github.com/pudthaiiii/go-ibooking/src/app/http/backend"
	am "github.com/pudthaiiii/go-ibooking/src/app/middleware/admin"

	"github.com/gofiber/fiber/v2"
)

func InitializeBackendRoute(app *fiber.App, c controller.BackendController, m am.Middleware) *fiber.App {
	routes := addPrototypeRoute(c.PrototypeController)

	// routes := append(
	// 	addPrototypeRoute(c.BackendPrototype),
	// 	// add route here
	// )

	prefix := app.Group("/v1/backend")

	for _, route := range routes {
		handler := route.HandlerFunc

		// add middleware for admin here
		handler = m.Authenticate(handler)

		prefix.Add(route.Method, route.Pattern, handler)
	}

	return app
}
