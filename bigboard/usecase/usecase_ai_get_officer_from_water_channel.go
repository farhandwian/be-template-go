package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
	"shared/usecase"
)

type AIGetOfficerFromWaterChannelReq struct {
	WaterChannelDoorID int `json:"water_channel_door_id"`
}

type AIGetOfficerFromWaterChannelRes struct {
	WaterChannelDoorID int               `json:"water_channel_door_id"`
	Name               string            `json:"name"`
	Officers           []usecase.Officer `json:"officers"`
}

type AIGetOfficerFromWaterChannelUseCase = core.ActionHandler[AIGetOfficerFromWaterChannelReq, AIGetOfficerFromWaterChannelRes]

func ImplAIGetOfficersFromWaterChannel(getOfficerGateway sharedGateway.GetWaterChannelOfficersByDoorID, getWaterChannelDoorDetailGateway sharedGateway.GetWaterChannelDoorByID, sendSSEGateway sharedGateway.SendSSEMessage) AIGetOfficerFromWaterChannelUseCase {
	return func(ctx context.Context, req AIGetOfficerFromWaterChannelReq) (*AIGetOfficerFromWaterChannelRes, error) {

		officerDetail, err := getOfficerGateway(ctx, sharedGateway.GetWaterChannelOfficersByDoorIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		waterChannelDoorDetail, err := getWaterChannelDoorDetailGateway(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		officerList := make([]usecase.Officer, 0)
		for _, officer := range officerDetail.Officers {
			officerList = append(officerList, usecase.Officer{
				Name:        officer.Name,
				Photo:       officer.Photo,
				PhoneNumber: officer.PhoneNumber,
				Task:        officer.Task,
			})
		}

		result := &AIGetOfficerFromWaterChannelRes{
			WaterChannelDoorID: waterChannelDoorDetail.WaterChannelDoor.ExternalID,
			Name:               waterChannelDoorDetail.WaterChannelDoor.Name,
			Officers:           officerList,
		}

		if _, err := sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "get-officer-from-water-channel-by-id",
			FunctionName: "showOfficer",
			Data:         result,
		}); err != nil {
			return nil, err
		}

		return result, nil
	}
}
