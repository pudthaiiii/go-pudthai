package middlewareAdmin

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type middleware struct {
	db          *gorm.DB
	redisClient redis.UniversalClient
}

type Middleware interface {
	Authenticate(next fiber.Handler, action string, subject string) fiber.Handler
}

func NewMiddleware(db *gorm.DB, redisClient redis.UniversalClient) Middleware {
	return &middleware{
		db:          db,
		redisClient: redisClient,
	}
}
