package category

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetAllCategories(ctx context.Context) ([]types.Category, error) {
	return s.repo.GetAllCategories(ctx)
}
