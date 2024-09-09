package registry

import (
	prototype "workshop/src/app/controller/prototype"
	"workshop/src/app/services"
)

// apply service to controller
func (r *registry) NewPrototypeController() prototype.PrototypeController {
	return prototype.NewPrototypeController(r.PrototypeInteractor())
}

// apply repository to service
func (r *registry) PrototypeInteractor() services.PrototypeInteractor {
	return services.NewPrototypeInteractor(r.NewUsageItemRepository(), r.NewProductRepository())
}
