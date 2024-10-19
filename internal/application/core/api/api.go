package api

import "github.com/andrei-kozel/owly-roles/internal/adapters/db"

type Application struct {
	db *db.Adapter
}

func NewApplication(db *db.Adapter) *Application {
	return &Application{db: db}
}
