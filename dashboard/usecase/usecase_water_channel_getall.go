package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
)

type GetAllWaterChannelUseCaseReq struct {
}

type GetAllWaterChannelUseCaseRes struct {
	WaterChannel []WaterChannel `json:"water_channel"`
}

type WaterChannel struct {
	ID         string `json:"id"`
	ExternalID int    `json:"external_id"`
	Name       string `json:"name"`
}

type ListWaterChannelUseCase = core.ActionHandler[GetAllWaterChannelUseCaseReq, GetAllWaterChannelUseCaseRes]

func ImplListWaterChannelUseCase(getAllWaterChannel gateway.GetWaterChannelListGateway) ListWaterChannelUseCase {
	return func(ctx context.Context, req GetAllWaterChannelUseCaseReq) (*GetAllWaterChannelUseCaseRes, error) {

		res, err := getAllWaterChannel(ctx, gateway.GetWaterChannelListReq{})
		if err != nil {
			return nil, err
		}

		waterChannels := make([]WaterChannel, 0)
		for _, waterChannel := range res.WaterChannels {
			waterChannels = append(waterChannels, WaterChannel{
				ID:         waterChannel.ID,
				ExternalID: waterChannel.ExternalID,
				Name:       waterChannel.Name,
			})
		}

		return &GetAllWaterChannelUseCaseRes{WaterChannel: waterChannels}, nil
	}
}
