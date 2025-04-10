package user

import (
	"context"
	"socialsecurity/internal/types"
)

func (s *Service) AddUser(ctx context.Context, req types.CreateUserRequest) error {
	hashpass, err := HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashpass
	user := types.NewUser(req)

	if err := s.repo.AddUser(ctx, user); err != nil {
		return err
	}

	return nil
}
