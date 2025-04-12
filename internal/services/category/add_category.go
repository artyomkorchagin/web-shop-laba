package category

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) AddCategory(ctx context.Context, ccr *types.CreateCategoryRequest) error {
	return s.repo.AddCategory(ctx, ccr)
}
