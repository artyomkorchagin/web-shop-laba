package cart

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetCartByUserID(ctx context.Context, userid uint) (*types.Cart, error) {
	cart, err := s.repo.GetCartByUserID(ctx, userid)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
