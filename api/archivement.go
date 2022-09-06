package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement"

	"github.com/NpoolPlatform/archivement-manager/api/detail"

	archivement1 "github.com/NpoolPlatform/archivement-middleware/pkg/archivement"

	errno "github.com/NpoolPlatform/archivement-middleware/pkg/errno"

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
		if errors.Is(err, errno.ErrAlreadyExists) {
			return &npool.BookKeepingResponse{}, status.Error(codes.AlreadyExists, err.Error())
		}
		return &npool.BookKeepingResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.BookKeepingResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, in *npool.DeleteRequest) (*npool.DeleteResponse, error) {
	if _, err := uuid.Parse(in.GetOrderID()); err != nil {
		logger.Sugar().Errorw("validate", "OrderID", in.GetOrderID(), "error", err)
		return &npool.DeleteResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if err := archivement1.Delete(ctx, in.GetOrderID()); err != nil {
		return &npool.DeleteResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteResponse{}, nil
}
