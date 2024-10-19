package grpc

import (
	"github.com/andrei-kozel/owly-proto/golang/roles"
	"github.com/andrei-kozel/owly-roles/internal/ports"
)

type Adapter struct {
	roles.UnimplementedRolesServiceServer
	api  ports.APIport
	port int
}
