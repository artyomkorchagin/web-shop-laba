package user

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
	"errors"
)

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {

	if email == "" {
		return nil, errors.New("email can't be empty")
	}

	return s.repo.GetUser(ctx, email)
}
