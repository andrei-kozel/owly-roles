package db

import (
	"context"

	"github.com/andrei-kozel/owly-roles/internal/application/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) AddRole(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	args := m.Called(ctx, role)
	return args.Get(0).(*domain.Role), args.Error(1)
}

func (m *MockRoleRepository) GetRole(ctx context.Context, id string) (*domain.Role, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Role), args.Error(1)
}

func (m *MockRoleRepository) DeleteRole(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRoleRepository) GetRoles(ctx context.Context) ([]*domain.Role, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Role), args.Error(1)
}
