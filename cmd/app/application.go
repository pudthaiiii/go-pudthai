package app

import (
	"go-ibooking/config"
	"go-ibooking/internal/infrastructure/datastore"
	"go-ibooking/internal/infrastructure/logger"
	"go-ibooking/internal/infrastructure/router"
	"go-ibooking/internal/registry"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type application struct {
	fiber *fiber.App
	cfg   *config.Config
}

func NewApplication() *application {
	config, err := config.NewConfig()
	if err != nil {
		log.Printf("unable to load config, %v", err)
	}

	config.Initialize()

	fiberConfig := fiber.Config{
		BodyLimit:       config.Get("FiberConfig")["BodyLimit"].(int),
		ReadBufferSize:  config.Get("FiberConfig")["ReadBufferSize"].(int),
		WriteBufferSize: config.Get("FiberConfig")["WriteBufferSize"].(int),
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		ErrorHandler:    errorHandler,
	}

	return &application{
		cfg:   config,
		fiber: fiber.New(fiberConfig),
	}
}

func (app *application) Fiber() *fiber.App {
	return app.fiber
}

func (app *application) Boot() {
	app.loadLogger()
	app.loadRouter()

	logger.Log.Info().Msg("Logger loaded")
}

// Load Router
func (app *application) loadRouter() {
	r := app.setupRegistry()

	router.InitializeRoute(app.fiber, r)
}

// Load Logger
func (app *application) loadLogger() {
	logger.NewInitializeLogger(app.cfg)
}

// setupRegistry
func (app *application) setupRegistry() registry.Registry {
	db := datastore.NewPgDatastore(app.cfg)
	redis := datastore.NewRedisDatastore(app.cfg)
	s3 := datastore.NewS3Datastore(app.cfg)
	cfg := app.cfg

	return registry.NewRegistry(db, redis.Client, s3, cfg)
}

// Passed
func (app *application) Listen() {
	port := app.cfg.Get("FiberConfig")["Port"].(string)

	if err := app.fiber.Listen(":" + port); err != nil {
		log.Printf("Server failed to start %v", err)
	}
}

func (app *application) DeferClose() {
	defer logger.CloseLogger()
}
