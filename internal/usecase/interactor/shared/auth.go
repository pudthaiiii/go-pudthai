package interactor

import (
	"context"
	"go-ibooking/internal/config"
	"go-ibooking/internal/entities"
	"go-ibooking/internal/enum"
	"go-ibooking/internal/infrastructure/cache"
	"go-ibooking/internal/model/dtos"
	"go-ibooking/internal/throw"
	"go-ibooking/internal/usecase/repository"
	"go-ibooking/internal/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type sharedAuthInteractor struct {
	userRepo        repository.UsersRepository
	accessTokenRepo repository.OauthAccessTokenRepository
	cacheManager    *cache.CacheManager
	cfg             *config.Config
}

func NewSharedAuthInteractor(
	usersRepo repository.UsersRepository,
	accessTokenRepo repository.OauthAccessTokenRepository,
	cacheManager *cache.CacheManager,
	cfg *config.Config,
) SharedAuthInteractor {
	return &sharedAuthInteractor{
		usersRepo,
		accessTokenRepo,
		cacheManager,
		cfg,
	}
}

type SharedAuthInteractor interface {
	Login(ctx context.Context, dto dtos.Login, userType string) (dtos.AuthJWT, error)
}

// Login
func (s *sharedAuthInteractor) Login(ctx context.Context, dto dtos.Login, userType string) (dtos.AuthJWT, error) {
	user, userErr := s.userRepo.FindUserByEmail(ctx, dto.Email, userType)
	if userErr != nil {
		return dtos.AuthJWT{}, throw.UserCredentialMismatch()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return dtos.AuthJWT{}, throw.UserCredentialMismatch()
	}

	generateJwt, err := s.generateJwt(ctx, user)
	if err != nil {
		return dtos.AuthJWT{}, err
	}

	return generateJwt, nil
}

func (s *sharedAuthInteractor) generateJwt(ctx context.Context, user entities.User) (dtos.AuthJWT, error) {
	accessTokenExpiresAt := s.cfg.Get("JWT")["JwtAccessExpiresInHour"].(string)
	refreshTokenExpiresAt := s.cfg.Get("JWT")["JwtRefreshExpiresInHour"].(string)

	access, refresh, err := s.accessTokenRepo.CreateTransaction(ctx, user.ID, accessTokenExpiresAt, refreshTokenExpiresAt)
	if err != nil {
		return dtos.AuthJWT{}, throw.GenerateJwtTokenError(err)
	}

	accessToken, err := s.signJwt(map[string]interface{}{
		"token": access.Token,
	}, user.Type, accessTokenExpiresAt)
	if err != nil {
		return dtos.AuthJWT{}, throw.GenerateJwtTokenError(err)
	}

	refreshToken, err := s.signJwt(map[string]interface{}{
		"token": refresh.Token,
	}, user.Type, refreshTokenExpiresAt)
	if err != nil {
		return dtos.AuthJWT{}, throw.GenerateJwtTokenError(err)
	}

	return dtos.AuthJWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *sharedAuthInteractor) signJwt(payload map[string]interface{}, userType string, expiresIn string) (string, error) {
	var (
		err         error
		secret      string
		tokenString string
	)

	switch userType {
	case string(enum.USER):
		secret = s.cfg.Get("JWT")["JwtSecret"].(string)
	case string(enum.ADMIN):
		secret = s.cfg.Get("JWT")["JwtSecretAdmin"].(string)
	case string(enum.MERCHANT):
		secret = s.cfg.Get("JWT")["JwtSecretBackend"].(string)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(utils.StringToInt(expiresIn))).Unix(),
		"iat": time.Now().Unix(),
	})

	claims := token.Claims.(jwt.MapClaims)
	for key, value := range payload {
		claims[key] = value
	}

	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", throw.GenerateJwtTokenError(err)
	}

	return tokenString, nil
}
