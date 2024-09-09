package controller

import (
	merchants "github.com/pudthaiiii/golang-cms/src/app/controller/merchants"
	prototype "github.com/pudthaiiii/golang-cms/src/app/controller/prototype"
)

type AppController struct {
	prototype.PrototypeController
	merchants.MerchantsController
}
