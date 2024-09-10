package registry

import "github.com/pudthaiiii/golang-cms/src/app/repository"

// apply entity(database) to repository
func (r *registry) NewUsageItemRepository() repository.UsageItemRepository {
	return repository.NewUsageItemRepository(r.db)
}

func (r *registry) NewProductRepository() repository.ProductRepository {
	return repository.NewProductRepository(r.db)
}
