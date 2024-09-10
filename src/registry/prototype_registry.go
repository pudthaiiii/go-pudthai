package registry

import (
	adminController "github.com/pudthaiiii/golang-cms/src/app/controller/admin/prototype"
	adminService "github.com/pudthaiiii/golang-cms/src/app/services/admin"
)

// apply service to controller
func (r *registry) NewPrototypeController() adminController.PrototypeController {
	return adminController.NewPrototypeController(r.AdminService())
}

// apply repository to service
func (r *registry) AdminService() adminService.PrototypeService {
	return adminService.NewPrototypeService(r.NewUsageItemRepository(), r.NewProductRepository())
}
