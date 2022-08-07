package archivement

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/NpoolPlatform/archivement-manager/pkg/db"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/general"

	detailcli "github.com/NpoolPlatform/archivement-manager/pkg/client/detail"
	detailcrud "github.com/NpoolPlatform/archivement-manager/pkg/crud/detail"
	detailmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/detail"

	generalcli "github.com/NpoolPlatform/archivement-manager/pkg/client/general"
	generalcrud "github.com/NpoolPlatform/archivement-manager/pkg/crud/general"
	generalmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"

	commonpb "github.com/NpoolPlatform/message/npool"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

func TryCreateGeneral(ctx context.Context, appID, userID, goodID, coinTypeID string) (string, error) {
	key := fmt.Sprintf("archivement-general:%v:%v:%v:%v", appID, userID, goodID, coinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return "", err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	general1, err := generalcli.GetGeneralOnly(ctx, &generalmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: appID,
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: userID,
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: goodID,
		},
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: coinTypeID,
		},
	})
	if err != nil {
		return "", err
	}
	if general1 != nil {
		return general1.ID, nil
	}

	general1, err = generalcli.CreateGeneral(ctx, &generalmgrpb.GeneralReq{
		AppID:      &appID,
		UserID:     &userID,
		GoodID:     &goodID,
		CoinTypeID: &coinTypeID,
	})
	if err != nil {
		return "", err
	}

	return general1.ID, nil
}

func detailKey(in *detailmgrpb.DetailReq) string {
	return fmt.Sprintf("archivement-detail:%v:%v:%v:%v",
		in.GetAppID(),
		in.GetUserID(),
		in.GetGoodID(),
		in.GetOrderID(),
	)
}

func BookKeeping(ctx context.Context, in *detailmgrpb.DetailReq) error { //nolint
	val, err := decimal.NewFromString(in.GetAmount())
	if err != nil {
		return err
	}
	if val.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("invalid amount")
	}

	var generalID string
	if generalID, err = TryCreateGeneral(
		ctx,
		in.GetAppID(), in.GetUserID(), in.GetGoodID(), in.GetCoinTypeID(),
	); err != nil {
		return err
	}

	key := detailKey(in)
	if err := redis2.TryLock(key, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	exist, err := detailcli.ExistDetailConds(ctx, &detailmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		UserID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetUserID(),
		},
		GoodID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetGoodID(),
		},
		OrderID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetOrderID(),
		},
	})
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("already exist")
	}

	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		c1, err := detailcrud.CreateSet(tx.Detail.Create(), in)
		if err != nil {
			return err
		}

		_, err = c1.Save(ctx)
		if err != nil {
			return err
		}

		info, err := tx.
			General.
			Query().
			Where(
				general.ID(uuid.MustParse(generalID)),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		selfUnits := uint32(0)
		selfAmount := decimal.NewFromInt(0).String()
		selfCommission := decimal.NewFromInt(0).String()

		if in.GetSelfOrder() {
			selfUnits += *in.Units
			selfAmount = in.GetUSDAmount()
			selfCommission = in.GetCommission()
		}

		c2, err := generalcrud.UpdateSet(info, &generalmgrpb.GeneralReq{
			AppID:           in.AppID,
			UserID:          in.UserID,
			GoodID:          in.GoodID,
			CoinTypeID:      in.CoinTypeID,
			TotalUnits:      in.Units,
			SelfUnits:       &selfUnits,
			TotalAmount:     in.USDAmount,
			SelfAmount:      &selfAmount,
			TotalCommission: in.Commission,
			SelfCommission:  &selfCommission,
		})
		if err != nil {
			return err
		}

		_, err = c2.Save(ctx)
		return err
	})
}
