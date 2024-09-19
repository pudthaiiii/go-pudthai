package registry

import (
	"go-ibooking/internal/model"
	a "go-ibooking/src/app/services/admin"

	"gorm.io/gorm"
)

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

func (r *registry) RegisterAuthService() a.AuthService {
	return a.NewAuthService(
		r.db.Model(&model.User{}).Session(&gorm.Session{NewDB: true}),
		r.RegisterUsersService(),
	)
}
