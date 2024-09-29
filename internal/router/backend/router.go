package router

import (
	smm "go-pudthai/internal/adapter/shared/middleware"
	bc "go-pudthai/internal/adapter/v1/backend/controllers"
	t "go-pudthai/internal/model/technical"

	"github.com/gofiber/fiber/v2"
)

func InitializeBackendRoute(app *fiber.App, c bc.BackendController, sm smm.Middleware) *fiber.App {
	var routes = t.Routes{}
	prefix := app.Group("/v1/backend")

	routes = append(routes, addAuthRoute(c.AuthController)...)

	for _, route := range routes {
		handler := route.HandlerFunc
		prefix.Name(route.Name)

		handler = sm.RequiredMerchant(handler, route.Action, route.Subject)
		handler = sm.Authenticate(handler, route.Action, route.Subject)
		handler = sm.GoogleRecaptcha(handler, route.Action, route.Subject)
		handler = sm.Log(handler, route.Name, route.Action, route.Subject)

		prefix.Add(route.Method, route.Path, handler)
	}

	return app
}
