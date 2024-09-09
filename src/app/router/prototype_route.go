package route

import (
	prototype "github.com/pudthaiiii/golang-cms/src/app/controller/prototype"
	technical "github.com/pudthaiiii/golang-cms/src/types"
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
