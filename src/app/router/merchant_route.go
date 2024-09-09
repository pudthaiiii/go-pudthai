package route

import (
	merchants "github.com/pudthaiiii/golang-cms/src/app/controller/merchants"
	technical "github.com/pudthaiiii/golang-cms/src/types"
)

func addMerchantRoute(c merchants.MerchantsController) technical.Routes {
	return technical.Routes{
		technical.Route{
			Name:        "GetAll",
			Method:      "GET",
			Pattern:     "/merchants",
			Operation:   "",
			Resource:    "",
			HandlerFunc: c.GetAll,
		},
	}
}
