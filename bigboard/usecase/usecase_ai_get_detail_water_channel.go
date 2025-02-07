package usecase

import (
	"context"
	"fmt"
	"shared/core"
	sharedGateway "shared/gateway"
	"shared/usecase"
)

type AiGetDetailWaterChannelReq struct {
	WaterChannelID int `json:"water_channel_id"`
}

type AiGetDetailWaterChannelResp struct {
	ExternalID            int               `json:"external_id"`
	Name                  string            `json:"name"`
	DebitActual           string            `json:"actual_debit"`      // get from WaterChannelDoor.DebitActual where WaterChannelDoor.ExternalID = x
	DebitRequirement      string            `json:"debit_requirement"` // get from WaterChannelDoor.DebitRequirement where WaterChannelDoor.ExternalID = x
	WaterSurfaceElevation string            `json:"water_surface_elevation"`
	Officers              []usecase.Officer `json:"officers"` // see `Officer` struct
	NCCTV                 int               `json:"number_of_cctv"`
	Latitude              string            `json:"latitude"`
	Longitude             string            `json:"longitude"`
	Status                string            `json:"status"`
}

type AiGetDetailWaterChannelUseCase = core.ActionHandler[AiGetDetailWaterChannelReq, AiGetDetailWaterChannelResp]

func ImplWaterChannelDetailUseCase(
	getWaterChannelDoorByID sharedGateway.GetWaterChannelDoorByID,
	getWaterChannelDevicesByDoorID sharedGateway.GetWaterChannelDevicesByDoorID,
	getWaterChannelOfficersByDoorID sharedGateway.GetWaterChannelOfficersByDoorID,
	getLatestDebitGateway sharedGateway.GetLatestDebit,
	getWaterSurfaceElevationGateway sharedGateway.GetLatestWaterSurfaceElevationByDoorID,
	sendSSEGateway sharedGateway.SendSSEMessage,
) AiGetDetailWaterChannelUseCase {
	return func(ctx context.Context, request AiGetDetailWaterChannelReq) (*AiGetDetailWaterChannelResp, error) {

		// Get WaterChannelDoor
		doorRes, err := getWaterChannelDoorByID(ctx, sharedGateway.GetWaterChannelDoorByIDReq{WaterChannelDoorID: request.WaterChannelID})
		if err != nil {
			return nil, err
		}
		door := doorRes.WaterChannelDoor

		// Get WaterChannelDevices
		devicesRes, err := getWaterChannelDevicesByDoorID(ctx, sharedGateway.GetWaterChannelDevicesByDoorIDReq{WaterChannelDoorID: door.ExternalID})
		if err != nil {
			return nil, err
		}

		// Get WaterChannelOfficers
		officersRes, err := getWaterChannelOfficersByDoorID(ctx, sharedGateway.GetWaterChannelOfficersByDoorIDReq{WaterChannelDoorID: door.ExternalID})
		if err != nil {
			return nil, err
		}

		debitRes, err := getLatestDebitGateway(ctx, sharedGateway.GetLatestDebitReq{WaterChannelDoorID: door.ExternalID})
		if err != nil {
			return nil, err
		}

		waterSurfaceElevationsRes, err := getWaterSurfaceElevationGateway(ctx, sharedGateway.GetLatestWaterSurfaceElevationByDoorIDReq{
			WaterChannelDoorID: door.ExternalID,
		})
		if err != nil {
			return nil, err
		}

		// Prepare response
		response := &AiGetDetailWaterChannelResp{
			ExternalID:            door.ExternalID,
			Name:                  door.Name,
			DebitActual:           fmt.Sprintf("%.2f liter per detik", debitRes.Debit.ActualDebit),
			DebitRequirement:      door.DebitRequirement + " liter per detik",
			WaterSurfaceElevation: fmt.Sprintf("%.2f cm", waterSurfaceElevationsRes.WaterSurfaceElevation.WaterLevel),
			Officers:              usecase.MapOfficersForAI(officersRes.Officers),
			NCCTV:                 len(usecase.MapCCTVs(devicesRes.Devices)),
			Latitude:              door.Latitude,
			Longitude:             door.Longitude,
			Status:                getWaterChannelDoorStatus(debitRes.Debit.ActualDebit, usecase.ParseFloat64(door.DebitRequirement)),
		}

		_, _ = sendSSEGateway(ctx, sharedGateway.SendSSEMessageReq{
			Subject:      "get-detail-from-water-channel",
			FunctionName: "showDetailWaterChannel",
			Data:         response,
		})

		return response, nil
	}
}
