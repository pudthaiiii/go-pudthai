package route

import (
	merchants "workshop/src/app/controller/merchants"
	technical "workshop/src/types"
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
