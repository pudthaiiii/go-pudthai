package middleware

import (
	"go-pudthai/internal/throw"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
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

		_, err = m.merchantRepo.FindByID(c.Context(), uint(merchantIDUint))
		if err != nil {
			return throw.MerchantNotFound()
		}

		return handler(c)
	}
}
