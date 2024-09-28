package interactor

import (
	"context"
	"errors"
	"go-ibooking/internal/adapter/v1/admin/dtos"
	"go-ibooking/internal/config"
	"go-ibooking/internal/infrastructure/cache"
	"go-ibooking/internal/throw"
	"go-ibooking/internal/usecase/repository"
	"time"

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
