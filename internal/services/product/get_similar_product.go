package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetSimilarProduct(ctx context.Context, userInput string) (types.Product, error) {
	return s.repo.GetSimilarProduct(ctx, userInput)
}
