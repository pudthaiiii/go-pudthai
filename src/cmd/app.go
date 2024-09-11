package cmd

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	adminRouter "github.com/pudthaiiii/go-ibooking/src/app/router/admin"
	"github.com/pudthaiiii/go-ibooking/src/types"

	_ "github.com/pudthaiiii/go-ibooking/src/docs"
	"github.com/pudthaiiii/go-ibooking/src/pkg"
	"github.com/pudthaiiii/go-ibooking/src/registry"
	"github.com/pudthaiiii/go-ibooking/src/utils"
)

type App struct {
	route    *fiber.App
	dbConfig types.PGConfig
	config   map[string]interface{}
}

func NewApplication(route *fiber.App, dbConfig types.PGConfig, config map[string]interface{}) *App {
	return &App{route: route, dbConfig: dbConfig, config: config}
}

func (app *App) Boot() {
	app.setup()

	app.listen()
}

func (app *App) setup() {
	app.route.Get("/swagger/*", swagger.HandlerDefault)

	// center middleware
	app.newMiddleware()

	adminRouter.InitializeAdminRoute(app.route, app.newRegistry().NewAppController(), app.newRegistry().NewAdminMiddleware())
}

func (app *App) newMiddleware() {
	app.route.Use(func(c *fiber.Ctx) error {
		fmt.Println("Middleware center")
		return c.Next()
	})
}

func (app *App) newRegistry() registry.Registry {
	db := pkg.NewPgDatastore(app.dbConfig)

	return registry.NewRegistry(db.DB)
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
