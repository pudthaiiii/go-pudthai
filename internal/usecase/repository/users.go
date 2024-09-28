package repository

import (
	"context"
	"fmt"
	"go-ibooking/internal/entities"
	"go-ibooking/internal/enum"
	"go-ibooking/internal/model/dtos"
	"strconv"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db: db}
}

type UsersRepository interface {
	CreateAdminUser(ctx context.Context, dto dtos.CreateUser, fileName string, password string, merchantID uint, userType string) (entities.User, error)
	FindUserByEmail(ctx context.Context, email string, userType string) (entities.User, error)
}

func (r *usersRepository) CreateAdminUser(ctx context.Context, dto dtos.CreateUser, fileName string, password string, merchantID uint, userType string) (entities.User, error) {
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
		Type:         userType,
	}

	if userType != string(enum.ADMIN) {
		merchantID, ok := ctx.Value("MerchantID").(string)
		if ok {
			merchantIDUint, err := strconv.ParseUint(merchantID, 10, 32)
			if err != nil {
				return user, err
			}

			user.MerchantID = uint(merchantIDUint)
		}
	}

	query := r.db.WithContext(ctx).Create(&user)
	if query.Error != nil {
		return user, query.Error
	}

	return user, nil
}

func (r *usersRepository) FindUserByEmail(ctx context.Context, email string, userType string) (entities.User, error) {
	var user entities.User

	Test := ctx.Value("TestMM")
	Test1 := ctx.Value("MerchantID")
	fmt.Println("Test", Test, Test1)

	fmt.Println("FInd")
	query := r.db.WithContext(ctx).Where("LOWER(email) = LOWER(?)", email)

	if userType != "" {
		query = query.Where("UPPER(type) = UPPER(?)", userType)
	}

	merchantID, ok := ctx.Value("MerchantID").(string)
	if ok {
		query = query.Where("merchant_id = ?", merchantID)
	}

	err := query.Preload("Role").First(&user).Error
	return user, err
}
