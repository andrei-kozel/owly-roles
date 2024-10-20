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

func TestCreateRole(t *testing.T) {
	mockRoleRepository := new(db.MockRoleRepository)
	application := api.NewApplication(mockRoleRepository)
	roleService := NewRoleService(application, 3000)

	role := &roles.CreateRoleRequest{
		Name:        "admin",
		Description: "admin role",
	}

	// Set up expectations
	mockRoleRepository.On("AddRole", mock.Anything, mock.AnythingOfType("*domain.Role")).Return(&domain.Role{
		Guid:        "12345",
		Name:        "admin",
		Description: "admin role",
	}, nil)

	// Call the method
	resp, err := roleService.CreateRole(context.Background(), role)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "12345", resp.Role.Id)
	assert.Equal(t, "admin", resp.Role.Name)
	assert.Equal(t, "admin role", resp.Role.Description)

	// Ensure that the expectations were met
	mockRoleRepository.AssertExpectations(t)
}

func TestGetRoles(t *testing.T) {
	mockRoleRepository := new(db.MockRoleRepository)
	application := api.NewApplication(mockRoleRepository)
	roleService := NewRoleService(application, 3000)

	// Set up expectations
	mockRoleRepository.On("GetRoles", mock.Anything).Return([]*domain.Role{
		{
			Guid:        "12345",
			Name:        "admin",
			Description: "admin role",
		},
		{
			Guid:        "67890",
			Name:        "user",
			Description: "user role",
		},
	}, nil)

	// Call the method
	resp, err := roleService.GetRoles(context.Background(), &roles.GetRolesRequest{})
	roles := resp.Roles

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(roles))
	assert.Equal(t, "12345", roles[0].Id)
	assert.Equal(t, "admin", roles[0].Name)
	assert.Equal(t, "admin role", roles[0].Description)
	assert.Equal(t, "67890", roles[1].Id)
	assert.Equal(t, "user", roles[1].Name)
	assert.Equal(t, "user role", roles[1].Description)

	// Ensure that the expectations were met
	mockRoleRepository.AssertExpectations(t)
}

func TestGetRole(t *testing.T) {
	mockRoleRepository := new(db.MockRoleRepository)
	application := api.NewApplication(mockRoleRepository)
	roleService := NewRoleService(application, 3000)

	// Set up expectations
	mockRoleRepository.On("GetRole", mock.Anything, mock.Anything).Return(&domain.Role{
		Guid:        "12345",
		Name:        "admin",
		Description: "admin role",
	}, nil)

	// Call the method
	resp, err := roleService.GetRole(context.Background(), &roles.GetRoleRequest{Id: "12345"})

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "12345", resp.Role.Id)
	assert.Equal(t, "admin", resp.Role.Name)
	assert.Equal(t, "admin role", resp.Role.Description)

	// Ensure that the expectations were met
	mockRoleRepository.AssertExpectations(t)
}

func TestDeleteRole(t *testing.T) {
	mockRoleRepository := new(db.MockRoleRepository)
	application := api.NewApplication(mockRoleRepository)
	roleService := NewRoleService(application, 3000)

	// Set up expectations
	mockRoleRepository.On("DeleteRole", mock.Anything, mock.Anything).Return(nil)

	// Call the method
	resp, err := roleService.DeleteRole(context.Background(), &roles.DeleteRoleRequest{Id: "12345"})

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "12345", resp.Role.Id)

	// Ensure that the expectations were met
	mockRoleRepository.AssertExpectations(t)
}
