package repository

import (
	"context"

	"github.com/pudthaiiii/golang-cms/src/app/model"
	"gorm.io/gorm"
)

type usageItemRepository struct {
	db *gorm.DB
}

type UsageItemRepository interface {
	SelectAll(ctx context.Context) ([]model.UsageItem, error)
}

func NewUsageItemRepository(db *gorm.DB) UsageItemRepository {
	return &usageItemRepository{db}
}

func (repo *usageItemRepository) SelectAll(ctx context.Context) ([]model.UsageItem, error) {
	var usageItems []model.UsageItem

	query := repo.db.WithContext(ctx).Unscoped()
	query = query.Select("id", "org_code", "product_code")
	query = query.Where("id = ?", "1")
	query = query.Where("org_code = ? OR product_code = ?", "PROD001", "PROD001")
	query = query.Limit(10000)
	query = query.Order("id desc")

	if err := query.Find(&usageItems).Error; err != nil {
		return nil, err
	}

	return usageItems, nil
}
