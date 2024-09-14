package main

import (
	ApiResource "go-ibooking/src/app/resources"

	log "go-ibooking/src/pkg/logger"

	"go-ibooking/src/cmd"

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
		ErrorHandler: ApiResource.ErrorHandler,
	})

	return app
}

func deferClose() {
	log.CloseLogger()
}
