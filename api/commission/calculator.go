package commission

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commission "github.com/NpoolPlatform/archivement-middleware/pkg/commission"
	constant "github.com/NpoolPlatform/archivement-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/archivement-middleware/pkg/tracer"

	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement/commission"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) CalculateOrderCommission(
	ctx context.Context,
	in *npool.CalculateOrderCommissionRequest,
) (*npool.CalculateOrderCommissionResponse, error) {
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

	span = tracer.TraceInvoker(span, "commission", "calculator", "CalculateOrderCommission")

	if err := commission.CalculateOrderCommission(ctx, in.GetOrderID()); err != nil {
		logger.Sugar().Errorw("CalculateOrderCommission", "error", err)
		return &npool.CalculateOrderCommissionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.CalculateOrderCommissionResponse{}, nil
}
