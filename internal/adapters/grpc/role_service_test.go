package grpc

import (
	"context"
	"testing"

	"github.com/andrei-kozel/owly-proto/golang/roles"
	"github.com/andrei-kozel/owly-roles/internal/adapters/db"
	"github.com/andrei-kozel/owly-roles/internal/application/core/api"
	"github.com/andrei-kozel/owly-roles/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRoleService() (*db.MockRoleRepository, *RoleService) {
	mockRoleRepository := new(db.MockRoleRepository)
	application := api.NewApplication(mockRoleRepository)
	roleService := NewRoleService(application, 3000)
	return mockRoleRepository, roleService
}

func TestCreateRole(t *testing.T) {
	mockRepo, roleService := setupRoleService()

	roleRequest := &roles.CreateRoleRequest{
		Name:        "admin",
		Description: "admin role",
	}

	expectedRole := &domain.Role{
		Guid:        "12345",
		Name:        "admin",
		Description: "admin role",
	}

	mockRepo.On("AddRole", mock.Anything, mock.AnythingOfType("*domain.Role")).Return(expectedRole, nil)

	resp, err := roleService.CreateRole(context.Background(), roleRequest)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "12345", resp.Role.Id)
	assert.Equal(t, "admin", resp.Role.Name)
	assert.Equal(t, "admin role", resp.Role.Description)

	mockRepo.AssertExpectations(t)
}

func TestGetRoles(t *testing.T) {
	mockRepo, roleService := setupRoleService()

	expectedRoles := []*domain.Role{
		{Guid: "12345", Name: "admin", Description: "admin role"},
		{Guid: "67890", Name: "user", Description: "user role"},
	}

	mockRepo.On("GetRoles", mock.Anything).Return(expectedRoles, nil)

	resp, err := roleService.GetRoles(context.Background(), &roles.GetRolesRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp.Roles))

	for i, role := range expectedRoles {
		assert.Equal(t, role.Guid, resp.Roles[i].Id)
		assert.Equal(t, role.Name, resp.Roles[i].Name)
		assert.Equal(t, role.Description, resp.Roles[i].Description)
	}

	mockRepo.AssertExpectations(t)
}

func TestGetRole(t *testing.T) {
	mockRepo, roleService := setupRoleService()

	expectedRole := &domain.Role{
		Guid:        "12345",
		Name:        "admin",
		Description: "admin role",
	}

	mockRepo.On("GetRole", mock.Anything, mock.Anything).Return(expectedRole, nil)

	resp, err := roleService.GetRole(context.Background(), &roles.GetRoleRequest{Id: "12345"})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "12345", resp.Role.Id)
	assert.Equal(t, "admin", resp.Role.Name)
	assert.Equal(t, "admin role", resp.Role.Description)

	mockRepo.AssertExpectations(t)
}

func TestDeleteRole(t *testing.T) {
	mockRepo, roleService := setupRoleService()

	mockRepo.On("DeleteRole", mock.Anything, mock.Anything).Return(nil)

	resp, err := roleService.DeleteRole(context.Background(), &roles.DeleteRoleRequest{Id: "12345"})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "12345", resp.Role.Id)

	mockRepo.AssertExpectations(t)
}
