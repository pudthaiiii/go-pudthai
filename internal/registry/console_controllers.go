package registry

import (
	console "go-ibooking/internal/adapter/controllers/v1/console"
)

// NewPrototypeController
func (r *registry) RegisterFeaturesController() console.FeaturesController {
	return console.NewFeaturesController(r.db)
}
