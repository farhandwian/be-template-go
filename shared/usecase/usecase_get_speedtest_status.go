package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type GetSpeedtestStatusReq struct {
}

type GetSpeedtestStatusRes struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
	Ping     float64 `json:"ping"`
}

type GetSpeedtestStatusUseCase = core.ActionHandler[GetSpeedtestStatusReq, GetSpeedtestStatusRes]

func ImplGetSpeedtestStatusUseCase(getSpeedtestGateway gateway.GetSpeedtestStatusGateway) GetSpeedtestStatusUseCase {
	return func(ctx context.Context, req GetSpeedtestStatusReq) (*GetSpeedtestStatusRes, error) {

		res, err := getSpeedtestGateway(ctx, gateway.GetSpeedtestStatusReq{})
		if err != nil {
			return nil, err
		}

		var speedtestResp GetSpeedtestStatusRes
		if res != nil {
			speedtestResp.Ping = res.SpeedtestResponse.Data.Ping
			speedtestResp.Download = res.SpeedtestResponse.Data.Download
			speedtestResp.Upload = res.SpeedtestResponse.Data.Upload
		} else {
			speedtestResp.Ping = 0
			speedtestResp.Download = 0
			speedtestResp.Upload = 0
		}

		return &speedtestResp, nil

	}
}
