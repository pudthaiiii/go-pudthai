package repository

import (
	"context"
	"fmt"
	"go-ibooking/internal/entities"
	b "go-ibooking/internal/model/business"
	"time"

	"gorm.io/gorm"
)

type oauthRefreshTokenRepository struct {
	db *gorm.DB
}

func NewOauthRefreshTokenRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &oauthRefreshTokenRepository{db: db}
}

type OauthRefreshTokenRepository interface {
	FindByToken(ctx context.Context, token string) (b.RefreshTokenResult, error)
	DeleteByID(ctx context.Context, id uint) error
}

func (r *oauthRefreshTokenRepository) FindByToken(ctx context.Context, token string) (b.RefreshTokenResult, error) {
	var (
		refreshTokenResult b.RefreshTokenResult
	)

	query := r.db.WithContext(ctx).
		Joins("JOIN oauth_access_tokens ON oauth_access_tokens.id = oauth_refresh_tokens.oauth_access_token_id").
		Joins("JOIN users ON oauth_access_tokens.user_id = users.id").
		Select("oauth_refresh_tokens.id, oauth_refresh_tokens.token, oauth_refresh_tokens.expires_at, oauth_access_tokens.user_id, users.type").
		Where("oauth_refresh_tokens.token = ? AND oauth_refresh_tokens.expires_at > ?", token, time.Now()).
		First(&refreshTokenResult)
	if query.Error != nil {
		return b.RefreshTokenResult{}, query.Error
	}

	return refreshTokenResult, nil
}

func (r *oauthRefreshTokenRepository) DeleteByID(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&entities.OauthRefreshToken{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with id %d", id)
	}

	return nil
}
