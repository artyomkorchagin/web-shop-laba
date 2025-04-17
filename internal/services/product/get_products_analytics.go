package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetProductsAnalytics(ctx context.Context) ([]types.Product, error) {
	return s.repo.GetProductsAnalytics(ctx)
}
