package application

import (
	"context"
	"socialsecurity/internal/types"
)

func (s *Service) GetBenefits(ctx context.Context) ([]types.Benefit, error) {

	benefits, err := s.repo.GetBenefits(ctx)
	if err != nil {
		return nil, err
	}

	return benefits, nil
}
