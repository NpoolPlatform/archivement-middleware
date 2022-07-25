package detail

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/archivement-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/archivement-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/archivement-manager/pkg/tracer/detail"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/archivement-manager/pkg/db"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent"
	"github.com/NpoolPlatform/archivement-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/archivementmgr/detail"

	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.DetailReq) (*ent.Detail, error) { //nolint
	var info *ent.Detail
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.Debug().Detail.Create()

		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.GoodID != nil {
			c.SetGoodID(uuid.MustParse(in.GetGoodID()))
		}
		if in.OrderID != nil {
			c.SetOrderID(uuid.MustParse(in.GetOrderID()))
		}
		if in.PaymentID != nil {
			c.SetPaymentID(uuid.MustParse(in.GetPaymentID()))
		}
		if in.CoinTypeID != nil {
			c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
		}
		if in.PaymentCoinTypeID != nil {
			c.SetPaymentCoinTypeID(uuid.MustParse(in.GetPaymentCoinTypeID()))
		}
		if in.PaymentCoinUSDCurrency != nil {
			currency, err := decimal.NewFromString(in.GetPaymentCoinUSDCurrency())
			if err != nil {
				return err
			}
			c.SetPaymentCoinUsdCurrency(currency)
		}
		if in.Amount != nil {
			amount, err := decimal.NewFromString(in.GetAmount())
			if err != nil {
				return err
			}
			c.SetAmount(amount)
		}
		if in.USDAmount != nil {
			amount, err := decimal.NewFromString(in.GetUSDAmount())
			if err != nil {
				return err
			}
			c.SetUsdAmount(amount)
		}
		if in.Units != nil {
			c.SetUnits(in.GetUnits())
		}
		if in.CreatedAt != nil {
			c.SetCreatedAt(in.GetCreatedAt())
		}

		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.DetailReq) ([]*ent.Detail, error) { //nolint
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Detail{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.DetailCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.Detail.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.GoodID != nil {
				bulk[i].SetGoodID(uuid.MustParse(info.GetGoodID()))
			}
			if info.OrderID != nil {
				bulk[i].SetOrderID(uuid.MustParse(info.GetOrderID()))
			}
			if info.PaymentID != nil {
				bulk[i].SetPaymentID(uuid.MustParse(info.GetPaymentID()))
			}
			if info.CoinTypeID != nil {
				bulk[i].SetCoinTypeID(uuid.MustParse(info.GetCoinTypeID()))
			}
			if info.PaymentCoinTypeID != nil {
				bulk[i].SetPaymentCoinTypeID(uuid.MustParse(info.GetPaymentCoinTypeID()))
			}
			if info.PaymentCoinUSDCurrency != nil {
				currency, err := decimal.NewFromString(info.GetPaymentCoinUSDCurrency())
				if err != nil {
					return err
				}
				bulk[i].SetPaymentCoinUsdCurrency(currency)
			}
			if info.Amount != nil {
				amount, err := decimal.NewFromString(info.GetAmount())
				if err != nil {
					return err
				}
				bulk[i].SetAmount(amount)
			}
			if info.USDAmount != nil {
				amount, err := decimal.NewFromString(info.GetUSDAmount())
				if err != nil {
					return err
				}
				bulk[i].SetUsdAmount(amount)
			}
			if info.Units != nil {
				bulk[i].SetUnits(info.GetUnits())
			}
			if info.CreatedAt != nil {
				bulk[i].SetCreatedAt(info.GetCreatedAt())
			}
		}
		rows, err = tx.Detail.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Detail, error) {
	var info *ent.Detail
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Detail.Query().Where(detail.ID(id)).Only(_ctx)
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.DetailQuery, error) { //nolint
	stm := cli.Detail.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.GoodID != nil {
		switch conds.GetGoodID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.OrderID != nil {
		switch conds.GetOrderID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.OrderID(uuid.MustParse(conds.GetOrderID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.PaymentID != nil {
		switch conds.GetPaymentID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.PaymentID(uuid.MustParse(conds.GetPaymentID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.PaymentCoinTypeID != nil {
		switch conds.GetPaymentCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(detail.PaymentCoinTypeID(uuid.MustParse(conds.GetPaymentCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.PaymentCoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(conds.GetPaymentCoinUSDCurrency().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetPaymentCoinUSDCurrency().GetOp() {
		case cruder.LT:
			stm.Where(detail.PaymentCoinUsdCurrencyLT(currency))
		case cruder.GT:
			stm.Where(detail.PaymentCoinUsdCurrencyGT(currency))
		case cruder.EQ:
			stm.Where(detail.PaymentCoinUsdCurrencyEQ(currency))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.Amount != nil {
		amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAmount().GetOp() {
		case cruder.LT:
			stm.Where(detail.AmountLT(amount))
		case cruder.GT:
			stm.Where(detail.AmountGT(amount))
		case cruder.EQ:
			stm.Where(detail.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	if conds.Units != nil {
		switch conds.GetUnits().GetOp() {
		case cruder.LT:
			stm.Where(detail.UnitsLT(conds.GetUnits().GetValue()))
		case cruder.GT:
			stm.Where(detail.UnitsGT(conds.GetUnits().GetValue()))
		case cruder.EQ:
			stm.Where(detail.UnitsEQ(conds.GetUnits().GetValue()))
		default:
			return nil, fmt.Errorf("invalid detail field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Detail, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Detail{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(detail.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Detail, error) {
	var info *ent.Detail
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Detail.Query().Where(detail.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Detail, error) {
	var info *ent.Detail
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Detail.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
