package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func JwtSign(payload map[string]interface{}, expiresIn string, secret string) (string, error) {
	// Parsing the expiration time from string to integer
	expirationHours, err := strconv.Atoi(expiresIn)
	if err != nil {
		return "", fmt.Errorf("invalid expiresIn format: %v", err)
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(expirationHours) * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	// Access token claims safely
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Add additional claims from the payload
		for key, value := range payload {
			claims[key] = value
		}
	} else {
		return "", fmt.Errorf("unable to parse claims")
	}

	// Sign the token using the provided secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
