package user

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetAllUsers(ctx context.Context) ([]types.User, error) {

	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
