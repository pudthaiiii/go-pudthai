package registry

import (
	"github.com/pudthaiiii/golang-cms/src/app/controller"
	"github.com/pudthaiiii/golang-cms/src/app/middleware"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
	NewMiddleWare() middleware.Middleware
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
		PrototypeController: r.NewPrototypeController(),
		MerchantsController: r.NewMerchantsController(),

		// add more controller here
	}

	return ac
}

func (r *registry) NewMiddleWare() middleware.Middleware {
	return middleware.NewMiddleware()
}
