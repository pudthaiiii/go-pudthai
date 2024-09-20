package repository

import (
	"context"
	"go-ibooking/internal/entities"
	"go-ibooking/internal/model/technical"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db: db}
}

type UsersRepository interface {
	CreateAdminUser(ctx context.Context, dto technical.CreateAdminUser, fileName string) (technical.ResponseAdminUser, error)
}

func (r *usersRepository) CreateAdminUser(ctx context.Context, dto technical.CreateAdminUser, fileName string) (technical.ResponseAdminUser, error) {
	var (
		response technical.ResponseAdminUser
		user     entities.User
	)

	// hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	user = entities.User{
		Email:        dto.Email,
		Password:     string(hashedPassword),
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

	queryBuilder := r.db.Create(&user)
	if queryBuilder.Error != nil {
		return response, queryBuilder.Error
	}

	response = technical.ResponseAdminUser{
		ID: user.ID,
	}

	return response, nil
}
