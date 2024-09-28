package middleware

import (
	"go-ibooking/internal/config"

	"github.com/gofiber/fiber/v2"
)

type middleware struct {
	cfg *config.Config
}

func NewConsoleMiddleware(cfg *config.Config) Middleware {
	return &middleware{
		cfg,
	}
}

type Middleware interface {
	CookieAuthenticate(handler fiber.Handler, action string, subject string) fiber.Handler
}
