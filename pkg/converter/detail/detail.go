package detail

import (
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/archivementmgr/detail"
)

func Ent2Grpc(row *ent.Detail) *npool.Detail {
	if row == nil {
		return nil
	}

	return &npool.Detail{
		ID:                     row.ID.String(),
		AppID:                  row.AppID.String(),
		UserID:                 row.UserID.String(),
		GoodID:                 row.GoodID.String(),
		OrderID:                row.OrderID.String(),
		PaymentID:              row.PaymentID.String(),
		CoinTypeID:             row.CoinTypeID.String(),
		PaymentCoinTypeID:      row.PaymentCoinTypeID.String(),
		PaymentCoinUSDCurrency: row.PaymentCoinUsdCurrency.String(),
		Units:                  row.Units,
		Amount:                 row.Amount.String(),
		USDAmount:              row.UsdAmount.String(),
		CreatedAt:              row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Detail) []*npool.Detail {
	infos := []*npool.Detail{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
