package cart

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

type Reader interface {
	GetCartByUserID(ctx context.Context, userID uint) (*types.Cart, error)
}

type Writer interface {
	AddToCart(ctx context.Context, userID uint, productID uint) error
}

type ReadWriter interface {
	Reader
	Writer
}
