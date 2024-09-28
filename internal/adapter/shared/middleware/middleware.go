package middleware

import (
	"go-ibooking/internal/config"
	"go-ibooking/internal/infrastructure/cache"
	"go-ibooking/internal/usecase/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type middleware struct {
	merchantRepo repository.MerchantsRepository
}

func NewSharedMiddleware(
	cfg *config.Config,
	cacheManager *cache.CacheManager,
	db *gorm.DB,
) Middleware {
	return &middleware{
		merchantRepo: repository.NewMerchantsRepository(db),
	}
}

type Middleware interface {
	Authenticate(handler fiber.Handler, action string, subject string) fiber.Handler
	RequiredMerchant(handler fiber.Handler, action string, subject string) fiber.Handler
}
