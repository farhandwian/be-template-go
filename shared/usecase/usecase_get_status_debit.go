package usecase

import (
	"context"
	"shared/core"
	gateway2 "shared/gateway"
)

type GetStatusDebitReq struct {
	WaterChannelDoorID int `json:"water_channel_door_id"`
}

type GetStatusDebitRes struct {
	DebitActual      float64 `json:"actual_debit"`      // get from WaterChannelDoor.DebitActual where WaterChannelDoor.ExternalID = x
	DebitRequirement float64 `json:"debit_requirement"` // get from WaterChannelDoor.DebitRequirement where WaterChannelDoor.ExternalID = x
	WaterGateStatus  bool    `json:"water_gate_status"` // get from WaterSurfaceElevation.Status where WaterSurfaceElevation.WaterChannelDoorID = x wuth the latest WaterSurfaceElevation.CreatedAt
	WaterLevel       float32 `json:"water_level"`       // get from WaterSurfaceElevation.WaterLevel where WaterSurfaceElevation.WaterChannelDoorID = x wuth the latest WaterSurfaceElevation.CreatedAt
	WaterLevelStatus bool    `json:"water_level_status"`
	DebitPer1cm      float32 `json:"debit_per_1cm"` // return -1 for now

}

type GetStatusDebitUseCase = core.ActionHandler[GetStatusDebitReq, GetStatusDebitRes]

func ImplGetStatusDebit(
	getWaterChannelDoorByID gateway2.GetWaterChannelDoorByID,
	getTMA gateway2.GetLatestWaterSurfaceElevationByDoorID,
	getDebit gateway2.GetLatestDebit,

) GetStatusDebitUseCase {
	return func(ctx context.Context, req GetStatusDebitReq) (*GetStatusDebitRes, error) {

		// Get WaterChannelDoor
		doorRes, err := getWaterChannelDoorByID(ctx, gateway2.GetWaterChannelDoorByIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}
		door := doorRes.WaterChannelDoor

		// Get latest WaterSurfaceElevation
		waterSurfaceElevationRes, err := getTMA(ctx, gateway2.GetLatestWaterSurfaceElevationByDoorIDReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		debitObj, err := getDebit(ctx, gateway2.GetLatestDebitReq{WaterChannelDoorID: req.WaterChannelDoorID})
		if err != nil {
			return nil, err
		}

		// Prepare response
		response := &GetStatusDebitRes{
			DebitActual:      debitObj.Debit.ActualDebit,
			DebitRequirement: ParseFloat64(door.DebitRequirement),
			DebitPer1cm:      0, // TODO where to get this value
			WaterLevelStatus: waterSurfaceElevationRes.WaterSurfaceElevation.Status,
		}

		if waterSurfaceElevationRes.WaterSurfaceElevation != nil {
			response.WaterLevel = waterSurfaceElevationRes.WaterSurfaceElevation.WaterLevel
		}

		response.WaterGateStatus = getWaterChannelDoorStatusBool(response.DebitActual, response.DebitRequirement)

		return response, nil
	}
}

func getWaterChannelDoorStatusBool(actualDebit, requiredDebit float64) bool {
	if actualDebit < requiredDebit {
		return false
	} else {
		return true
	}
}
