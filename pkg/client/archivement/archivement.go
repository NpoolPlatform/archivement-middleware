//nolint
package archivement

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	detailpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"
	npool "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement"

	constant "github.com/NpoolPlatform/archivement-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func BookKeeping(ctx context.Context, in *detailpb.DetailReq) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.BookKeeping(ctx, &npool.BookKeepingRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}

func Delete(ctx context.Context, orderID string) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.Delete(ctx, &npool.DeleteRequest{
			OrderID: orderID,
		})
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	return err
}
