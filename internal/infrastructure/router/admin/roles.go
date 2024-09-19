package router

import (
	controller "go-ibooking/src/app/http/admin/controllers/role"
	"go-ibooking/src/types"
)

func addRolesRoute(c controller.RoleController) types.Routes {
	return types.Routes{
		types.Route{
			Name:        "Paginate",
			Method:      "GET",
			Path:        "/roles",
			Action:      "",
			Subject:     "",
			HandlerFunc: c.Paginate,
		},

		types.Route{
			Name:        "Create",
			Method:      "POST",
			Path:        "/role",
			Action:      "",
			Subject:     "",
			HandlerFunc: c.Create,
		},
	}
}
