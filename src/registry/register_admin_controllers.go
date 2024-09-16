package registry

import (
	auth "go-ibooking/src/app/http/admin/controllers/auth"
	prototype "go-ibooking/src/app/http/admin/controllers/prototype"
	role "go-ibooking/src/app/http/admin/controllers/role"
	users "go-ibooking/src/app/http/admin/controllers/users"
)

// NewPrototypeController
func (r *registry) RegisterPrototypeController() prototype.PrototypeController {
	return prototype.NewPrototypeController(r.RegisterPrototypeService())
}

// RoleController
func (r *registry) RegisterRoleController() role.RoleController {
	return role.NewRoleController(r.RegisterRoleService())
}

// UsersController
func (r *registry) RegisterUsersController() users.UsersController {
	return users.NewUsersController(r.RegisterUsersService())
}

// AuthController
func (r *registry) RegisterAuthController() auth.AuthController {
	return auth.NewAuthController(r.RegisterAuthService())
}
