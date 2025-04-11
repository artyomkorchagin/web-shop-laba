package user

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

type Reader interface {
	GetUser(ctx context.Context, email string) (*types.User, error)
	GetAllUsers(ctx context.Context) ([]types.User, error)
}

type Writer interface {
	AddUser(ctx context.Context, a *types.User) error
	UpdateUser(ctx context.Context, u *types.User) error
}

type ReadWriter interface {
	Reader
	Writer
}
