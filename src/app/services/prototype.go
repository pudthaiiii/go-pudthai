package services

import (
	"context"
	repo "workshop/src/app/repository"
	resource "workshop/src/app/resources/prototype"
)

type prototypeService struct {
	usageItemRepository repo.UsageItemRepository
	productRepository   repo.ProductRepository
}

type PrototypeInteractor interface {
	Paginate(ctx context.Context) ([]resource.UsageItemResponse, error)
}

func NewPrototypeInteractor(usageItemRepository repo.UsageItemRepository, productRepository repo.ProductRepository) PrototypeInteractor {
	return &prototypeService{usageItemRepository, productRepository}
}

func (p *prototypeService) Paginate(ctx context.Context) ([]resource.UsageItemResponse, error) {
	data, err := p.usageItemRepository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []resource.UsageItemResponse

	// Transform the data to include only selected fields
	for _, item := range data {
		response = append(response, resource.UsageItemResponse{
			ID:          item.ID,
			OrgCode:     item.OrgCode,
			ProductCode: item.ProductCode,
		})
	}

	return response, nil

}
