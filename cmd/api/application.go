package api

import (
	"go-ibooking/internal/config"
	"go-ibooking/internal/events"
	"go-ibooking/internal/infrastructure/cache"
	"go-ibooking/internal/infrastructure/datastore"
	"go-ibooking/internal/infrastructure/logger"
	"go-ibooking/internal/infrastructure/mailer"
	"go-ibooking/internal/infrastructure/recaptcha"
	"go-ibooking/internal/infrastructure/router"
	"go-ibooking/internal/registry"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type application struct {
	fiber *fiber.App
	cfg   *config.Config
}

func NewApiApplication() *application {
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

func (app *application) Config() *config.Config {
	return app.cfg
}

func (app *application) Boot() {
	app.loadLogger()
	app.loadRouter()

	app.listen()
}

func (app *application) loadRouter() {
	r := app.setupRegistry()

	router.InitializeRoute(app.fiber, r)
}

func (app *application) loadLogger() {
	logger.NewInitializeLogger(app.cfg)
}

func (app *application) registryListener(db *gorm.DB, cacheManager *cache.CacheManager) events.EventListener {
	mailer := mailer.NewMailer(app.cfg)
	listener := events.NewEventListener(mailer, db, cacheManager)

	go listener.Listen()

	return listener
}

func (app *application) setupRegistry() registry.Registry {
	s3 := datastore.NewS3Datastore(app.cfg)
	db := datastore.NewPgDatastore(app.cfg)
	cfg := app.cfg
	redis := datastore.NewRedisDatastore(app.cfg)
	recaptcha := recaptcha.NewRecaptchaProvider(app.cfg)
	cacheManager := cache.NewCacheManager(redis.Client, 5*time.Minute)

	listener := app.registryListener(db, cacheManager)

	return registry.NewRegistry(db, redis.Client, s3, cfg, recaptcha, cacheManager, listener)
}

func (app *application) listen() {
	port := app.cfg.Get("FiberConfig")["Port"].(string)

	if err := app.fiber.Listen(":" + port); err != nil {
		log.Printf("Server failed to start %v", err)
	}
}

func (app *application) DeferClose() {
	defer logger.CloseLogger()
}
