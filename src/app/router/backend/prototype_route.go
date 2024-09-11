package route

import (
	controller "github.com/pudthaiiii/go-ibooking/src/app/http/backend/controllers/prototype"
	technical "github.com/pudthaiiii/go-ibooking/src/types"
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
