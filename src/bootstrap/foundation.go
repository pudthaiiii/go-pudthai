package bootstrap

import (
	router "github.com/pudthaiiii/golang-cms/src/app/router"

	"github.com/pudthaiiii/golang-cms/src/pkg"
	"github.com/pudthaiiii/golang-cms/src/registry"

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
