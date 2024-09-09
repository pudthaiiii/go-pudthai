package repository

import (
	"context"
	"workshop/src/app/entities"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepository interface {
	SelectAll(ctx context.Context) ([]entities.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) SelectAll(ctx context.Context) ([]entities.Product, error) {
	var productList []entities.Product

	productBuilder := r.db.WithContext(ctx).Unscoped().
		Limit(1000).Find(&productList)

	if err := productBuilder.Error; err != nil {
		return nil, err
	}

	return productList, nil
}
