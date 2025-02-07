package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
	"shared/usecase"
)

type AIGetCCTVFromWaterChannelReq struct {
	WaterChannelDoorIDs []int `json:"water_channel_door_ids"`
}

type AIGetCCTVFromWaterChannelRes struct {
	Message    string       `json:"message"`
	CCTVDetail []CCTVDetail `json:"cctv"`
}

type CCTVDetail struct {
	WaterChannelDoorID int            `json:"water_channel_door_id"`
	Name               string         `json:"name"`
	CCTV               []usecase.CCTV `json:"cctv"`
}

type AIGetAllCCTVFromWaterChannelUseCase = core.ActionHandler[AIGetCCTVFromWaterChannelReq, AIGetCCTVFromWaterChannelRes]

func ImplAIGetAllCCTVFromWaterChannel(
	getWaterChannelDoorDevicesGateway sharedGateway.GetWaterChannelDevicesByDoorID,
	getWaterChannelDoorDetailGateway sharedGateway.GetWaterChannelDoorByID,
	sendSSEGateway sharedGateway.SendSSEMessage,
) AIGetAllCCTVFromWaterChannelUseCase {
	return func(ctx context.Context, req AIGetCCTVFromWaterChannelReq) (*AIGetCCTVFromWaterChannelRes, error) {

		cctvListData := make([]CCTVDetail, 0)
		for _, doorID := range req.WaterChannelDoorIDs {

			cctvList := make([]usecase.CCTV, 0)
			waterChannelDoorDetail, err := getWaterChannelDoorDetailGateway(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: doorID})
			if err != nil {
				return nil, err
			}

			devices, err := getWaterChannelDoorDevicesGateway(ctx, sharedGateway.GetWaterChannelDevicesByDoorIDReq{WaterChannelDoorID: doorID})
			if err != nil {
				return nil, err
			}
			for _, device := range devices.Devices {
				if device.Category == "cctv" {

					humanDetected := false
					garbageDetected := false
					if device.DetectedObject == "human" {
						humanDetected = true
					} else if device.DetectedObject == "garbage" {
						garbageDetected = true
					}

					cctvList = append(cctvList, usecase.CCTV{
						ExternalID:      device.ExternalID,
						Name:            device.Name,
						IPAddress:       device.IPAddress,
						HumanDetected:   humanDetected,
						GarbageDetected: garbageDetected,
					})
				}
			}
			cctvListData = append(cctvListData, CCTVDetail{
				WaterChannelDoorID: doorID,
				Name:               waterChannelDoorDetail.WaterChannelDoor.Name,
				CCTV:               cctvList,
			})
		}

		_, _ = sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "get-all-cctv-from-water-channel",
			FunctionName: "showAllCCTV",
			Data:         cctvListData,
		})

		return &AIGetCCTVFromWaterChannelRes{
			Message:    "All CCTV has been opened. Wait for few moments if not shown yet.",
			CCTVDetail: cctvListData,
		}, nil
	}
}
