package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func JwtSign(payload map[string]interface{}, expiresIn string, secret string) (string, error) {
	expires := StringToInt(expiresIn)

	tokenG := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * time.Duration(expires)).Unix(),
		"iat": time.Now().Unix(),
	})

	// Add payload to the claims
	claims := tokenG.Claims.(jwt.MapClaims)
	for key, value := range payload {
		claims[key] = value
	}

	// Sign the token with the secret
	tokenString, err := tokenG.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
