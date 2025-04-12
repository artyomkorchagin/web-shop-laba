package category

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

type Reader interface {
	GetAllCategories(ctx context.Context) ([]types.Category, error)
}

type Writer interface {
	AddCategory(ctx context.Context, ccr *types.CreateCategoryRequest) error
}

type ReadWriter interface {
	Reader
	Writer
}
