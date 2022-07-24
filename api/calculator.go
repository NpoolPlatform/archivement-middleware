package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	ordercli "github.com/NpoolPlatform/cloud-hashing-order/pkg/client"

	npool "github.com/NpoolPlatform/message/npool/commissionmgr"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CalculateOrderCommission(ctx context.Context, in *npool.CalculateOrderCommissionRequest) (*npool.CalculateOrderCommissionResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGeneral")
	defer span.End()

	span.SetAttributes(attribute.String("OrderID", in.GetOrderID()))

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	_, err = uuid.Parse(in.GetOrderID())
	if err != nil {
		logger.Sugar().Errorw("CalculateOrderCommission", "error", err)
		return &npool.CalculateOrderCommissionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err := ordercli.GetOrder(ctx, in.GetOrderID())
	if err != nil {
		logger.Sugar().Errorw("CalculateOrderCommission", "error", err)
		return &npool.CalculateOrderCommissionResponse{}, status.Error(codes.Internal, err.Error())
	}

	payment, err := ordercli.GetOrderPayment(ctx, in.GetOrderID())
	if err != nil {
		logger.Sugar().Errorw("CalculateOrderCommission", "error", err)
		return &npool.CalculateOrderCommissionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
