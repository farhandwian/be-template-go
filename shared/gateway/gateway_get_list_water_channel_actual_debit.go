package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetListWaterChannelActualDebitReq struct{}

type GetListWaterChannelActualDebitRes struct {
	WaterChannelActualDebit []model.ActualDebitData
}

type GetListWaterChannelActualDebitGateway = core.ActionHandler[GetListWaterChannelActualDebitReq, GetListWaterChannelActualDebitRes]

func ImplGetListWaterChannelActualDebit(tsDB *gorm.DB) GetListWaterChannelActualDebitGateway {
	return func(ctx context.Context, request GetListWaterChannelActualDebitReq) (*GetListWaterChannelActualDebitRes, error) {

		var actualDebits []model.ActualDebitData

		err := tsDB.Raw(`
			SELECT timestamp,
				water_channel_door_id,
				latest_actual_debits.latest_debit as actual_debit
			FROM latest_actual_debits;
		`).Scan(&actualDebits).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListWaterChannelActualDebitRes{
			WaterChannelActualDebit: actualDebits,
		}, nil
	}
}
