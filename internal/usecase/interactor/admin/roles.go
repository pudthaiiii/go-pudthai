package services

import (
	"context"
	"fmt"

	"go-ibooking/internal/model"
	"go-ibooking/internal/model/scopes"
	"go-ibooking/src/utils"

	"gorm.io/gorm"

	throw "go-ibooking/src/app/exception"
	req "go-ibooking/src/app/http/admin/dtos/request"
	res "go-ibooking/src/app/http/admin/dtos/response"

	dtoReq "go-ibooking/src/app/http/admin/dtos/request/roles"
	dtoRes "go-ibooking/src/app/http/admin/dtos/response/roles"
)

type RolesService interface {
	Paginate(ctx context.Context, params req.PaginateRequest) (result []dtoRes.RolePaginateResponse, paginate res.Pagination, err error)
	Create(ctx context.Context, dto dtoReq.RoleCreateRequest) (dtoRes.CreateRoleResponse, error)
}

type rolesService struct {
	roleRepo *gorm.DB
}

func NewRoleService(roleRepo *gorm.DB) RolesService {
	return &rolesService{
		roleRepo,
	}
}

func (s *rolesService) Create(ctx context.Context, dto dtoReq.RoleCreateRequest) (dtoRes.CreateRoleResponse, error) {
	response := dtoRes.CreateRoleResponse{}

	role := model.Role{
		Name:             dto.Name,
		IsCorporateAdmin: dto.IsCorporateAdmin,
		IsActive:         dto.IsActive,
		MerchantID:       1,
	}

	queryBuilder1 := s.roleRepo.Create(&role)
	if queryBuilder1.Error != nil {
		return response, throw.Error(910101, queryBuilder1.Error)
	}

	response = dtoRes.CreateRoleResponse{
		ID: role.ID,
	}

	return response, nil
}

func (s *rolesService) Paginate(ctx context.Context, params req.PaginateRequest) (result []dtoRes.RolePaginateResponse, paginate res.Pagination, err error) {
	var totalRecord int64

	fmt.Println("params", params)
	roles := []model.Role{}
	keySearch := []string{"name", "description"}

	countBuilder := s.roleRepo.
		Model(&model.Role{}).
		Scopes(
			scopes.WithSearch(params.Search, keySearch),
			scopes.WithIsActive(params.Filters.IsActive),
		).
		Count(&totalRecord)

	if countBuilder.Error != nil {
		return result, paginate, throw.Error(910102, countBuilder.Error)
	}

	if totalRecord == 0 {
		return result, paginate, nil
	}

	queryBuilder := s.roleRepo.
		Scopes(
			scopes.WithSearchAndPaginate(params.Search, keySearch, params.Page, params.PerPage),
			scopes.WithIsActive(params.Filters.IsActive),
		).
		Find(&roles)

	if queryBuilder.Error != nil {
		return result, paginate, throw.Error(910102, queryBuilder.Error)
	}

	result = make([]dtoRes.RolePaginateResponse, len(roles))
	for i, role := range roles {
		result[i] = transformRoleToResponse(role)
	}

	paginate = utils.CalculationPaginate(params.Page, params.PerPage, totalRecord)

	return result, paginate, nil
}

func transformRoleToResponse(role model.Role) dtoRes.RolePaginateResponse {
	return dtoRes.RolePaginateResponse{
		Uuid:             role.Uuid,
		Name:             role.Name,
		Description:      role.Description,
		IsActive:         role.IsActive,
		IsCorporateAdmin: role.IsCorporateAdmin,
		MerchantID:       role.MerchantID,
		CreatedAt:        role.CreatedAt.String(),
		UpdatedAt:        role.UpdatedAt.String(),
	}
}
