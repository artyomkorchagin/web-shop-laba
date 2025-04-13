package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetCartByUserEmail(ctx context.Context, email string) ([]types.Product, error) {
	return s.repo.GetCartByUserEmail(ctx, email)
}
