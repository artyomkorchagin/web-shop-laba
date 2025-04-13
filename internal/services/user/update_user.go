package user

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) UpdateUser(ctx context.Context, u *types.User) error {

	return s.repo.UpdateUser(ctx, u)
}
