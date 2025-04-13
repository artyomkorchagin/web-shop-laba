package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetProductById(ctx context.Context, id int) (types.Product, error) {
	return s.repo.GetProductById(ctx, id)
}
