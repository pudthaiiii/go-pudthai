package services

import (
	"context"
	"fmt"

	"go-ibooking/src/app/model"
	"go-ibooking/src/pkg/logger"

	"gorm.io/gorm"

	req "go-ibooking/src/app/http/admin/dtos/request/roles"
	res "go-ibooking/src/app/http/admin/dtos/response/roles"
)

type RoleService interface {
	Paginate(ctx context.Context) ([]model.Role, error)
	Create(ctx context.Context, req req.RoleCreateRequest) (res.CreateRoleResponse, error)
}

type roleService struct {
	roleRepo *gorm.DB
}

func NewRoleService(roleRepo *gorm.DB) RoleService {
	return &roleService{
		roleRepo,
	}
}

func (s *roleService) Create(ctx context.Context, req req.RoleCreateRequest) (res.CreateRoleResponse, error) {
	response := res.CreateRoleResponse{}

	role := model.Role{
		Name:             req.Name,
		IsCorporateAdmin: req.IsCorporateAdmin,
		IsActive:         req.IsActive,
		MerchantID:       1,
	}

	result := s.roleRepo.Create(&role)
	if result.Error != nil {
		logger.Log.Err(result.Error).Msg("Failed to create roles")
		return response, result.Error
	}

	response = res.CreateRoleResponse{
		ID: role.ID,
	}

	return response, nil
}

func (s *roleService) Paginate(ctx context.Context) ([]model.Role, error) {
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
