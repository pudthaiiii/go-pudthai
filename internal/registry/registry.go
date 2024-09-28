package registry

import (
	cc "go-ibooking/internal/adapter/console/controllers"
	cm "go-ibooking/internal/adapter/console/middleware"
	sm "go-ibooking/internal/adapter/shared/middleware"
	ca "go-ibooking/internal/adapter/v1/admin/controllers"
	"go-ibooking/internal/config"
	"go-ibooking/internal/events"
	"go-ibooking/internal/infrastructure/cache"
	"go-ibooking/internal/infrastructure/datastore"
	"go-ibooking/internal/infrastructure/recaptcha"

	"gorm.io/gorm"
)

type registry struct {
	db           *gorm.DB
	s3           *datastore.S3Datastore
	cfg          *config.Config
	recaptcha    *recaptcha.RecaptchaProvider
	cacheManager *cache.CacheManager
	listener     events.EventListener
}

type Registry interface {
	NewAdminController() ca.AdminController
	NewConsoleController() cc.ConsoleController
	NewConsoleMiddleware() cm.Middleware
	NewSharedMiddleware() sm.Middleware
}

func NewRegistry(
	db *gorm.DB,
	s3 *datastore.S3Datastore,
	cfg *config.Config,
	recaptcha *recaptcha.RecaptchaProvider,
	cacheManager *cache.CacheManager,
	listener events.EventListener,
) Registry {
	return &registry{
		db:           db,
		s3:           s3,
		cfg:          cfg,
		recaptcha:    recaptcha,
		cacheManager: cacheManager,
		listener:     listener,
	}
}

func (r *registry) NewConsoleController() cc.ConsoleController {
	ac := cc.ConsoleController{
		DatabaseController: r.NewConsoleDatabaseController(),
	}

	return ac
}

func (r *registry) NewAdminController() ca.AdminController {
	ac := ca.AdminController{
		AuthController:  r.NewAdminAuthController(),
		UsersController: r.NewAdminUsersController(),
	}

	return ac
}

func (r *registry) NewConsoleMiddleware() cm.Middleware {
	return cm.NewConsoleMiddleware(r.cfg)
}

func (r *registry) NewSharedMiddleware() sm.Middleware {
	return sm.NewSharedMiddleware(r.cfg, r.cacheManager, r.db)
}
