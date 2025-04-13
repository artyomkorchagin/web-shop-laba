package order

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

type Reader interface {
}

type Writer interface {
	CreateOrder(ctx context.Context, order *types.Order) error
}

type ReadWriter interface {
	Reader
	Writer
}
