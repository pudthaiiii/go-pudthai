package middleware

import (
	"context"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *middleware) RequiredMerchant(handler fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !strings.Contains(c.Path(), "admin") {
			return handler(c)
		}

		merchantID := c.Get("Merchant-Id")
		if merchantID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error",
				"error":  "Merchant-Id header is required",
			})
		}

		merchantIDUint, err := strconv.ParseUint(merchantID, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "error",
				"error":  "Invalid Merchant-Id",
			})
		}

		merchant, err := m.merchantRepo.FindByID(c.Context(), uint(merchantIDUint))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "error",
				"error":  "Merchant not found",
			})
		}

		c.Locals("Merchant", merchant)
		c.Locals("MerchantID", merchantID)

		ctx := context.WithValue(c.Context(), "MerchantID", merchantID)
		ctx = context.WithValue(c.Context(), "Merchant", merchant)

		c.SetUserContext(ctx)

		return handler(c)
	}
}
