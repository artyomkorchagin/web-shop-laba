package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetAllProducts(ctx context.Context) ([]types.Product, error) {
	return s.repo.GetAllProducts(ctx)
}
