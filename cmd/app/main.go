package main

import (
	"go-ibooking/config"
	"go-ibooking/src/cmd"
	"go-ibooking/src/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize environment and config
	cfg, err := cmd.InitializeEnv()
	if err != nil {
		logger.Write.Err(err).Msg("Error loading configuration")
		return
	}
	logger.NewInitializeLogger()
	defer logger.CloseLogger()

	// Initialize and configure Fiber app
	app := setupFiber(cfg)

	// Create and bootstrap the application
	application := cmd.NewApplication(app, cfg)
	application.Boot()
}

func setupFiber(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit:       cfg.BodyLimit,
		ReadBufferSize:  cfg.BodyLimit,
		WriteBufferSize: cfg.BodyLimit,
		ReadTimeout:     time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(cfg.WriteTimeout) * time.Second,
		ErrorHandler:    ApiResource.ErrorHandler,
	})

	return app
}
