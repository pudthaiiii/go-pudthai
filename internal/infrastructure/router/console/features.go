package router

import (
	controller "go-ibooking/internal/adapter/v1/controllers/console"
	"go-ibooking/internal/model/technical"
)

func addFeaturesRoute(c controller.FeaturesController) technical.Routes {
	return technical.Routes{
		technical.Route{
			Name:        "AutoMigrate",
			Method:      "GET",
			Path:        "/auto-migrate",
			Action:      "",
			Subject:     "",
			HandlerFunc: c.AutoMigrate,
		},
	}
}
