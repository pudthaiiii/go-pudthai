package router

import (
	cc "go-ibooking/internal/adapter/console/controllers"
	t "go-ibooking/internal/model/technical"
)

func addDatabaseRoute(c cc.DatabaseController) t.Routes {
	return t.Routes{
		t.Route{
			Name: "AutoMigrate", Method: "GET", Path: "/auto-migrate", Action: string(t.COOKIE), Subject: string(t.CONSOLE), HandlerFunc: c.AutoMigrate,
		},
	}
}
