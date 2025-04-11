package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) AddProduct(ctx context.Context, req *types.CreateProductRequest) error {
	product := types.NewProduct(req)

	if err := s.repo.AddProduct(ctx, product); err != nil {
		return err
	}

	return nil
}
