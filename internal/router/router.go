package router

import (
	_ "go-ibooking/docs"
	ra "go-ibooking/internal/router/admin"
	rb "go-ibooking/internal/router/backend"
	rc "go-ibooking/internal/router/console"
	rf "go-ibooking/internal/router/frontend"

	"go-ibooking/internal/registry"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func InitializeRoute(fiber *fiber.App, r registry.Registry) *fiber.App {
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
	fiber.Use(cors.New(cors.Config{
		AllowOrigins: "https://example.com",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE",
	}))
}
