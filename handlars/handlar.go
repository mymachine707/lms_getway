package handlars

import (
	"lms/lms_getway/clients"
	"lms/lms_getway/config"
)

// Handler ...
type handler struct {
	cfg        config.Config
	grpcClient *clients.GrpcClients
}

func NewHandler(cfg config.Config, grpcClient *clients.GrpcClients) handler {
	return handler{
		cfg:        cfg,
		grpcClient: grpcClient,
	}
}
