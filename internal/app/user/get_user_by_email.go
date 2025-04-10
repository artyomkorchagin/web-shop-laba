package user

import (
	"context"
	"errors"
	"socialsecurity/internal/types"
)

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {

	if email == "" {
		return nil, errors.New("email can't be empty")
	}

	user, err := s.repo.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
