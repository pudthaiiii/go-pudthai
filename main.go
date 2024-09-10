package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pudthaiiii/golang-cms/src/cmd"

	resource "github.com/pudthaiiii/golang-cms/src/resource"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: resource.ErrorHandler,
	})

	config := map[string]interface{}{
		"app": "Golang CMS",
	}

	application := cmd.NewApplication(app, config)
	application.InitializeEnv()
	application.Boot()
}
