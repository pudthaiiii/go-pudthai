package business

import "github.com/google/uuid"

type UserInfo struct {
	ID         uint
	Uuid       uuid.UUID
	MerchantID uint
	FullName   string
	Email      string
	Mobile     string
	Company    string
	IsAllBU    int
	RoleID     uint
	Type       string
}
