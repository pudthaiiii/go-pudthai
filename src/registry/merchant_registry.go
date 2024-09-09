package registry

import (
	merchant "github.com/pudthaiiii/golang-cms/src/app/controller/merchants"
)

// apply service to controller
func (r *registry) NewMerchantsController() merchant.MerchantsController {
	return merchant.NewMerchantsController(r.PrototypeInteractor())
}
