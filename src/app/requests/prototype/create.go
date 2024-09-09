package requests

type PrototypeCreateRequest struct {
	ProductCode string `json:"productCode" validate:"required,min=3,email"`
}
