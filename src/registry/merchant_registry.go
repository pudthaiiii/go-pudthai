package registry

import (
	merchant "workshop/src/app/controller/merchants"
)

// apply service to controller
func (r *registry) NewMerchantsController() merchant.MerchantsController {
	return merchant.NewMerchantsController(r.PrototypeInteractor())
}
