package db

import (
	"context"

	"github.com/andrei-kozel/owly-roles/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Guid        string
	Name        string
	Description string
}

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(dataSourceUrl string) (*RoleRepository, error) {
	db, err := gorm.Open(postgres.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Role{})
	if err != nil {
		return nil, err
	}

	return &RoleRepository{db: db}, nil
}

func (r *RoleRepository) GetRole(ctx context.Context, id string) (*domain.Role, error) {
	var role domain.Role
	result := r.db.WithContext(ctx).First(&role, "guid = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

func (r *RoleRepository) AddRole(context context.Context, role *domain.Role) (*domain.Role, error) {
	newRole := Role{
		Guid:        role.Guid,
		Name:        role.Name,
		Description: role.Description,
	}

	result := r.db.Create(&newRole)
	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.Role{
		Guid:        newRole.Guid,
		Name:        newRole.Name,
		Description: newRole.Description,
	}, nil
}

func (r *RoleRepository) DeleteRole(ctx context.Context, id string) error {
	result := r.db.Delete(&Role{}, "guid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RoleRepository) GetRoles(ctx context.Context) ([]*domain.Role, error) {
	var roles []Role
	result := r.db.WithContext(ctx).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	var domainRoles []*domain.Role
	for _, role := range roles {
		domainRoles = append(domainRoles, &domain.Role{
			Guid:        role.Guid,
			Name:        role.Name,
			Description: role.Description,
		})
	}

	return domainRoles, nil
}
