package middleware

import (
	"fmt"
	t "go-pudthai/internal/model/technical"
	"go-pudthai/internal/throw"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func (m *middleware) Authenticate(handler fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if isAuthenticate(action) {
			return handler(c)
		}

		tokenString := c.Get("Authorization")
		if err := m.validateTokenString(tokenString); err != nil {
			return err
		}

		tokenString = tokenString[7:]
		secret := m.getSecret(c.Route().Path)

		fmt.Println(secret, c.Route().Path)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, throw.InvalidJwtToken(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return m.unauthorizedResponse(c)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims)
		}

		return handler(c)
	}
}

func (m *middleware) validateTokenString(tokenString string) error {
	if tokenString == "" {
		return m.unauthorizedResponse(nil)
	}

	if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
		return m.unauthorizedResponse(nil)
	}

	return nil
}

func (m *middleware) unauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"code":    900401,
		"message": "UNAUTHORIZED",
	})
}

func isAuthenticate(action string) bool {
	return action == string(t.NONE)
}

func (m *middleware) getSecret(path string) string {
	var secretKey string

	switch {
	case strings.Contains(path, "/v1/admin"):
		secretKey = m.cfg.Get("JWT")["JwtSecretAdmin"].(string)
	case strings.Contains(path, "/v1/backend"):
		secretKey = m.cfg.Get("JWT")["JwtSecretBackend"].(string)
	case strings.Contains(path, "/v1/frontend"):
		secretKey = m.cfg.Get("JWT")["JwtSecret"].(string)
	default:
		secretKey = m.cfg.Get("JWT")["JwtSecret"].(string)
	}

	return secretKey
}
