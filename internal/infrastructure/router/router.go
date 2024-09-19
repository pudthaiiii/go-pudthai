package router

import (
	routerAdmin "go-ibooking/internal/infrastructure/router/admin"
	routerConsole "go-ibooking/internal/infrastructure/router/console"
	"go-ibooking/internal/registry"

	"github.com/gofiber/fiber/v2"
)

func InitializeRoute(fiber *fiber.App, r registry.Registry) *fiber.App {
	routerAdmin.InitializeAdminRoute(fiber, r.NewAdminController(), r.NewAdminMiddleware())
	routerConsole.InitializeConsoleRoute(fiber, r.NewConsoleController(), r.NewAdminMiddleware())

	return fiber
}
