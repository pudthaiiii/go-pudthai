package middleware

import (
	"go-ibooking/internal/enum/permission"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *middleware) CookieAuthenticate(handler fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if action == string(permission.COOKIE) && subject == string(permission.CONSOLE) {
			headerCookie := c.Get("Cookie", "")
			cookiePairs := strings.Split(headerCookie, "; ")
			cookieBag := make(map[string]string)

			for _, pair := range cookiePairs {
				kv := strings.SplitN(pair, ":", 2)
				if len(kv) == 2 {
					key := strings.TrimSpace(kv[0])
					value := strings.TrimSpace(kv[1])
					cookieBag[key] = value
				}
			}

			secret := m.cfg.Get("Cookie")["Secret"].(string)
			name := m.cfg.Get("Cookie")["Name"].(string)

			if cookieBag[name] != secret {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code":    100001,
					"message": "AUTH_CREDENTIAL_MISMATCH",
				})
			}
		}

		return handler(c)
	}
}
