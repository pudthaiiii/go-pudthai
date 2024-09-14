package services

import (
	"context"
	"fmt"

	"go-ibooking/src/app/model"
	"go-ibooking/src/pkg/logger"

	"gorm.io/gorm"
)

type prototypeService struct {
	merchantRepo *gorm.DB
	roleRepo     *gorm.DB
}

func NewPrototypeService(merchantRepo *gorm.DB, roleRepo *gorm.DB) PrototypeService {
	return &prototypeService{
		merchantRepo,
		roleRepo,
	}
}

type PrototypeService interface {
	Paginate(ctx context.Context) ([]model.Role, error)
	Create(ctx context.Context) (model.Role, error)
	// Paginate(ctx context.Context) ([]model.Merchant, error)
}

func (s *prototypeService) Create(ctx context.Context) (model.Role, error) {
	role := model.Role{
		Name:       "TEST",
		MerchantID: 1,
	}

	result := s.roleRepo.Create(&role)
	if result.Error != nil {
		logger.Log.Err(result.Error).Msg("Failed to create roles")
	}

	return role, nil
}

func (s *prototypeService) Paginate(ctx context.Context) ([]model.Role, error) {
	var response []model.Role

	var totalCount int64 = 0

	// Get total count
	err := s.roleRepo.Model(&model.Role{}).Count(&totalCount).Error
	if err != nil {
		return nil, err
	}

	// Get paginated results
	err = s.roleRepo.
		Preload("Merchant", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, name")
		}).
		Limit(10).
		Offset(0).
		Find(&response).Error

	if err != nil {
		return nil, err
	}

	fmt.Println("Total count: ", totalCount)
	return response, nil
}

// func (s *prototypeService) Paginate(ctx context.Context) ([]model.Merchant, error) {
// 	var response []model.Merchant

// 	err := s.merchantRepo.
// 		Preload("Roles").
// 		Find(&response).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }
