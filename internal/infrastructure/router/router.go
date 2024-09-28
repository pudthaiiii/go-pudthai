package router

import (
	_ "go-ibooking/docs"
	ra "go-ibooking/internal/infrastructure/router/admin"
	rc "go-ibooking/internal/infrastructure/router/console"
	"go-ibooking/internal/registry"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func InitializeRoute(fiber *fiber.App, r registry.Registry) *fiber.App {
	fiber.Get("/swagger/*", swagger.HandlerDefault)

	ra.InitializeAdminRoute(fiber, r.NewAdminController())

	rc.InitializeConsoleRoute(fiber, r.NewConsoleController(), r.NewConsoleMiddleware())

	return fiber
}
