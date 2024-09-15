package dtos

import "github.com/google/uuid"

type RolePaginateResponse struct {
	ID               uint      `json:"id"`
	Uuid             uuid.UUID `json:"uuid"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	IsActive         int       `json:"isActive"`
	IsCorporateAdmin int       `json:"isCorporateAdmin"`
	MerchantID       uint      `json:"merchantId"`
	CreatedAt        string    `json:"createdAt"`
	UpdatedAt        string    `json:"updatedAt"`
	Merchant         RoleMerchant
}

type RoleMerchant struct {
	ID   uint      `json:"id"`
	Name string    `json:"name"`
	Uuid uuid.UUID `json:"uuid"`
}
