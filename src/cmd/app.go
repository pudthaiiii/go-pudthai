package cmd

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "go-ibooking/docs"

	adminRouter "go-ibooking/src/app/router/admin"
	backendRouter "go-ibooking/src/app/router/backend"
	"go-ibooking/src/pkg"
	"go-ibooking/src/pkg/logger"
	"go-ibooking/src/registry"
)

type App struct {
	route *fiber.App
}

func NewApplication(route *fiber.App) *App {
	return &App{route: route}
}

func (app *App) Boot() {
	app.setup()

	app.listen()
}

func (app *App) setup() {
	app.route.Get("/swagger/*", swagger.HandlerDefault)

	// center middleware
	app.newMiddleware()

	// apply registry
	r := app.newRegistry()

	logger.Write.Info().Msg("Successfully connected to Kafka")

	// apply router
	adminRouter.InitializeAdminRoute(app.route, r.NewAdminController(), r.NewAdminMiddleware())
	backendRouter.InitializeBackendRoute(app.route, r.NewBackendController(), r.NewAdminMiddleware())
}

func (app *App) newMiddleware() {
	app.route.Use(func(c *fiber.Ctx) error {
		fmt.Println("Middleware center")
		return c.Next()
	})
}

func (app *App) newRegistry() registry.Registry {
	db := pkg.NewPgDatastore()
	redis := pkg.NewRedisDatastore()
	s3 := pkg.NewS3Datastore()

	return registry.NewRegistry(db, redis.Client, s3)
}

func (app *App) listen() {
	port := os.Getenv("PORT")

	logger.Write.Info().Msg(fmt.Sprintf("Server started on port %s", port))
	if err := app.route.Listen(":" + port); err != nil {
		logger.Write.Err(err).Msg("Server failed to start")
	}
}

func InitializeEnv() {
	godotenv.Load()
}
