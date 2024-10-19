package main

import (
	"github.com/andrei-kozel/go-utils/utils/prettylog"
	"github.com/andrei-kozel/owly-roles/internal/adapters/db"
	"github.com/andrei-kozel/owly-roles/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	// Load the config
	config.Configurations()

	// Setup the logger
	log := prettylog.SetupLoggger(config.AppConfig.Env)
	log.Info("Service started", "config", config.AppConfig)

	// Connect to the database
	_, err := db.NewAdapter(config.AppConfig.PostgresUrl)
	if err != nil {
		log.Error("Failed to open the database", "error", err)
		panic(err)
	}

	log.Info("Connected to the database")
}
