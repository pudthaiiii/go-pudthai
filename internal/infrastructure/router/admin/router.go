package router

import (
	_ "go-ibooking/docs"
	ca "go-ibooking/internal/adapter/v1/controllers/admin"
	t "go-ibooking/internal/model/technical"

	"github.com/gofiber/fiber/v2"
)

func InitializeAdminRoute(app *fiber.App, c ca.AdminController) *fiber.App {
	var routes = t.Routes{}
	prefix := app.Group("/v1/admin")

	routes = append(routes, addAdminAuthRoute(c.AuthController)...)
	routes = append(routes, addAdminUsersRoute(c.UsersController)...)

	for _, route := range routes {
		handler := route.HandlerFunc
		// handler = m.Authenticate(handler, route.Action, route.Subject)

		prefix.Name(route.Name)
		prefix.Add(route.Method, route.Path, handler)
	}

	return app
}
