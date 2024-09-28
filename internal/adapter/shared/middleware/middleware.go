package middleware

import (
	"go-ibooking/internal/config"
	"go-ibooking/internal/infrastructure/cache"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type middleware struct {
	cfg          *config.Config
	cacheManager *cache.CacheManager
	db           *gorm.DB
}

func NewSharedMiddleware(
	cfg *config.Config,
	cacheManager *cache.CacheManager,
	db *gorm.DB,
) Middleware {
	return &middleware{
		cfg,
		cacheManager,
		db,
	}
}

type Middleware interface {
	Authenticate(handler fiber.Handler, action string, subject string) fiber.Handler
}
