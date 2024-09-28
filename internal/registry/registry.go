package registry

import (
	cc "go-ibooking/internal/adapter/console/controllers"
	cm "go-ibooking/internal/adapter/console/middleware"
	sm "go-ibooking/internal/adapter/shared/middleware"
	ac "go-ibooking/internal/adapter/v1/admin/controllers"
	ab "go-ibooking/internal/adapter/v1/backend/controllers"
	af "go-ibooking/internal/adapter/v1/frontend/controllers"
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
	NewAdminController() ac.AdminController
	NewBackendController() ab.BackendController
	NewFrontendController() af.FrontendController
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

// Console Controllers
func (r *registry) NewConsoleController() cc.ConsoleController {
	ac := cc.ConsoleController{
		DatabaseController: r.NewConsoleDatabaseController(),
	}

	return ac
}

// Admin Controllers
func (r *registry) NewAdminController() ac.AdminController {
	ac := ac.AdminController{
		AuthController:  r.NewAdminAuthController(),
		UsersController: r.NewAdminUsersController(),
	}

	return ac
}

// Backend Controllers
func (r *registry) NewBackendController() ab.BackendController {
	ac := ab.BackendController{
		AuthController: r.NewBackendAuthController(),
	}

	return ac
}

// Frontend Controllers
func (r *registry) NewFrontendController() af.FrontendController {
	ac := af.FrontendController{
		AuthController: r.NewFrontendAuthController(),
	}

	return ac
}

// Console Middleware
func (r *registry) NewConsoleMiddleware() cm.Middleware {
	return cm.NewConsoleMiddleware(r.cfg)
}

// Shared Middleware
func (r *registry) NewSharedMiddleware() sm.Middleware {
	return sm.NewSharedMiddleware(r.cfg, r.cacheManager, r.db)
}
