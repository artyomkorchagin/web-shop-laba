package user

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) GetAllUsers(ctx context.Context) ([]types.User, error) {

	return s.repo.GetAllUsers(ctx)
}
