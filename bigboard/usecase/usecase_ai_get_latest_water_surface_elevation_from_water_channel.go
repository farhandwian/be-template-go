package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
)

type AIGetWaterSurfaceElevationFromWaterChannelReq struct {
	WaterChannelDoorID int `json:"water_channel_door_id"`
}

type AIGetWaterSurfaceElevationFromWaterChannelRes struct {
	WaterChannelDoorID    int     `json:"water_channel_door_id"`
	Name                  string  `json:"name"`
	WaterSurfaceElevation float32 `json:"water_surface_elevation"`
}

type AIGetWaterSurfaceElevationFromWaterChannelUseCase = core.ActionHandler[AIGetWaterSurfaceElevationFromWaterChannelReq, AIGetWaterSurfaceElevationFromWaterChannelRes]

func ImplAIGetWaterSurfaceElevationFromWaterChannel(getWaterSurfaceElevationGateway sharedGateway.GetLatestWaterSurfaceElevationByDoorID, getWaterChannelDoorDetailGateway sharedGateway.GetWaterChannelDoorByID, sendSSEGateway sharedGateway.SendSSEMessage) AIGetWaterSurfaceElevationFromWaterChannelUseCase {
	return func(ctx context.Context, req AIGetWaterSurfaceElevationFromWaterChannelReq) (*AIGetWaterSurfaceElevationFromWaterChannelRes, error) {

		waterChannelDoorDetail, err := getWaterChannelDoorDetailGateway(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		waterSurfaceElevation, err := getWaterSurfaceElevationGateway(ctx, sharedGateway.GetLatestWaterSurfaceElevationByDoorIDReq{
			WaterChannelDoorID: waterChannelDoorDetail.WaterChannelDoor.ExternalID,
		})
		if err != nil {
			return nil, err
		}

		result := &AIGetWaterSurfaceElevationFromWaterChannelRes{
			WaterChannelDoorID:    waterSurfaceElevation.WaterSurfaceElevation.WaterChannelDoorID,
			Name:                  waterChannelDoorDetail.WaterChannelDoor.Name,
			WaterSurfaceElevation: waterSurfaceElevation.WaterSurfaceElevation.WaterLevel,
		}

		if _, err := sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "get-latest-water-surface-elevation-from-water-channel-by-id",
			FunctionName: "showLatestWaterSurfaceElevation",
			Data:         result,
		}); err != nil {
			return nil, err
		}

		return result, nil
	}
}
