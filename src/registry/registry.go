package registry

import (
	"github.com/pudthaiiii/golang-cms/src/app/controller"
	am "github.com/pudthaiiii/golang-cms/src/app/middleware/admin"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
	NewAdminMiddleware() am.Middleware
}

func NewRegistry(
	db *gorm.DB,
) Registry {
	return &registry{
		db,
	}
}

func (r *registry) NewAppController() controller.AppController {
	ac := controller.AppController{
		AdminPrototype: r.NewPrototypeController(),

		// add more controller here
	}

	return ac
}

func (r *registry) NewAdminMiddleware() am.Middleware {
	return am.NewMiddleware()
}
