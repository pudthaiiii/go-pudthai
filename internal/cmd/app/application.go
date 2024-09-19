package app

import (
	"go-ibooking/internal/adapter/resources"
	"go-ibooking/internal/config"
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
		ErrorHandler:    resources.ErrorHandler,
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

	return registry.NewRegistry(db, redis.Client, s3)
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

// func (app *application) setup() {
// 	app.route.Get("/swagger/*", swagger.HandlerDefault)

// 	// Add middleware
// 	app.addMiddleware()

// 	// Initialize registry and setup routes
// 	// r := app.setupRegistry()

// 	logger.Write.Info().Msg("Successfully connected to Kafka")

// 	// Initialize routes
// 	// admin.InitializeAdminRoute(app.route, r.NewAdminController(), r.NewAdminMiddleware())
// 	// backend.InitializeBackendRoute(app.route, r.NewBackendController(), r.NewAdminMiddleware())
// }

// func (app *application) addMiddleware() {
// 	app.route.Use(func(c *fiber.Ctx) error {
// 		fmt.Println("Middleware center")
// 		return c.Next()
// 	})
// }

// func (app *application) SetupRegistry() registry.Registry {
// 	db := pkg.NewPgDatastore()
// 	redis := pkg.NewRedisDatastore()
// 	s3 := pkg.NewS3Datastore()

// 	return registry.NewRegistry(db, redis.Client, s3)
// }

// func (app *application) listen() {
// 	port := app.config.Port

// 	logger.Write.Info().Msg(fmt.Sprintf("Server started on port %s", port))
// 	if err := app.route.Listen(":" + port); err != nil {
// 		logger.Write.Err(err).Msg("Server failed to start")
// 	}
// }
