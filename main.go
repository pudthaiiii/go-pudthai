package main

import (
	ApiResource "go-ibooking/src/app/resources"
	"go-ibooking/src/utils"

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
	bodyLimit, _ := utils.CalFileSize("20mb")

	app := fiber.New(fiber.Config{
		BodyLimit:    int(bodyLimit),
		ErrorHandler: ApiResource.ErrorHandler,
	})

	return app
}

func deferClose() {
	log.CloseLogger()
}
