package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	archivement.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	archivement.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := archivement.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
