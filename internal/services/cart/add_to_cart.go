package cart

import "context"

func (s *Service) AddToCart(ctx context.Context, userID uint, productID uint) error {
	return s.repo.AddToCart(ctx, userID, productID)
}
