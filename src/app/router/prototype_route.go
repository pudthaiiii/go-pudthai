package route

import (
	prototype "workshop/src/app/controller/prototype"
	technical "workshop/src/types"
)

func addPrototypeRoute(c prototype.PrototypeController) technical.Routes {
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
