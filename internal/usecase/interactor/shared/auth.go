package interactor

import (
	"context"
	"errors"
	"fmt"
	"go-ibooking/internal/adapter/v1/admin/dtos"
	"go-ibooking/internal/config"
	"go-ibooking/internal/enum"
	"go-ibooking/internal/infrastructure/cache"
	"go-ibooking/internal/throw"
	"go-ibooking/internal/usecase/repository"
	"go-ibooking/internal/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type sharedAuthInteractor struct {
	userRepo         repository.UsersRepository
	accessTokenRepo  repository.OauthAccessTokenRepository
	refreshTokenRepo repository.OauthRefreshTokenRepository
	cacheManager     *cache.CacheManager
	cfg              *config.Config
}

func NewSharedAuthInteractor(
	usersRepo repository.UsersRepository,
	accessTokenRepo repository.OauthAccessTokenRepository,
	refreshTokenRepo repository.OauthRefreshTokenRepository,
	cacheManager *cache.CacheManager,
	cfg *config.Config,
) SharedAuthInteractor {
	return &sharedAuthInteractor{
		usersRepo,
		accessTokenRepo,
		refreshTokenRepo,
		cacheManager,
		cfg,
	}
}

type SharedAuthInteractor interface {
	Login(ctx context.Context, dto dtos.Login, userType string) (dtos.AuthJWT, error)
	Refresh(ctx context.Context, dto dtos.Refresh, userType string) (dtos.AuthJWT, error)
}

// Login
func (s *sharedAuthInteractor) Login(ctx context.Context, dto dtos.Login, userType string) (dtos.AuthJWT, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, dto.Email, userType)
	if err != nil {
		return dtos.AuthJWT{}, throw.UserCredentialMismatch()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return dtos.AuthJWT{}, throw.UserCredentialMismatch()
	}

	return s.generateJwt(ctx, user.Type, user.ID)
}

// Refresh
func (s *sharedAuthInteractor) Refresh(ctx context.Context, dto dtos.Refresh, userType string) (dtos.AuthJWT, error) {
	claims, err := s.decodeJwt(dto.RefreshToken, userType)
	if err != nil {
		return dtos.AuthJWT{}, throw.InvalidJwtToken(err)
	}

	token := claims["token"].(string)

	refreshTokenResult, err := s.refreshTokenRepo.FindByToken(ctx, token)
	if err != nil {
		return dtos.AuthJWT{}, throw.InvalidJwtToken(err)
	}

	if time.Now().After(refreshTokenResult.ExpiresAt) {
		return dtos.AuthJWT{}, throw.InvalidJwtToken(errors.New("refreshTokenResult token has expired"))
	}

	generateJwt, err := s.generateJwt(ctx, refreshTokenResult.Type, refreshTokenResult.UserID)
	if err != nil {
		return dtos.AuthJWT{}, err
	}

	if err := s.refreshTokenRepo.DeleteByID(ctx, refreshTokenResult.ID); err != nil {
		return dtos.AuthJWT{}, throw.InvalidJwtToken(err)
	}

	return generateJwt, nil
}

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
