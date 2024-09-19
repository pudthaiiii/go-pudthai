package registry

import (
	"go-ibooking/internal/infrastructure/datastore"
	a "go-ibooking/src/app/http/admin"

	"github.com/go-redis/redis/v8"

	am "go-ibooking/src/app/http/middleware/admin"

	consoleController "go-ibooking/internal/adapter/controllers/v1/console"

	"gorm.io/gorm"
)

type registry struct {
	db    *gorm.DB
	redis redis.UniversalClient
	s3    *datastore.S3Datastore
}

type Registry interface {
	NewAdminController() a.AdminController
	NewConsoleController() consoleController.ConsoleController
	NewAdminMiddleware() am.Middleware
}

func NewRegistry(
	db *gorm.DB,
	redisClient redis.UniversalClient,
	s3 *datastore.S3Datastore,
) Registry {
	return &registry{
		db:    db,
		redis: redisClient,
		s3:    s3,
	}
}

func (r *registry) NewAdminController() a.AdminController {
	ac := a.AdminController{
		RoleController:  r.RegisterRoleController(),
		UsersController: r.RegisterUsersController(),
		AuthController:  r.RegisterAuthController(),
		// add more controller here
	}

	return ac
}

func (r *registry) NewConsoleController() consoleController.ConsoleController {
	ac := consoleController.ConsoleController{
		FeaturesController: r.RegisterFeaturesController(),
		// add more controller here
	}

	return ac
}

func (r *registry) NewAdminMiddleware() am.Middleware {
	return am.NewMiddleware(r.db, r.redis)
}
