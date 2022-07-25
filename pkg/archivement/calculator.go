package archivement

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/shopspring/decimal"

	ordercli "github.com/NpoolPlatform/cloud-hashing-order/pkg/client"
	orderconst "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"
	orderpb "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	goodscli "github.com/NpoolPlatform/cloud-hashing-goods/pkg/client"
	goodspb "github.com/NpoolPlatform/message/npool/cloud-hashing-goods"

	commonpb "github.com/NpoolPlatform/message/npool"
	detailpb "github.com/NpoolPlatform/message/npool/archivementmgr/detail"
	generalpb "github.com/NpoolPlatform/message/npool/archivementmgr/general"

	detailcrud "github.com/NpoolPlatform/archivement-manager/pkg/crud/detail"
	generalcrud "github.com/NpoolPlatform/archivement-manager/pkg/crud/general"

	constant "github.com/NpoolPlatform/archivement-manager/pkg/message/const"
	"github.com/NpoolPlatform/archivement-manager/pkg/referral"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"
)

func calculateArchivement(ctx context.Context, order *orderpb.Order, payment *orderpb.Payment, good *goodspb.GoodInfo) error {
	inviters, _, err := referral.GetReferrals(ctx, order.AppID, order.UserID)
	if err != nil {
		return err
	}

	amount := decimal.NewFromFloat(payment.Amount).String()
	usdAmount := decimal.NewFromFloat(payment.Amount).Mul(decimal.NewFromFloat(payment.CoinUSDCurrency)).String()
	currency := decimal.NewFromFloat(payment.CoinUSDCurrency).String()

	for _, inviter := range inviters {
		myInviter := inviter

		_, err = detailcrud.Create(ctx, &detailpb.DetailReq{
			AppID:                  &payment.AppID,
			UserID:                 &myInviter,
			GoodID:                 &order.GoodID,
			OrderID:                &order.ID,
			PaymentID:              &payment.ID,
			CoinTypeID:             &good.CoinInfoID,
			PaymentCoinTypeID:      &payment.CoinInfoID,
			PaymentCoinUSDCurrency: &currency,
			Units:                  &order.Units,
			Amount:                 &amount,
			USDAmount:              &usdAmount,
		})
		if err != nil {
			return err
		}

		general, err := generalcrud.RowOnly(ctx, &generalpb.Conds{
			AppID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: payment.AppID,
			},
			UserID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: payment.UserID,
			},
			GoodID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: payment.GoodID,
			},
			CoinTypeID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: good.CoinInfoID,
			},
		})
		if err != nil {
			return err
		}

		selfUnits := uint32(0)
		if inviter == payment.UserID {
			selfUnits = order.Units
		}

		if general == nil {
			_, err = generalcrud.Create(ctx, &generalpb.GeneralReq{
				AppID:      &payment.AppID,
				UserID:     &myInviter,
				GoodID:     &order.GoodID,
				CoinTypeID: &good.CoinInfoID,
				Amount:     &amount,
				TotalUnits: &order.Units,
				SelfUnits:  &selfUnits,
			})
			if err != nil {
				return err
			}
			continue
		}

		generalID := general.ID.String()

		_, err = generalcrud.AddFields(ctx, &generalpb.GeneralReq{
			ID:         &generalID,
			Amount:     &amount,
			TotalUnits: &order.Units,
			SelfUnits:  &selfUnits,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func CalculateOrderArchivement(ctx context.Context, orderID string) error {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGeneral")
	defer span.End()

	span.SetAttributes(attribute.String("OrderID", orderID))

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	order, err := ordercli.GetOrder(ctx, orderID)
	if err != nil {
		return err
	}

	good, err := goodscli.GetGood(ctx, order.GoodID)
	if err != nil {
		return err
	}

	payment, err := ordercli.GetOrderPayment(ctx, orderID)
	if err != nil {
		return err
	}

	switch payment.State {
	case orderconst.PaymentStateDone:
	default:
		logger.Sugar().Errorw("CalculateOrderArchivement", "payment", payment.ID, "state", payment.State)
		return fmt.Errorf("invalid payment state")
	}

	if err := calculateArchivement(ctx, order, payment, good); err != nil {
		return err
	}

	return nil
}
