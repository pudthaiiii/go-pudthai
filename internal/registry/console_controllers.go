package registry

import (
	console "go-ibooking/internal/adapter/v1/controllers/console"
)

// NewPrototypeController
func (r *registry) RegisterFeaturesController() console.FeaturesController {
	return console.NewFeaturesController(r.db)
}
