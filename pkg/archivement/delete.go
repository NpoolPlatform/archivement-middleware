package archivement

import (
	"context"
	"time"

	detailcrud "github.com/NpoolPlatform/archivement-manager/pkg/crud/detail"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	generalcrud "github.com/NpoolPlatform/archivement-manager/pkg/crud/general"
	generalmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"

	"github.com/NpoolPlatform/archivement-manager/pkg/db"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/general"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/shopspring/decimal"
)

func Delete(ctx context.Context, orderID string) error {
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		details, _, err := detailcrud.Rows(ctx, &detailmgrpb.Conds{
			OrderID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: orderID,
			},
		}, 0, 0)

		for _, val := range details {
			generalInfo, err := generalcrud.RowOnly(ctx, &generalmgrpb.Conds{
				AppID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: val.AppID.String(),
				},
				UserID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: val.UserID.String(),
				},
				GoodID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: val.GoodID.String(),
				},
				CoinTypeID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: val.CoinTypeID.String(),
				},
			})
			if err != nil {
				return err
			}
			if generalInfo == nil {
				logger.Sugar().Errorw("Delete", "detail", val, "err", "details incorrect data")
				continue
			}

			info, err := tx.
				General.
				Query().
				Where(
					general.ID(generalInfo.ID),
				).
				ForUpdate().
				Only(ctx)
			if err != nil {
				return err
			}

			units := -val.Units
			usdAmount := val.UsdAmount.Neg().String()

			selfUnits := uint32(0)
			selfAmount := decimal.NewFromInt(0).String()

			if val.SelfOrder {
				selfUnits += units
				selfAmount = usdAmount
			}

			c2, err := generalcrud.UpdateSet(info, &generalmgrpb.GeneralReq{
				TotalUnits:  &units,
				SelfUnits:   &selfUnits,
				TotalAmount: &usdAmount,
				SelfAmount:  &selfAmount,
			})
			if err != nil {
				return err
			}
			_, err = c2.Save(ctx)
			if err != nil {
				return err
			}

			err = tx.Detail.
				UpdateOneID(val.ID).
				SetDeletedAt(uint32(time.Now().Unix())).
				Exec(ctx)
			if err != nil {
				return err
			}
		}
		return err
	})
}
