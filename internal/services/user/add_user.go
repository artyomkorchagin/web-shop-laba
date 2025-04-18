package user

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) AddUser(ctx context.Context, req types.CreateUserRequest) error {
	hashpass, err := HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashpass
	user := types.NewUser(req)

	return s.repo.AddUser(ctx, user)
}
