package registry

import (
	mc "github.com/pudthaiiii/go-ibooking/src/app/http"
	am "github.com/pudthaiiii/go-ibooking/src/app/middleware/admin"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() mc.AppController
	NewAdminMiddleware() am.Middleware
}

func NewRegistry(
	db *gorm.DB,
) Registry {
	return &registry{
		db,
	}
}

func (r *registry) NewAppController() mc.AppController {
	ac := mc.AppController{
		AdminPrototype: r.NewPrototypeController(),

		// add more controller here
	}

	return ac
}

func (r *registry) NewAdminMiddleware() am.Middleware {
	return am.NewMiddleware()
}
