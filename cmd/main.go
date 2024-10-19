package main

import (
	"os"

	"github.com/andrei-kozel/go-utils/utils/prettylog"
	"github.com/andrei-kozel/owly-roles/internal/adapters/db"
	"github.com/andrei-kozel/owly-roles/internal/adapters/grpc"
	"github.com/andrei-kozel/owly-roles/internal/application/core/api"
	"github.com/andrei-kozel/owly-roles/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	var attr string
	if len(os.Args) > 1 {
		attr = os.Args[1]
	}
	config.Configurations(attr)

	// Setup the logger
	log := prettylog.SetupLoggger(config.AppConfig.Env)
	log.Info("Service started", "config", config.AppConfig)

	// Connect to the database
	db, err := db.NewRoleRepository(config.AppConfig.PostgresUrl, log)
	if err != nil {
		log.Error("Failed to open the database", "error", err)
		panic(err)
	}
	log.Info("Connected to the database")

	application := api.NewApplication(db)

	grpcAdapter := grpc.NewRoleService(application, config.AppConfig.ApplicationPort)
	grpcAdapter.Start()
}
