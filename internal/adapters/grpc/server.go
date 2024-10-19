package grpc

import (
	"fmt"
	"net"

	"github.com/andrei-kozel/go-utils/utils/prettylog"
	"github.com/andrei-kozel/owly-proto/golang/roles"
	"github.com/andrei-kozel/owly-roles/internal/config"
	"google.golang.org/grpc"
)

func (r *RoleService) Start() {
	var err error

	log := prettylog.SetupLoggger(config.AppConfig.Env)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", r.port))
	if err != nil {
		log.Error("failed to listen", "error", err)
	}
	log.Info("listening", "port", r.port)

	grpcServer := grpc.NewServer()
	roles.RegisterRolesServiceServer(grpcServer, r)
	if err := grpcServer.Serve(listen); err != nil {
		log.Error("failed to serve", "error", err)
	}
}
