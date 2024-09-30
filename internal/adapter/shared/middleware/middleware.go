package middleware

import (
	"go-pudthai/internal/config"
	"go-pudthai/internal/infrastructure/cache"
	"go-pudthai/internal/infrastructure/recaptcha"
	"go-pudthai/internal/usecase/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type middleware struct {
	merchantRepo    repository.MerchantsRepository
	accessTokenRepo repository.OauthAccessTokenRepository
	recaptcha       *recaptcha.RecaptchaProvider
	cfg             *config.Config
}

func NewSharedMiddleware(
	cfg *config.Config,
	cacheManager *cache.CacheManager,
	db *gorm.DB,
	recaptcha *recaptcha.RecaptchaProvider,
) Middleware {
	return &middleware{
		merchantRepo:    repository.NewMerchantsRepository(db),
		accessTokenRepo: repository.NewOauthAccessTokenRepository(db),
		recaptcha:       recaptcha,
		cfg:             cfg,
	}
}

type Middleware interface {
	Authenticate(handler fiber.Handler, action string, subject string) fiber.Handler
	RequiredMerchant(handler fiber.Handler, action string, subject string) fiber.Handler
	GoogleRecaptcha(handler fiber.Handler, action string, subject string) fiber.Handler
}
