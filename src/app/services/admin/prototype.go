package services

import (
	"context"

	dtos "github.com/pudthaiiii/golang-cms/src/app/controller/admin/dtos/response/prototype"
	repo "github.com/pudthaiiii/golang-cms/src/app/repository"
)

type prototypeService struct {
	usageItemRepo repo.UsageItemRepository
	productRepo   repo.ProductRepository
}

type PrototypeService interface {
	Paginate(ctx context.Context) ([]dtos.UsageItemResponse, error)
}

func NewPrototypeService(
	usageItemRepo repo.UsageItemRepository,
	productRepo repo.ProductRepository) PrototypeService {
	return &prototypeService{
		usageItemRepo,
		productRepo,
	}
}

func (s *prototypeService) Paginate(ctx context.Context) ([]dtos.UsageItemResponse, error) {
	var response []dtos.UsageItemResponse

	data, err := s.usageItemRepo.SelectAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range data {
		response = append(response, dtos.UsageItemResponse{
			ID:          item.ID,
			OrgCode:     item.OrgCode,
			ProductCode: item.ProductCode,
		})
	}

	return response, nil
}
