package registry

import (
	admin "github.com/pudthaiiii/go-ibooking/src/app/http/admin"
	backend "github.com/pudthaiiii/go-ibooking/src/app/http/backend"

	am "github.com/pudthaiiii/go-ibooking/src/app/middleware/admin"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAdminController() admin.AdminController
	NewBackendController() backend.BackendController
	NewAdminMiddleware() am.Middleware
}

func NewRegistry(
	db *gorm.DB,
) Registry {
	return &registry{
		db,
	}
}

func (r *registry) NewAdminController() admin.AdminController {
	ac := admin.AdminController{
		AdminPrototype: r.NewPrototypeController(),

		// add more controller here
	}

	return ac
}

func (r *registry) NewBackendController() backend.BackendController {
	ac := backend.BackendController{
		PrototypeController: r.NewPrototypeController(),

		// add more controller here
	}

	return ac
}

func (r *registry) NewAdminMiddleware() am.Middleware {
	return am.NewMiddleware()
}
