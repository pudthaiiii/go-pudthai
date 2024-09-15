package dtos

type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}
