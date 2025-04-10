package user

import (
	"context"
	"errors"
	"socialsecurity/internal/types"
)

func (s *Service) GetUser(ctx context.Context, email, password string) (*types.User, error) {

	if email == "" || password == "" {
		return nil, errors.New("email or password can't be empty")
	}

	user, err := s.repo.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}

	if !VerifyPassword(password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
