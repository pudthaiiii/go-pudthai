package route

import (
	controller "github.com/pudthaiiii/golang-cms/src/app/controller/admin/prototype"
	technical "github.com/pudthaiiii/golang-cms/src/types"
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
