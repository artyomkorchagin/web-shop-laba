package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) AddProduct(ctx context.Context, req *types.CreateProductRequest) error {
	product := types.NewProduct(req)
	return s.repo.AddProduct(ctx, product)
}
