package router

import (
	smm "go-ibooking/internal/adapter/shared/middleware"
	fc "go-ibooking/internal/adapter/v1/frontend/controllers"
	t "go-ibooking/internal/model/technical"

	"github.com/gofiber/fiber/v2"
)

func InitializeFrontendRoute(app *fiber.App, c fc.FrontendController, sm smm.Middleware) *fiber.App {
	var routes = t.Routes{}
	prefix := app.Group("/v1/frontend")

	routes = append(routes, addAuthRoute(c.AuthController)...)

	for _, route := range routes {
		handler := route.HandlerFunc

		handler = sm.RequiredMerchant(handler, route.Action, route.Subject)

		handler = sm.Authenticate(handler, route.Action, route.Subject)

		prefix.Name(route.Name)
		prefix.Add(route.Method, route.Path, handler)
	}

	return app
}
