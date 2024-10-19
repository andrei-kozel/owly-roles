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
