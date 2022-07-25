package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/archivementmgr"

	"github.com/NpoolPlatform/archivement-manager/api/archivement"
	"github.com/NpoolPlatform/archivement-manager/api/commission"
	"github.com/NpoolPlatform/archivement-manager/api/detail"
	"github.com/NpoolPlatform/archivement-manager/api/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	archivementmgr.UnimplementedArchivementManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	archivementmgr.RegisterArchivementManagerServer(server, &Server{})
	commission.Register(server)
	archivement.Register(server)
	detail.Register(server)
	general.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := archivementmgr.RegisterArchivementManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := commission.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := archivement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := detail.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := general.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
