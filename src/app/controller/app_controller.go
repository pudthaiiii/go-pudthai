package controller

import (
	merchants "workshop/src/app/controller/merchants"
	prototype "workshop/src/app/controller/prototype"
)

type AppController struct {
	prototype.PrototypeController
	merchants.MerchantsController
}
