package ports

import (
	"context"

	"github.com/andrei-kozel/owly-common/domain"
)

type APIport interface {
	AddRole(ctx context.Context, role *domain.Role) (*domain.Role, error)
	DeleteRole(ctx context.Context, id string) error
	GetRoles(ctx context.Context) ([]*domain.Role, error)
	GetRole(ctx context.Context, id string) (*domain.Role, error)
}
