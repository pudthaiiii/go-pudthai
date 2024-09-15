package registry

import (
	a "go-ibooking/src/app/http/admin"
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
	NewAdminController() a.AdminController
	NewBackendController() backend.BackendController
	NewAdminMiddleware() am.Middleware
}

func NewRegistry(
	db *gorm.DB,
	redisClient redis.UniversalClient,
	s3 *pkg.S3Datastore,
) Registry {
	return &registry{
		db:    db,
		redis: redisClient,
		s3:    s3,
	}
}

func (r *registry) NewAdminController() a.AdminController {
	ac := a.AdminController{
		PrototypeController: r.RegisterPrototypeController(),
		RoleController:      r.RegisterRoleController(),
		UsersController:     r.RegisterUsersController(),
		// add more controller here
	}

	return ac
}

func (r *registry) NewBackendController() backend.BackendController {
	ac := backend.BackendController{
		PrototypeController: r.RegisterPrototypeController(),

		// add more controller here
	}

	return ac
}

func (r *registry) NewAdminMiddleware() am.Middleware {
	return am.NewMiddleware()
}
