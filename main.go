package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pudthaiiii/go-ibooking/src/cmd"

	resource "github.com/pudthaiiii/go-ibooking/src/resource"
)

func main() {
	app := initFiberRouter()

	application := cmd.NewApplication(app)

	application.Boot()
}

func initFiberRouter() *fiber.App {
	cmd.InitializeEnv()

	app := fiber.New(fiber.Config{
		ErrorHandler: resource.ErrorHandler,
	})

	return app
}
