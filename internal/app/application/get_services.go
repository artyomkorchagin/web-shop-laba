package application

import (
	"context"
	"socialsecurity/internal/types"
)

func (s *Service) GetServices(ctx context.Context) ([]types.Service, error) {

	services, err := s.repo.GetServices(ctx)
	if err != nil {
		return nil, err
	}

	return services, nil
}
