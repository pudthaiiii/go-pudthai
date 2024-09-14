package controllers

import (
	prototype "go-ibooking/src/app/http/admin/controllers/prototype"
	role "go-ibooking/src/app/http/admin/controllers/role"
)

type AdminController struct {
	prototype.PrototypeController
	role.RoleController
}
