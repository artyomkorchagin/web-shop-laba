package cart

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

type Reader interface {
	GetAllCategories(ctx context.Context) ([]types.Category, error)
}

type Writer interface {
	AddCategory(ctx context.Context, name string) error
}

type ReadWriter interface {
	Reader
	Writer
}
