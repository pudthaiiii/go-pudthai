package router

import (
	routerAdmin "go-ibooking/internal/infrastructure/router/admin"
	routerConsole "go-ibooking/internal/infrastructure/router/console"
	"go-ibooking/internal/registry"

	"github.com/gofiber/fiber/v2"
)

func InitializeRoute(fiber *fiber.App, r registry.Registry) *fiber.App {
	routerConsole.InitializeConsoleRoute(fiber, r.NewConsoleController(), r.NewAdminMiddleware())
	routerAdmin.InitializeAdminRoute(fiber, r.NewAdminController(), r.NewAdminMiddleware())

	return fiber
}
