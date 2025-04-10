package application

import (
	"context"
	"socialsecurity/internal/types"
)

func (s Service) ListAllApplications(ctx context.Context) ([]*types.Application, error) {

	apps, err := s.repo.ListAllApplications(ctx)
	if err != nil {
		return []*types.Application{}, err
	}
	return apps, nil
}
