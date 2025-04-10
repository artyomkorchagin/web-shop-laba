package user

import (
	"context"
	"socialsecurity/internal/types"
)

func (s *Service) UpdateUser(ctx context.Context, u *types.User) error {
	if err := s.repo.UpdateUser(ctx, u); err != nil {
		return err
	}

	return nil
}
