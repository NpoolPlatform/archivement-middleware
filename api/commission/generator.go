package commission

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/archivementmw/commission"
)

func (s *Server) CreateUserGoodCommissions(
	ctx context.Context,
	in *npool.CreateUserGoodCommissionsRequest,
) (*npool.CreateUserGoodCommissionsResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) CreateAppUserGoodCommissions(
	ctx context.Context,
	in *npool.CreateAppUserGoodCommissionsRequest,
) (*npool.CreateAppUserGoodCommissionsResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}
