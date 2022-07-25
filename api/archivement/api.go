package archivement

import (
	"github.com/NpoolPlatform/message/npool/archivementmgr/archivement"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	archivement.UnimplementedArchivementServer
}

func Register(server grpc.ServiceRegistrar) {
	archivement.RegisterArchivementServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
