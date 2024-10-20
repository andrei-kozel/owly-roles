package api

import (
	"context"

	"github.com/andrei-kozel/owly-roles/internal/application/core/domain"
	"github.com/andrei-kozel/owly-roles/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a *Application) AddRole(context context.Context, role *domain.Role) (*domain.Role, error) {
	return a.db.AddRole(context, role)
}

func (a *Application) DeleteRole(context context.Context, id string) error {
	return a.db.DeleteRole(context, id)
}

func (a *Application) GetRoles(context context.Context) ([]*domain.Role, error) {
	return a.db.GetRoles(context)
}

func (a *Application) GetRole(context context.Context, id string) (*domain.Role, error) {
	return a.db.GetRole(context, id)
}
