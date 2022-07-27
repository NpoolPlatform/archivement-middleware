package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/archivementmw"

	"github.com/NpoolPlatform/archivement-middleware/api/archivement"
	"github.com/NpoolPlatform/archivement-middleware/api/commission"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	archivementmw.UnimplementedArchivementMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	archivementmw.RegisterArchivementMiddlewareServer(server, &Server{})
	commission.Register(server)
	archivement.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := archivementmw.RegisterArchivementMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := commission.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := archivement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
