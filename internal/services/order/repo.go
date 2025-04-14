package order

import (
	"context"
)

type Reader interface {
}

type Writer interface {
	AddOrder(ctx context.Context, email, address string) error
}

type ReadWriter interface {
	Reader
	Writer
}
