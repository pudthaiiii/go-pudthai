package registry

import (
	prototype "github.com/pudthaiiii/golang-cms/src/app/controller/prototype"
	"github.com/pudthaiiii/golang-cms/src/app/services"
)

// apply service to controller
func (r *registry) NewPrototypeController() prototype.PrototypeController {
	return prototype.NewPrototypeController(r.PrototypeInteractor())
}

// apply repository to service
func (r *registry) PrototypeInteractor() services.PrototypeInteractor {
	return services.NewPrototypeInteractor(r.NewUsageItemRepository(), r.NewProductRepository())
}
