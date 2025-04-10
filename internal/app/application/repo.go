package application

import (
	"context"
	"socialsecurity/internal/types"

	"github.com/google/uuid"
)

type Reader interface {
	ListAllApplications(ctx context.Context) ([]*types.Application, error)
	ListUsersApplications(ctx context.Context, user_id uuid.UUID) ([]*types.Application, error)
	GetServices(ctx context.Context) ([]types.Service, error)
	GetBenefits(ctx context.Context) ([]types.Benefit, error)
	GetApplicationByID(ctx context.Context, id uuid.UUID) (*types.Application, error)
}

type Writer interface {
	WorkApplication(ctx context.Context, input types.WorkApplicationRequest) error
	AddApplication(ctx context.Context, a *types.Application) error
	EditApplication(ctx context.Context, a *types.Application) error
}

type ReadWriter interface {
	Reader
	Writer
}
