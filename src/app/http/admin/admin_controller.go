package controllers

import (
	auth "go-ibooking/src/app/http/admin/controllers/auth"
	prototype "go-ibooking/src/app/http/admin/controllers/prototype"
	role "go-ibooking/src/app/http/admin/controllers/role"
	users "go-ibooking/src/app/http/admin/controllers/users"
)

type AdminController struct {
	prototype.PrototypeController
	role.RoleController
	users.UsersController
	auth.AuthController
}
