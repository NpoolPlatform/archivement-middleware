package general

import (
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/archivementmgr/general"
)

func Ent2Grpc(row *ent.General) *npool.General {
	if row == nil {
		return nil
	}

	return &npool.General{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		GoodID:     row.GoodID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		Amount:     row.Amount.String(),
		TotalUnits: row.TotalUnits,
		SelfUnits:  row.SelfUnits,
	}
}

func Ent2GrpcMany(rows []*ent.General) []*npool.General {
	infos := []*npool.General{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
