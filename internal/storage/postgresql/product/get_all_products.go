package psqlProduct

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (r *Repository) GetAllProducts(ctx context.Context) ([]types.Product, error) {
	return []types.Product{}, nil
}
