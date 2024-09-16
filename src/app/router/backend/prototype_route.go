package route

import (
	controller "go-ibooking/src/app/http/backend/controllers/prototype"
	technical "go-ibooking/src/types"
)

func addPrototypeRoute(c controller.PrototypeController) technical.Routes {
	return technical.Routes{
		technical.Route{
			Name:        "Paginate",
			Method:      "GET",
			Path:        "/prototype",
			Action:      "",
			Subject:     "",
			HandlerFunc: c.Paginate,
		},

		technical.Route{
			Name:        "Paginate",
			Method:      "POST",
			Path:        "/prototype",
			Action:      "",
			Subject:     "",
			HandlerFunc: c.Create,
		},
	}
}
