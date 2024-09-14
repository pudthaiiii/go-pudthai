package dtos

type RoleCreateRequest struct {
	Name             string `json:"name" validate:"required"`
	IsCorporateAdmin int    `json:"isCorporateAdmin" validate:"required,oneOrZero"`
	IsActive         int    `json:"isActive" validate:"required,oneOrZero"`
}
