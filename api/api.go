package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/commissionmgr"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	commissionmgr.UnimplementedCommissionManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	commissionmgr.RegisterCommissionManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return commissionmgr.RegisterCommissionManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
