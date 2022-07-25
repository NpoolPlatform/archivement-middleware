package detail

import (
	"github.com/NpoolPlatform/message/npool/archivementmgr/detail"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	detail.UnimplementedArchivementDetailServer
}

func Register(server grpc.ServiceRegistrar) {
	detail.RegisterArchivementDetailServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
