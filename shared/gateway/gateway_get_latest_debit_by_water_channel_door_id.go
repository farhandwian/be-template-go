package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetLatestDebitReq struct {
	WaterChannelDoorID int
}

type GetLatestDebitRes struct {
	Debit model.ActualDebitData
}

type GetLatestDebit = core.ActionHandler[GetLatestDebitReq, GetLatestDebitRes]

func ImplGetLatestDebit(db *gorm.DB) GetLatestDebit {
	return func(ctx context.Context, request GetLatestDebitReq) (*GetLatestDebitRes, error) {
		var debit model.ActualDebitData

		result := db.Raw(`
			SELECT water_channel_door_id, latest_debit as actual_debit, timestamp
			FROM latest_actual_debits 
			WHERE water_channel_door_id = ?
			LIMIT 1
		`, request.WaterChannelDoorID).Scan(&debit)

		if result.Error != nil {
			return nil, core.NewInternalServerError(result.Error)
		}

		return &GetLatestDebitRes{
			Debit: debit,
		}, nil
	}
}
