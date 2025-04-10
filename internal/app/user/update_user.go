package user

import (
	"artyomkorchagin/web-shop/internal/types"
	"context"
)

func (s *Service) UpdateUser(ctx context.Context, u *types.User) error {
	if err := s.repo.UpdateUser(ctx, u); err != nil {
		return err
	}

	return nil
}
