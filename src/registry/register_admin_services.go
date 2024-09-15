package registry

import (
	"go-ibooking/src/app/model"
	a "go-ibooking/src/app/services/admin"

	"gorm.io/gorm"
)

// NewPrototypeService
func (r *registry) RegisterPrototypeService() a.PrototypeService {
	return a.NewPrototypeService(
		r.db.Model(&model.Merchant{}),
		r.db.Model(&model.Role{}),
	)
}

func (r *registry) RegisterRoleService() a.RolesService {
	return a.NewRoleService(
		r.db.Model(&model.Role{}).Session(&gorm.Session{NewDB: true}),
	)
}

func (r *registry) RegisterUsersService() a.UsersService {
	return a.NewUsersService(
		r.db.Model(&model.User{}).Session(&gorm.Session{NewDB: true}),
		r.s3,
	)
}
