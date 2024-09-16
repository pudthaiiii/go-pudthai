package middlewareAdmin

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func (m *middleware) Authenticate(next fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		action, subject = strings.ToLower(action), strings.ToLower(subject)

		if isAuthWithoutToken(action, subject) {
			return next(c)
		}

		if isInvalidToken(tokenString) {
			return unauthorizedException(c)
		}

		if err := validateToken(tokenString); err != nil {
			return unauthorizedException(c)
		}

		return next(c)
	}
}

// func (m *middleware) getUserFromJwtToken() {
// 	// m.db.Model(&model.User{}).Where("id = ?", token.Claims.(jwt.MapClaims)["id"]).First(&user)
// }

// func (m *middleware) checkPermission() {
// 	// m.db.Model(&model.User{}).Where("id = ?", token.Claims.(jwt.MapClaims)["id"]).First(&user)
// }

func isAuthWithoutToken(action, subject string) bool {
	return action == "login" || subject == "auth"
}

func isInvalidToken(token string) bool {
	return token == "" || !strings.HasPrefix(token, "Bearer ")
}

func validateToken(tokenString string) error {
	bearerToken := strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_ADMIN_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return err
	}

	return nil
}

func unauthorizedException(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    900401,
		"message": "UNAUTHORIZED",
	})
}
