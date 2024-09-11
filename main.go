package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pudthaiiii/go-ibooking/src/cmd"

	dbConfig "github.com/pudthaiiii/go-ibooking/src/config/database"
	resource "github.com/pudthaiiii/go-ibooking/src/resource"
)

func main() {
	cmd.InitializeEnv()

	app := fiber.New(fiber.Config{
		ErrorHandler: resource.ErrorHandler,
	})

	config := map[string]interface{}{
		"dbConfig": "dbConfig",
	}

	application := cmd.NewApplication(app, dbConfig.GetPGConfig(), config)

	application.Boot()
}
