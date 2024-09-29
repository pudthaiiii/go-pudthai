package middleware

import (
	"go-pudthai/internal/infrastructure/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (m *middleware) Log(handler fiber.Handler, name string, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		logger.Write.Info().
			Msgf("%s\t%s\t%s\t%s\t%s\t%s",
				c.Method(),
				c.OriginalURL(),
				name,
				time.Since(start),
				action,
				subject,
			)

		return handler(c)
	}
}
