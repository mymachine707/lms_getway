package handlars

import (
	"mymachine707/clients"
	"mymachine707/config"
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
