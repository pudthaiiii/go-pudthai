package repository

import (
	"context"
	"go-pudthai/internal/entities"
	t "go-pudthai/internal/model/technical"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db: db}
}

type UsersRepository interface {
	CreateAdminUser(ctx context.Context, save entities.User) (entities.User, error)
	FindUserByEmail(ctx context.Context, email string, userType string) (entities.User, error)
}

func (r *usersRepository) CreateAdminUser(ctx context.Context, user entities.User) (entities.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *usersRepository) FindUserByEmail(ctx context.Context, email string, userType string) (entities.User, error) {
	var user entities.User

	query := r.db.WithContext(ctx).Where("LOWER(email) = LOWER(?)", email)

	if userType != "" {
		query = query.Where("LOWER(type) = LOWER(?)", userType)
	}

	merchantID, ok := ctx.Value(t.MerchantID).(string)
	if ok {
		query = query.Where("merchant_id = ?", merchantID)
	}

	err := query.Preload("Role").First(&user).Error
	return user, err
}
