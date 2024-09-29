package router

import (
	_ "go-pudthai/docs"
	"go-pudthai/internal/registry"
	ra "go-pudthai/internal/router/admin"
	rb "go-pudthai/internal/router/backend"
	rc "go-pudthai/internal/router/console"
	rf "go-pudthai/internal/router/frontend"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"go.elastic.co/apm/module/apmfiber"
)

func InitializeRoute(fiber *fiber.App, r registry.Registry) *fiber.App {
	fiber.Static("/static", "./public")
	fiber.Get("/swagger/*", swagger.HandlerDefault)

	appleMiddlewares(fiber)

	// Admin
	ra.InitializeAdminRoute(fiber, r.NewAdminController(), r.NewSharedMiddleware())

	// Backend
	rb.InitializeBackendRoute(fiber, r.NewBackendController(), r.NewSharedMiddleware())

	// Frontend
	rf.InitializeFrontendRoute(fiber, r.NewFrontendController(), r.NewSharedMiddleware())

	// Console
	rc.InitializeConsoleRoute(fiber, r.NewConsoleController(), r.NewConsoleMiddleware())

	return fiber
}

func appleMiddlewares(fiber *fiber.App) {

	fiber.Use(apmfiber.Middleware())

	fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE",
	}))
}
