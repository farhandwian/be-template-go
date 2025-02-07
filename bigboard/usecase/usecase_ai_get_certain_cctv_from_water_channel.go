package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
	"shared/usecase"
)

type AIGetCertainCCTVFromWaterChannelReq struct {
	WaterChannelDoorID int   `json:"water_channel_door_id"`
	CCTVIndex          []int `json:"cctv_index"`
}

type AIGetCertainCCTVFromWaterChannelRes struct {
	WaterChannelDoorID int            `json:"water_channel_door_id"`
	Name               string         `json:"name"`
	CCTVDetail         []usecase.CCTV `json:"cctv"`
}

type AIGetCertainCCTVFromWaterChannelUseCase = core.ActionHandler[AIGetCertainCCTVFromWaterChannelReq, AIGetCertainCCTVFromWaterChannelRes]

func ImplAIGetCertainCCTVFromWaterChannel(
	getWaterChannelDoorDevicesGateway sharedGateway.GetWaterChannelDevicesByDoorID,
	getWaterChannelDoorDetailGateway sharedGateway.GetWaterChannelDoorByID,
	sendSSEGateway sharedGateway.SendSSEMessage,
) AIGetCertainCCTVFromWaterChannelUseCase {
	return func(ctx context.Context, req AIGetCertainCCTVFromWaterChannelReq) (*AIGetCertainCCTVFromWaterChannelRes, error) {

		cctvList := make([]usecase.CCTV, 0)
		waterChannelDoorDetail, err := getWaterChannelDoorDetailGateway(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		devices, err := getWaterChannelDoorDevicesGateway(ctx, sharedGateway.GetWaterChannelDevicesByDoorIDReq{WaterChannelDoorID: waterChannelDoorDetail.WaterChannelDoor.ExternalID})
		if err != nil {
			return nil, err
		}

		for _, device := range devices.Devices {
			if device.Category == "cctv" {
				cctvList = append(cctvList, usecase.CCTV{
					ExternalID: device.ExternalID,
					Name:       device.Name,
					IPAddress:  device.IPAddress,
				})
			}
		}

		var filteredCCTVList []usecase.CCTV
		for _, cctvReqIndex := range req.CCTVIndex {
			if cctvReqIndex-1 >= 0 && cctvReqIndex-1 < len(cctvList) {
				filteredCCTVList = append(filteredCCTVList, cctvList[cctvReqIndex-1])
			}
		}

		result := &AIGetCertainCCTVFromWaterChannelRes{
			WaterChannelDoorID: waterChannelDoorDetail.WaterChannelDoor.ExternalID,
			Name:               waterChannelDoorDetail.WaterChannelDoor.Name,
			CCTVDetail:         filteredCCTVList,
		}

		_, _ = sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "get-certain-cctv-from-water-channel",
			FunctionName: "showCertainCCTV",
			Data:         result,
		})

		return result, nil
	}
}
