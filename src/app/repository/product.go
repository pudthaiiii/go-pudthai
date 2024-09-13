package repository

import (
	"context"

	"go-ibooking/src/app/model"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	SelectAll(ctx context.Context) ([]model.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) SelectAll(ctx context.Context) ([]model.Product, error) {
	var productList []model.Product

	productBuilder := r.db.WithContext(ctx).Unscoped().
		Limit(1000).Find(&productList)

	if err := productBuilder.Error; err != nil {
		return nil, err
	}

	return productList, nil
}
