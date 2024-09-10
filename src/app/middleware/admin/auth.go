package middlewareAdmin

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (middleware) Authenticate(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Your middleware logic here
		// Call the next handler
		fmt.Println("Middleware: Admin Authentication")
		return next(c)
	}
}
