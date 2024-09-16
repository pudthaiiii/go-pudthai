package main

import (
	ApiResource "go-ibooking/src/app/resources"
	"go-ibooking/src/utils"
	"time"

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
	bodyLimit, _ := utils.CalFileSize("100mb")

	app := fiber.New(fiber.Config{
		BodyLimit:       int(bodyLimit),
		ReadBufferSize:  int(bodyLimit),
		WriteBufferSize: int(bodyLimit),
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		ErrorHandler:    ApiResource.ErrorHandler,
	})

	return app
}

func deferClose() {
	log.CloseLogger()
}
