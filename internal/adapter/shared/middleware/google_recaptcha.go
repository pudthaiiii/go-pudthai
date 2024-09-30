package middleware

import (
	"go-pudthai/internal/throw"
	"go-pudthai/internal/utils"

	"github.com/gofiber/fiber/v2"
)

var routeNotAllow = []string{"Login", "Refresh"}

func (m *middleware) GoogleRecaptcha(handler fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		enabled := m.cfg.Get("GoogleRecaptcha")["RecaptchaEnabled"].(string)
		if !utils.StringToBool(enabled) {
			return handler(c)
		}

		if !isAuthRoute(c.Route().Name) {
			return handler(c)
		}

		token := c.Get("Recaptcha")
		if token == "" {
			return throw.ValidateRecaptchaError()
		}

		if err := m.verifyRecaptchaToken(token); err != nil {
			return err
		}

		return handler(c)
	}
}

func (m *middleware) verifyRecaptchaToken(token string) error {
	ok, err := m.recaptcha.VerifyToken(token)
	if err != nil {
		return throw.RecaptchaError()
	}

	if !ok {
		return throw.RecaptchaError()
	}

	return nil
}

func isAuthRoute(routeName string) bool {
	for _, route := range routeNotAllow {
		if route == routeName {
			return true
		}
	}

	return false
}
