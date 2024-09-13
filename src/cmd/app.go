package cmd

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	adminRouter "go-ibooking/src/app/router/admin"
	backendRouter "go-ibooking/src/app/router/backend"

	_ "go-ibooking/docs"
	"go-ibooking/src/pkg"
	"go-ibooking/src/registry"
	"go-ibooking/src/utils"
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
	port := utils.RequireEnv("PORT", "3000")

	log.Printf("Server started on port %s", port)
	if err := app.route.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func InitializeEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}
