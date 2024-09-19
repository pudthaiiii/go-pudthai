package router

import (
	console "go-ibooking/internal/adapter/controllers/v1/console"
	"go-ibooking/src/types"
)

func addFeaturesRoute(c console.FeaturesController) types.Routes {
	return types.Routes{
		types.Route{
			Name:        "AutoMigrate",
			Method:      "GET",
			Path:        "/auto-migrate",
			Action:      "",
			Subject:     "",
			HandlerFunc: c.AutoMigrate,
		},
	}
}
