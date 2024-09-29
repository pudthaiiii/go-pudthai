package repository

import (
	"context"
	"go-pudthai/internal/entities"

	"gorm.io/gorm"
)

type merchantsRepository struct {
	db *gorm.DB
}

func NewMerchantsRepository(db *gorm.DB) MerchantsRepository {
	return &merchantsRepository{db: db}
}

type MerchantsRepository interface {
	FindByID(ctx context.Context, id uint) (entities.Merchant, error)
}

func (r *merchantsRepository) FindByID(ctx context.Context, id uint) (entities.Merchant, error) {
	var (
		merchant entities.Merchant
	)

	query := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&merchant)
	if query.Error != nil {
		return entities.Merchant{}, query.Error
	}

	return merchant, nil
}
