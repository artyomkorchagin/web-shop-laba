package order

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

type Reader interface {
	GetOrderHistory(ctx context.Context, email string) ([]types.Order, error)
}

type Writer interface {
	AddOrder(ctx context.Context, email, address string) error
}

type ReadWriter interface {
	Reader
	Writer
}
