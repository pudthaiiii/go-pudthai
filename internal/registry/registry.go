package registry

import (
	"go-ibooking/internal/adapter/v1/controllers"
	cc "go-ibooking/internal/adapter/v1/controllers/console"
	"go-ibooking/internal/config"
	"go-ibooking/internal/events"
	"go-ibooking/internal/infrastructure/cache"
	"go-ibooking/internal/infrastructure/datastore"
	"go-ibooking/internal/infrastructure/recaptcha"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type registry struct {
	db           *gorm.DB
	redis        redis.UniversalClient
	s3           *datastore.S3Datastore
	cfg          *config.Config
	recaptcha    *recaptcha.RecaptchaProvider
	cacheManager *cache.CacheManager
	listener     events.EventListener
}

type Registry interface {
	NewController() controllers.AppController
	NewConsoleController() cc.ConsoleController
	// NewAdminMiddleware() am.Middleware
}

func NewRegistry(
	db *gorm.DB,
	redisClient redis.UniversalClient,
	s3 *datastore.S3Datastore,
	cfg *config.Config,
	recaptcha *recaptcha.RecaptchaProvider,
	cacheManager *cache.CacheManager,
	listener events.EventListener,
) Registry {
	return &registry{
		db:           db,
		redis:        redisClient,
		s3:           s3,
		cfg:          cfg,
		recaptcha:    recaptcha,
		cacheManager: cacheManager,
		listener:     listener,
	}
}

func (r *registry) NewController() controllers.AppController {
	ac := controllers.AppController{
		UsersController: r.RegisterUsersController(),
	}

	return ac
}

func (r *registry) NewConsoleController() cc.ConsoleController {
	ac := cc.ConsoleController{
		FeaturesController: r.RegisterFeaturesController(),
	}

	return ac
}

// func (r *registry) NewAdminMiddleware() am.Middleware {
// 	return am.NewMiddleware(r.db, r.redis)
// }
