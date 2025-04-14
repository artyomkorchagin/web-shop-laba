package order

import "context"

func (s *Service) AddOrder(ctx context.Context, email, address string) error {
	return s.repo.AddOrder(ctx, email, address)
}
