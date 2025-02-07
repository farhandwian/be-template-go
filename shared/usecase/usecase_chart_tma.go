package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"time"
)

type ChartTMAReq struct {
	WaterChannelDoorID int       `json:"water_channel_door_id"`
	MinTime            time.Time `json:"min_time"`
	MaxTime            time.Time `json:"max_time"`
}

type ChartTMARes struct {
	Charts []gateway.WaterSurfaceElevation `json:"charts"`
}

type ChartTMAUseCase = core.ActionHandler[ChartTMAReq, ChartTMARes]

func ImplChartTMAUseCase(
	getWaterElevation gateway.GetWaterSurfaceElevation,
) ChartTMAUseCase {

	return func(ctx context.Context, req ChartTMAReq) (*ChartTMARes, error) {

		// Query TMA
		weObj, err := getWaterElevation(ctx, gateway.GetWaterSurfaceElevationReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
			MinTime:            req.MinTime,
			MaxTime:            req.MaxTime,
		})
		if err != nil {
			return nil, err
		}

		return &ChartTMARes{
			Charts: weObj.Results,
		}, nil
	}

}
