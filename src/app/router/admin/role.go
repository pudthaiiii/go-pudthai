package route

import (
	controller "go-ibooking/src/app/http/admin/controllers/role"
	"go-ibooking/src/types"
)

func addRoleRoute(c controller.RoleController) types.Routes {
	return types.Routes{
		types.Route{
			Name:        "Paginate",
			Method:      "GET",
			Pattern:     "/roles",
			Operation:   "",
			Resource:    "",
			HandlerFunc: c.Paginate,
		},

		types.Route{
			Name:        "Create",
			Method:      "POST",
			Pattern:     "/role",
			Operation:   "",
			Resource:    "",
			HandlerFunc: c.Create,
		},
	}
}
