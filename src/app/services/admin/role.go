package services

import (
	"context"

	"go-ibooking/src/app/model"
	"go-ibooking/src/pkg/logger"

	"gorm.io/gorm"

	req "go-ibooking/src/app/http/admin/dtos/request/roles"
	res "go-ibooking/src/app/http/admin/dtos/response/roles"
)

type RoleService interface {
	Paginate(ctx context.Context) ([]res.RolePaginateResponse, error)
	Create(ctx context.Context, dtoReq req.RoleCreateRequest) (res.CreateRoleResponse, error)
}

type roleService struct {
	roleRepo *gorm.DB
}

func NewRoleService(roleRepo *gorm.DB) RoleService {
	return &roleService{
		roleRepo,
	}
}

func (s *roleService) Create(ctx context.Context, dtoReq req.RoleCreateRequest) (res.CreateRoleResponse, error) {
	response := res.CreateRoleResponse{}

	role := model.Role{
		Name:             dtoReq.Name,
		IsCorporateAdmin: dtoReq.IsCorporateAdmin,
		IsActive:         dtoReq.IsActive,
		MerchantID:       1,
	}

	queryBuilder1 := s.roleRepo.Create(&role)
	if queryBuilder1.Error != nil {
		logger.Log.Err(queryBuilder1.Error).Msg("Failed to create roles")
		return response, queryBuilder1.Error
	}

	response = res.CreateRoleResponse{
		ID: role.ID,
	}

	return response, nil
}

func (s *roleService) Paginate(ctx context.Context) ([]res.RolePaginateResponse, error) {
	roles := []model.Role{}
	response := []res.RolePaginateResponse{}

	queryBuilder := s.roleRepo.
		Find(&roles)

	if queryBuilder.Error != nil {
		logger.Log.Err(queryBuilder.Error).Msg("Failed to fetch roles")
		return response, queryBuilder.Error
	}

	for _, role := range roles {
		response = append(response, res.RolePaginateResponse{
			ID:               role.ID,
			Uuid:             role.Uuid,
			Name:             role.Name,
			IsActive:         role.IsActive,
			IsCorporateAdmin: role.IsCorporateAdmin,
			MerchantID:       role.MerchantID,
			CreatedAt:        role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        role.UpdatedAt.Format("2006-01-02 15:04:05"),
			Merchant: res.RoleMerchant{
				ID:   role.Merchant.ID,
				Name: role.Merchant.Name,
				Uuid: role.Merchant.Uuid,
			},
		})
	}

	return response, nil
}
