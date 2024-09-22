package repository

import (
	"context"
	"go-ibooking/internal/entities"
	"go-ibooking/internal/model/dtos"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db: db}
}

type UsersRepository interface {
	CreateAdminUser(ctx context.Context, dto dtos.CreateUser, fileName string, password string) (entities.User, error)
	FindUserByEmail(ctx context.Context, email string, userType string) (entities.User, error)
}

func (r *usersRepository) CreateAdminUser(ctx context.Context, dto dtos.CreateUser, fileName string, password string) (entities.User, error) {
	var user = entities.User{
		Email:        dto.Email,
		Password:     string(password),
		RoleID:       dto.RoleID,
		ProfileImage: fileName,
		IsActive:     dto.IsActive,
		IsAllBU:      dto.IsAllBU,
		FullName:     dto.FullName,
		Mobile:       dto.Mobile,
		Company:      dto.Company,
		Type:         dto.Type,
		MerchantID:   1,
	}

	query := r.db.WithContext(ctx).Create(&user)
	if query.Error != nil {
		return user, query.Error
	}

	return user, nil
}

func (r *usersRepository) FindUserByEmail(ctx context.Context, email string, userType string) (entities.User, error) {
	var user entities.User

	query := r.db.WithContext(ctx).Where("email = ?", email)

	if userType != "" {
		query = query.Where("type = ?", userType)
	}

	err := query.First(&user).Error
	return user, err
}
