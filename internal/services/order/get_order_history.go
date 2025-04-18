package order

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetOrderHistory(ctx context.Context, email string) ([]types.Order, error) {
	orders, err := s.repo.GetOrderHistory(ctx, email)
	if err != nil {
		return nil, err
	}
	for i := range orders {
		sum := 0
		for _, p := range orders[i].Products {
			sum += int(p.Sum)
		}
		orders[i].TotalSum = sum
	}

	return orders, nil
}
