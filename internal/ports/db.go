package ports

import (
	"context"

	"github.com/andrei-kozel/owly-roles/internal/application/core/domain"
)

type DBPort interface {
	GetRole(ctx context.Context, id string) (*domain.Role, error)
	AddRole(ctx context.Context, role *domain.Role) (*domain.Role, error)
	DeleteRole(ctx context.Context, id string) error
	GetRoles(ctx context.Context) ([]*domain.Role, error)
}
