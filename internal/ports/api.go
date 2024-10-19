package ports

import "context"

type APIport interface {
	AddRole(ctx context.Context, role string) error
	RemoveRole(ctx context.Context, role string) error
	GetRoles(ctx context.Context) ([]string, error)
	GetRole(ctx context.Context, id string) (string, error)
}
