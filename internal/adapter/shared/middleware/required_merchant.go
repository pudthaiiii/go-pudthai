package middleware

import (
	"context"
	"fmt"
	"go-ibooking/internal/throw"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type contextKey string

const (
	merchantIDKey contextKey = "MerchantID"
	merchantKey   contextKey = "Merchant"
	testMM        contextKey = "TestMM"
)

func (m *middleware) RequiredMerchant(handler fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if strings.Contains(c.Path(), "/v1/admin") {
			return handler(c)
		}

		merchantID := c.Get("Merchant-Id")
		if merchantID == "" {
			return throw.MerchantNotFound()
		}

		merchantIDUint, err := strconv.ParseUint(merchantID, 10, 32)
		if err != nil {
			return throw.MerchantNotFound()
		}

		merchant, err := m.merchantRepo.FindByID(c.Context(), uint(merchantIDUint))
		if err != nil {
			return throw.MerchantNotFound()
		}

		c.Locals("Merchant", merchant)
		c.Locals("MerchantID", merchantID)

		ctx := context.WithValue(c.Context(), merchantIDKey, merchantID)
		ctx = context.WithValue(ctx, merchantKey, merchant)

		c.SetUserContext(ctx)

		return handler(c)
	}
}

func (m *middleware) Authenticate(handler fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user from context
		fmt.Println("Get user from context", action, subject)

		c.Locals(testMM, "Test1")
		ctx := context.WithValue(c.Context(), testMM, "Test1")

		c.SetUserContext(ctx)
		return handler(c)
	}
}
