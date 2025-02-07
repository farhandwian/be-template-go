package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"time"
)

type ChartDebitReq struct {
	WaterChannelDoorID int       `json:"water_channel_door_id"`
	MinTime            time.Time `json:"min_time"`
	MaxTime            time.Time `json:"max_time"`
}

type ChartDebitRes struct {
	Charts []gateway.ActualDebit `json:"charts"`
}

type ChartDebitUseCase = core.ActionHandler[ChartDebitReq, ChartDebitRes]

func ImplChartDebitUseCase(
	getDebit gateway.GetDebit,
) ChartDebitUseCase {

	return func(ctx context.Context, req ChartDebitReq) (*ChartDebitRes, error) {

		debitObj, err := getDebit(ctx, gateway.GetDebitReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
			MinTime:            req.MinTime,
			MaxTime:            req.MaxTime,
		})
		if err != nil {
			return nil, err
		}

		return &ChartDebitRes{
			Charts: debitObj.Debits,
		}, nil
	}

}
