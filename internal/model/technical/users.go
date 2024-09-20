package technical

type CreateAdminUser struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Mobile   string `json:"mobile" validate:"required,min=10"`
	Company  string `json:"company"`
	IsActive int    `json:"isActive" validate:"required,oneOrZero"`
	Type     string `json:"type" validate:"required"`
	RoleID   uint   `json:"roleId" validate:"required"`
	IsAllBU  int    `json:"isAllBu" validate:"required,oneOrZero"`
}

type ResponseAdminUser struct {
	ID uint `json:"id"`
}
