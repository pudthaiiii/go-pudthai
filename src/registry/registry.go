package registry

import (
	admin "go-ibooking/src/app/http/admin"
	backend "go-ibooking/src/app/http/backend"
	"go-ibooking/src/pkg"

	"github.com/go-redis/redis/v8"

	am "go-ibooking/src/app/middleware/admin"

	"gorm.io/gorm"
)

type registry struct {
	db    *gorm.DB
	redis redis.UniversalClient
	s3    *pkg.S3Datastore
}

type Registry interface {
	NewAdminController() admin.AdminController
	NewBackendController() backend.BackendController
	NewAdminMiddleware() am.Middleware
}

func NewRegistry(
	db *gorm.DB,
	// redisClient redis.UniversalClient,
	s3 *pkg.S3Datastore,
) Registry {
	return &registry{
		db: db,
		// redis: redisClient,
		s3: s3,
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
