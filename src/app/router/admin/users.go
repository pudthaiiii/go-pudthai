package route

import (
	controller "go-ibooking/src/app/http/admin/controllers/users"
	"go-ibooking/src/types"
)

func addUsersRoute(c controller.UsersController) types.Routes {
	return types.Routes{
		types.Route{
			Name:        "Create",
			Method:      "POST",
			Path:        "/users",
			Action:      "",
			Subject:     "",
			HandlerFunc: c.Create,
		},
	}
}
