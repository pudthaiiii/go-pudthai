package route

import (
	controller "go-ibooking/src/app/http/admin/controllers/prototype"
	technical "go-ibooking/src/types"
)

func addPrototypeRoute(c controller.PrototypeController) technical.Routes {
	return technical.Routes{
		technical.Route{
			Name:        "Paginate",
			Method:      "GET",
			Pattern:     "/prototype",
			Operation:   "",
			Resource:    "",
			HandlerFunc: c.Paginate,
		},

		technical.Route{
			Name:        "Paginate",
			Method:      "POST",
			Pattern:     "/prototype",
			Operation:   "",
			Resource:    "",
			HandlerFunc: c.Create,
		},
	}
}
