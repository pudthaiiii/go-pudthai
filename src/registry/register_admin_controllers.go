package registry

import (
	prototype "go-ibooking/src/app/http/admin/controllers/prototype"
	role "go-ibooking/src/app/http/admin/controllers/role"
)

// NewPrototypeController
func (r *registry) RegisterPrototypeController() prototype.PrototypeController {
	return prototype.NewPrototypeController(r.RegisterPrototypeService())
}

// RoleController
func (r *registry) RegisterRoleController() role.RoleController {
	return role.NewRoleController(r.RegisterRoleService())
}
