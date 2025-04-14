package product

import "context"

func (s *Service) IncreaseViewCount(ctx context.Context, product_id int) error {
	return s.repo.IncreaseViewCount(ctx, product_id)
}
