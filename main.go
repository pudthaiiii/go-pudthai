package main

import (
	"go-ibooking/src/cmd"
	log "go-ibooking/src/pkg/logger"
	resource "go-ibooking/src/resource"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cmd.InitializeEnv()
	log.NewInitializeLogger()

	defer deferClose()

	app := initFiberRouter()

	application := cmd.NewApplication(app)

	application.Boot()
}

func initFiberRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: resource.ErrorHandler,
	})

	return app
}

func deferClose() {
	log.CloseLogger()
}
