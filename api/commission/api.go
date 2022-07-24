package commission

import (
	"context"

	"github.com/NpoolPlatform/message/npool/archivementmgr/commission"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	commission.UnimplementedCommissionServer
}

func Register(server grpc.ServiceRegistrar) {
	commission.RegisterCommissionServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return commission.RegisterCommissionHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
