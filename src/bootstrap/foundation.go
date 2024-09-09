package bootstrap

import (
	router "workshop/src/app/router"

	"workshop/src/pkg"
	"workshop/src/registry"

	"github.com/gofiber/fiber/v2"
)

func start(app *fiber.App) {
	db := pkg.ConnectPgSql()

	// package registry
	r := registry.NewRegistry(
		db.DB,
	)
	app = router.InitializeRoute(app, r.NewAppController(), r.NewMiddleWare())

	startListen(app)
}
