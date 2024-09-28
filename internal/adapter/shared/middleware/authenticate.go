package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (m *middleware) Authenticate(handler fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user from context
		fmt.Println("Get user from context", action, subject)

		return handler(c)
	}
}
