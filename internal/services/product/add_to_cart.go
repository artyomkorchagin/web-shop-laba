package product

import "context"

func (s *Service) AddToCart(ctx context.Context, email string, productID uint) error {
	return s.repo.AddToCart(ctx, email, productID)
}
