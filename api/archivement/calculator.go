package archivement

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	archivement "github.com/NpoolPlatform/archivement-manager/pkg/archivement"
	constant "github.com/NpoolPlatform/archivement-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/archivement-manager/pkg/tracer"

	npool "github.com/NpoolPlatform/message/npool/archivementmgr/archivement"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) CalculateOrderArchivement(
	ctx context.Context,
	in *npool.CalculateOrderArchivementRequest,
) (*npool.CalculateOrderArchivementResponse, error) {
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
		logger.Sugar().Errorw("CalculateOrderArchivement", "error", err)
		return &npool.CalculateOrderArchivementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceInvoker(span, "archivement", "calculator", "CalculateOrderArchivement")

	if err := archivement.CalculateOrderArchivement(ctx, in.GetOrderID()); err != nil {
		logger.Sugar().Errorw("CalculateOrderArchivement", "error", err)
		return &npool.CalculateOrderArchivementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.CalculateOrderArchivementResponse{}, nil
}
