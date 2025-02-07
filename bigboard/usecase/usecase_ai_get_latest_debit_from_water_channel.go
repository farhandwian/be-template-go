package usecase

import (
	"context"
	"shared/core"
	sharedGateway "shared/gateway"
)

type AIGetDebitFromWaterChannelReq struct {
	WaterChannelDoorID int `json:"water_channel_door_id"`
}

type AIGetDebitFromWaterChannelRes struct {
	WaterChannelDoorID int     `json:"water_channel_door_id"`
	Name               string  `json:"name"`
	Debit              float64 `json:"debit"`
}

type AIGetDebitFromWaterChannelUseCase = core.ActionHandler[AIGetDebitFromWaterChannelReq, AIGetDebitFromWaterChannelRes]

func ImplAIGetDebitFromWaterChannel(getDebitGateway sharedGateway.GetLatestDebit, getWaterChannelDoorDetailGateway sharedGateway.GetWaterChannelDoorByID, sendSSEGateway sharedGateway.SendSSEMessage) AIGetDebitFromWaterChannelUseCase {
	return func(ctx context.Context, req AIGetDebitFromWaterChannelReq) (*AIGetDebitFromWaterChannelRes, error) {

		// Get Debit
		debitData, err := getDebitGateway(ctx, sharedGateway.GetLatestDebitReq{
			WaterChannelDoorID: req.WaterChannelDoorID,
		})
		if err != nil {
			return nil, err
		}

		waterChannelDoorDetail, err := getWaterChannelDoorDetailGateway(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		result := &AIGetDebitFromWaterChannelRes{
			WaterChannelDoorID: debitData.Debit.WaterChannelDoorID,
			Name:               waterChannelDoorDetail.WaterChannelDoor.Name,
			Debit:              debitData.Debit.ActualDebit,
		}

		if _, err = sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "get-latest-debit-from-water-channel-by-id",
			FunctionName: "showLatestDebit",
			Data:         result,
		}); err != nil {
			return nil, err
		}

		return result, nil
	}
}
