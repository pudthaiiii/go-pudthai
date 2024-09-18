package main

import (
	"fmt"
	"go-ibooking/config"
	"go-ibooking/src/pkg"
	"go-ibooking/src/pkg/logger"
	"go-ibooking/src/registry"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type App struct {
	route  *fiber.App
	config *config.Config
}

func NewApplication(route *fiber.App, cfg *config.Config) *App {
	return &App{route: route, config: cfg}
}

func (app *App) Boot() {
	app.setup()
	app.listen()
}

func (app *App) setup() {
	app.route.Get("/swagger/*", swagger.HandlerDefault)

	// Add middleware
	app.addMiddleware()

	// Initialize registry and setup routes
	// r := app.setupRegistry()

	logger.Write.Info().Msg("Successfully connected to Kafka")

	// Initialize routes
	// admin.InitializeAdminRoute(app.route, r.NewAdminController(), r.NewAdminMiddleware())
	// backend.InitializeBackendRoute(app.route, r.NewBackendController(), r.NewAdminMiddleware())
}

func (app *App) addMiddleware() {
	app.route.Use(func(c *fiber.Ctx) error {
		fmt.Println("Middleware center")
		return c.Next()
	})
}

func (app *App) setupRegistry() registry.Registry {
	db := pkg.NewPgDatastore()
	redis := pkg.NewRedisDatastore()
	s3 := pkg.NewS3Datastore()

	return registry.NewRegistry(db, redis.Client, s3)
}

func (app *App) listen() {
	port := app.config.Port

	logger.Write.Info().Msg(fmt.Sprintf("Server started on port %s", port))
	if err := app.route.Listen(":" + port); err != nil {
		logger.Write.Err(err).Msg("Server failed to start")
	}
}

func InitializeEnv() (*config.Config, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
