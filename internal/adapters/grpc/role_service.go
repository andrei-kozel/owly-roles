package grpc

import (
	"context"

	"github.com/andrei-kozel/owly-common/domain"
	"github.com/andrei-kozel/owly-proto/golang/roles"
	"github.com/andrei-kozel/owly-roles/internal/ports"
	"github.com/google/uuid"
)

type RoleService struct {
	roles.UnimplementedRolesServiceServer
	api  ports.APIport
	port int
}

func NewRoleService(api ports.APIport, port int) *RoleService {
	return &RoleService{api: api, port: port}
}

func (r RoleService) CreateRole(ctx context.Context, role *roles.CreateRoleRequest) (*roles.RoleResponse, error) {
	newRole := domain.Role{
		Guid:        uuid.NewString(),
		Name:        role.Name,
		Description: role.Description,
	}

	addedRole, err := r.api.AddRole(ctx, &newRole)
	if err != nil {
		return nil, err
	}

	return &roles.RoleResponse{
		Role: &roles.Role{
			Id:          addedRole.Guid,
			Name:        addedRole.Name,
			Description: addedRole.Description,
		},
	}, nil
}

func (r RoleService) DeleteRole(ctx context.Context, role *roles.DeleteRoleRequest) (*roles.RoleResponse, error) {
	err := r.api.DeleteRole(ctx, role.Id)
	if err != nil {
		return nil, err
	}

	return &roles.RoleResponse{
		Role: &roles.Role{
			Id: role.Id,
		},
	}, nil
}

func (r RoleService) GetRoles(ctx context.Context, _ *roles.GetRolesRequest) (*roles.GetRolesResponse, error) {
	response, err := r.api.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	rls := make([]*roles.Role, 0, len(response))
	for _, role := range response {
		rls = append(rls, &roles.Role{
			Id:          role.Guid,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	return &roles.GetRolesResponse{
		Roles: rls,
	}, nil
}

func (r RoleService) GetRole(ctx context.Context, role *roles.GetRoleRequest) (*roles.GetRoleResponse, error) {
	response, err := r.api.GetRole(ctx, role.Id)
	if err != nil {
		return nil, err
	}

	return &roles.GetRoleResponse{
		Role: &roles.Role{
			Id:          response.Guid,
			Name:        response.Name,
			Description: response.Description,
		},
	}, nil
}
