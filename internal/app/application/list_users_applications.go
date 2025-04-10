package application

import (
	"context"
	"socialsecurity/internal/types"

	"github.com/google/uuid"
)

func (s Service) ListUsersApplications(ctx context.Context, user_id uuid.UUID) ([]*types.Application, error) {

	apps, err := s.repo.ListUsersApplications(ctx, user_id)
	if err != nil {
		return []*types.Application{}, err
	}
	return apps, nil
}
