package repository

import (
	"context"
	"fmt"
	"go-ibooking/internal/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type oauthAccessTokenRepository struct {
	db *gorm.DB
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &oauthAccessTokenRepository{db: db}
}

type OauthAccessTokenRepository interface {
	CreateTransaction(ctx context.Context, userID uint, accessExpiresAt string, refreshExpiresAt string) (entities.OauthAccessToken, entities.OauthRefreshToken, error)
}

func (r *oauthAccessTokenRepository) CreateTransaction(ctx context.Context, userID uint, accessExpiresAt string, refreshExpiresAt string) (entities.OauthAccessToken, entities.OauthRefreshToken, error) {
	var accessToken entities.OauthAccessToken
	var refreshToken entities.OauthRefreshToken

	err := r.db.Transaction(func(tx *gorm.DB) error {
		newAccessUUID := uuid.New()

		accessDuration, err := time.ParseDuration(accessExpiresAt + "h")
		if err != nil {
			return fmt.Errorf("invalid access token expiresAt format: %w", err)
		}

		accessExpiredAt := time.Now().Add(accessDuration)

		accessToken = entities.OauthAccessToken{
			Token:     newAccessUUID.String(),
			ExpiresAt: &accessExpiredAt,
			UserID:    userID,
		}

		if err := tx.WithContext(ctx).Create(&accessToken).Error; err != nil {
			return fmt.Errorf("failed to create access token: %w", err)
		}

		newRefreshUUID := uuid.New()

		refreshDuration, err := time.ParseDuration(refreshExpiresAt + "h")
		if err != nil {
			return fmt.Errorf("invalid refresh token expiresAt format: %w", err)
		}

		refreshExpiredAt := time.Now().Add(refreshDuration)

		refreshToken = entities.OauthRefreshToken{
			Token:              newRefreshUUID.String(),
			ExpiresAt:          &refreshExpiredAt,
			OauthAccessTokenID: accessToken.ID,
		}

		if err := tx.WithContext(ctx).Create(&refreshToken).Error; err != nil {
			return fmt.Errorf("failed to create refresh token: %w", err)
		}

		return nil
	})

	if err != nil {
		return entities.OauthAccessToken{}, entities.OauthRefreshToken{}, err
	}

	return accessToken, refreshToken, nil
}
