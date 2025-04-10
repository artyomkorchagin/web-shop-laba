package user

import (
	"context"
	"socialsecurity/internal/types"
)

func (s *Service) GetAllUsers(ctx context.Context) ([]types.User, error) {

	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
