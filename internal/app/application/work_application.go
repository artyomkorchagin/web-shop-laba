package application

import (
	"context"
	"socialsecurity/internal/types"
)

func (s Service) WorkApplication(ctx context.Context, input types.WorkApplicationRequest) error {

	err := s.repo.WorkApplication(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
