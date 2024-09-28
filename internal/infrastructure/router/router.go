package router

import (
	"fmt"
	_ "go-ibooking/docs"
	ra "go-ibooking/internal/infrastructure/router/admin"
	rc "go-ibooking/internal/infrastructure/router/console"

	"go-ibooking/internal/registry"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func InitializeRoute(fiber *fiber.App, r registry.Registry) *fiber.App {
	fiber.Get("/swagger/*", swagger.HandlerDefault)

	appleMiddlewares(fiber)

	ra.InitializeAdminRoute(fiber, r.NewAdminController(), r.NewSharedMiddleware())

	rc.InitializeConsoleRoute(fiber, r.NewConsoleController(), r.NewConsoleMiddleware())

	return fiber
}

func appleMiddlewares(fiber *fiber.App) {
	fmt.Println("appleMiddlewares")

	fiber.Use(cors.New(cors.Config{
		AllowOrigins: "https://example.com",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE",
	}))
}
