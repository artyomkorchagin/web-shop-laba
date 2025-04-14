package product

import (
	"context"
)

func (s *Service) AddToCart(ctx context.Context, email string, productID int, amount int) error {
	return s.repo.AddToCart(ctx, email, productID, amount)
}
