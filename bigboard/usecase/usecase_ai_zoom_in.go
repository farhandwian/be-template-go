package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type AiZoomInReq struct {
	WaterChannelDoorID int `json:"water_channel_door_id"`
}

type AiZoomInRes struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Name      string `json:"name"`
}

type AiZoomInUseCase = core.ActionHandler[AiZoomInReq, AiZoomInRes]

func ImplAiZoomInUseCase(getWaterChannelDoorByID gateway.GetWaterChannelDoorByID, sendSSEMessageGateway gateway.SendSSEMessage) AiZoomInUseCase {
	return func(ctx context.Context, req AiZoomInReq) (*AiZoomInRes, error) {
		waterChannelDoor, err := getWaterChannelDoorByID(ctx, gateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		resData := &AiZoomInRes{
			Longitude: waterChannelDoor.WaterChannelDoor.Longitude,
			Latitude:  waterChannelDoor.WaterChannelDoor.Latitude,
			Name:      waterChannelDoor.WaterChannelDoor.Name,
		}

		_, _ = sendSSEMessageGateway(ctx, gateway.SendSSEMessageReq{
			Subject:      "zoom-in",
			FunctionName: "zoomIn",
			Data:         resData,
		})
		return resData, nil

	}
}
