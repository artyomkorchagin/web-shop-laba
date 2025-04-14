package product

import "context"

func (s *Service) IncreasePurchaseCount(ctx context.Context, product_id int) error {
	return s.repo.IncreasePurchaseCount(ctx, product_id)
}
