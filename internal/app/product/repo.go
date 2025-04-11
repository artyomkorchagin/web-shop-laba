package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

type Reader interface {
	GetAllProducts(ctx context.Context) ([]types.Product, error)
}

type Writer interface {
	AddProduct(ctx context.Context, p *types.Product) error
}

type ReadWriter interface {
	Reader
	Writer
}
