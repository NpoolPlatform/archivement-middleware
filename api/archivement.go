package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement"

	"github.com/NpoolPlatform/archivement-manager/api/detail"

	archivement1 "github.com/NpoolPlatform/archivement-middleware/pkg/archivement"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) BookKeeping(ctx context.Context, in *npool.BookKeepingRequest) (*npool.BookKeepingResponse, error) {
	if err := detail.Validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		return &npool.BookKeepingResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := archivement1.BookKeeping(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("BookKeeping", "error", err)
		return &npool.BookKeepingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}
