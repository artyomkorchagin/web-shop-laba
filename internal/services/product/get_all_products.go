package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetAllProducts(ctx context.Context) ([]types.Product, error) {
	products, err := s.repo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
