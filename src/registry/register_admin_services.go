package registry

import (
	"go-ibooking/src/app/model"
	a "go-ibooking/src/app/services/admin"
)

// NewPrototypeService
func (r *registry) RegisterPrototypeService() a.PrototypeService {
	return a.NewPrototypeService(
		r.db.Model(&model.Merchant{}),
		r.db.Model(&model.Role{}),
	)
}

func (r *registry) RegisterRoleService() a.RoleService {
	return a.NewRoleService(
		r.db.Model(&model.Role{}),
	)
}
