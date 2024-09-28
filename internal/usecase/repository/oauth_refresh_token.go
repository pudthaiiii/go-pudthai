package repository

import (
	"context"
	"fmt"
	"go-ibooking/internal/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type oauthRefreshTokenRepository struct {
	db *gorm.DB
}

func NewOauthRefreshTokenRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &oauthRefreshTokenRepository{db: db}
}

type OauthRefreshTokenRepository interface {
	Create(ctx context.Context, accessTokenID uint, expiresAt string) (entities.OauthRefreshToken, error)
}

func (r *oauthRefreshTokenRepository) Create(ctx context.Context, accessTokenID uint, expiresAt string) (entities.OauthRefreshToken, error) {
	newUUID := uuid.New()

	duration, err := time.ParseDuration(expiresAt + "h")
	if err != nil {
		return entities.OauthRefreshToken{}, err
	}

	expiredAt := time.Now().Add(duration)

	var refreshToken = entities.OauthRefreshToken{
		Token:              newUUID.String(),
		ExpiresAt:          &expiredAt,
		OAuthAccessTokenID: accessTokenID,
	}

	fmt.Println("refreshToken", refreshToken)
	query := r.db.WithContext(ctx).Create(&refreshToken)
	if query.Error != nil {
		return entities.OauthRefreshToken{}, query.Error
	}

	return refreshToken, nil
}
