package product

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

type Reader interface {
	GetCartByUserEmail(ctx context.Context, email string) ([]types.Product, error)
	GetAllProducts(ctx context.Context) ([]types.Product, error)
	GetProductById(ctx context.Context, id int) (types.Product, error)
}

type Writer interface {
	IncreasePurchaseCount(ctx context.Context, product_id int) error
	IncreaseViewCount(ctx context.Context, product_id int) error
	AddToCart(ctx context.Context, email string, productID, amount int) error
	AddProduct(ctx context.Context, p *types.Product) error
}

type ReadWriter interface {
	Reader
	Writer
}
