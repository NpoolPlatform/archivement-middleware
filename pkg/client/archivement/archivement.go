package archivement

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/archivement-middleware/pkg/message/const"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/archivementmw/archivement"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.ArchivementClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get good payment connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewArchivementClient(conn)

	return fn(_ctx, cli)
}

func CalculateArchivement(ctx context.Context, orderID string) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.ArchivementClient) (cruder.Any, error) {
		_, err := cli.CalculateOrderArchivement(ctx, &npool.CalculateOrderArchivementRequest{
			OrderID: orderID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail calculate order archivement: %v", err)
		}
		return nil, nil
	})
	if err != nil {
		return fmt.Errorf("fail calculate order archivement: %v", err)
	}
	return nil
}
