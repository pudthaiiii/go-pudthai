package bootstrap

import (
	"log"

	ApiResource "github.com/pudthaiiii/golang-cms/src/resource"
	"github.com/pudthaiiii/golang-cms/src/utils"

	"github.com/gofiber/fiber/v2"
)

func routerInit() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: ApiResource.ErrorHandler,
	})

	return app
}

func startListen(app *fiber.App) {
	port := utils.RequireEnv("PORT", "3000")
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
