package interactor

import (
	"context"
	"errors"
	"fmt"
	"go-ibooking/internal/adapter/shared/dtos"
	"go-ibooking/internal/enum"
	"go-ibooking/internal/throw"
	"go-ibooking/internal/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

// generateJwt generates jwt token
func (s *sharedAuthInteractor) generateJwt(ctx context.Context, userType string, userID uint) (dtos.AuthJWT, error) {
	accessTokenExpiresAt := s.cfg.Get("JWT")["JwtAccessExpiresInHour"].(string)
	refreshTokenExpiresAt := s.cfg.Get("JWT")["JwtRefreshExpiresInHour"].(string)

	access, refresh, err := s.accessTokenRepo.CreateTransaction(ctx, userID, accessTokenExpiresAt, refreshTokenExpiresAt)
	if err != nil {
		return dtos.AuthJWT{}, throw.GenerateJwtTokenError(err)
	}

	accessToken, err := s.signJwt(map[string]interface{}{"token": access.Token}, userType, accessTokenExpiresAt)
	if err != nil {
		return dtos.AuthJWT{}, throw.GenerateJwtTokenError(err)
	}

	refreshToken, err := s.signJwt(map[string]interface{}{"token": refresh.Token}, userType, refreshTokenExpiresAt)
	if err != nil {
		return dtos.AuthJWT{}, throw.GenerateJwtTokenError(err)
	}

	return dtos.AuthJWT{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

// signJwt signs jwt token
func (s *sharedAuthInteractor) signJwt(payload map[string]interface{}, userType string, expiresIn string) (string, error) {
	secret := s.getSecret(userType)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(utils.StringToInt(expiresIn))).Unix(),
		"iat": time.Now().Unix(),
	})

	claims := token.Claims.(jwt.MapClaims)
	for key, value := range payload {
		claims[key] = value
	}

	return token.SignedString([]byte(secret))
}

// decodeJwt decodes jwt token
func (s *sharedAuthInteractor) decodeJwt(tokenString string, userType string) (map[string]interface{}, error) {
	secret := s.getSecret(userType)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, throw.InvalidJwtToken(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, throw.InvalidJwtToken(errors.New("invalid token"))
}

// getSecret gets jwt secret
func (s *sharedAuthInteractor) getSecret(userType string) string {
	switch userType {
	case string(enum.USER):
		return s.cfg.Get("JWT")["JwtSecret"].(string)
	case string(enum.ADMIN):
		return s.cfg.Get("JWT")["JwtSecretAdmin"].(string)
	case string(enum.MERCHANT):
		return s.cfg.Get("JWT")["JwtSecretBackend"].(string)
	}
	return ""
}
