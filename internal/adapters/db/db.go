package db

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	id           string
	name         string
	descriptioan string
}

type Adapter struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Role{})
	if err != nil {
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) Get(id string) (*Role, error) {
	var role Role
	result := a.db.First(&role, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (a *Adapter) Save(role *Role) error {
	result := a.db.Save(role)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
