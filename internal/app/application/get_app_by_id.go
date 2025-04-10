package application

import (
	"context"

	"socialsecurity/internal/types"

	"github.com/google/uuid"
)

func (s *Service) GetApplicationByID(ctx context.Context, id uuid.UUID) (*types.Application, error) {

	app, err := s.repo.GetApplicationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return app, nil
}
