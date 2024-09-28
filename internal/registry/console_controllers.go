package registry

import (
	cc "go-ibooking/internal/adapter/console/controllers"
)

// NewPrototypeController
func (r *registry) NewConsoleDatabaseController() cc.DatabaseController {
	return cc.NewDatabaseController(r.db)
}
