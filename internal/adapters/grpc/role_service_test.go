package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/andrei-kozel/owly-common/domain"
	"github.com/andrei-kozel/owly-proto/golang/roles"
	"github.com/andrei-kozel/owly-roles/internal/adapters/db"
	"github.com/andrei-kozel/owly-roles/internal/application/core/api"
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

func TestGetRole_Error(t *testing.T) {
	mockRepo, roleService := setupRoleService()

	// Simulate an error returned by the GetRole method in the mock repository
	mockRepo.On("GetRole", mock.Anything, mock.Anything).Return((*domain.Role)(nil), errors.New("role not found"))

	// Call the GetRole method
	resp, err := roleService.GetRole(context.Background(), &roles.GetRoleRequest{Id: "nonexistent-id"})

	// Assertions
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "role not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestDeleteRole_Error(t *testing.T) {
	mockRepo, roleService := setupRoleService()

	// Simulate an error returned by the DeleteRole method in the mock repository
	mockRepo.On("DeleteRole", mock.Anything, mock.Anything).Return((*domain.Role)(nil), errors.New("role not found"))

	// Call the DeleteRole method
	resp, err := roleService.DeleteRole(context.Background(), &roles.DeleteRoleRequest{Id: "nonexistent-id"})

	// Assertions
	assert.Nil(t, resp)                            // Response should be nil if there's an error
	assert.Error(t, err)                           // There should be an error
	assert.Equal(t, "role not found", err.Error()) // Error message should match the expected error
	mockRepo.AssertExpectations(t)                 // Ensure the mock expectations were met
}

func TestGetRoles_Error(t *testing.T) {
	mockRepo, roleService := setupRoleService()

	// Simulate an error returned by the GetRoles method in the mock repository
	mockRepo.On("GetRoles", mock.Anything).Return(([]*domain.Role)(nil), errors.New("error fetching roles"))

	// Call the GetRoles method
	resp, err := roleService.GetRoles(context.Background(), &roles.GetRolesRequest{})

	// Assertions
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "error fetching roles", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestCreateRole_Error(t *testing.T) {
	mockRepo, roleService := setupRoleService()

	roleRequest := &roles.CreateRoleRequest{
		Name:        "admin",
		Description: "admin role",
	}

	// Simulate an error returned by the AddRole method in the mock repository
	mockRepo.On("AddRole", mock.Anything, mock.AnythingOfType("*domain.Role")).Return((*domain.Role)(nil), errors.New("error creating role"))

	// Call the CreateRole method
	resp, err := roleService.CreateRole(context.Background(), roleRequest)

	// Assertions
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "error creating role", err.Error())
	mockRepo.AssertExpectations(t)
}
