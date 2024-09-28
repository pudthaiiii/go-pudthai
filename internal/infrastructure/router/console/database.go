package router

import (
	cc "go-ibooking/internal/adapter/console/controllers"
	"go-ibooking/internal/enum/permission"
	t "go-ibooking/internal/model/technical"
)

func addDatabaseRoute(c cc.DatabaseController) t.Routes {
	return t.Routes{
		t.Route{
			Name: "AutoMigrate", Method: "GET", Path: "/auto-migrate", Action: string(permission.COOKIE), Subject: string(permission.CONSOLE), HandlerFunc: c.AutoMigrate,
		},
	}
}
