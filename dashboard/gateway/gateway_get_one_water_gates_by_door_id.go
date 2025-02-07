package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetOneWaterGatesByDoorIDReq struct {
	WaterChannelDoorID int
	DeviceID           int
}

type GetOneWaterGatesByDoorIDRes struct {
	WaterGates *model.WaterGateData
}

type GetOneWaterGatesByDoorID = core.ActionHandler[GetOneWaterGatesByDoorIDReq, GetOneWaterGatesByDoorIDRes]

func ImplGetOneWaterGatesByDoorID(db *gorm.DB) GetOneWaterGatesByDoorID {
	return func(ctx context.Context, req GetOneWaterGatesByDoorIDReq) (*GetOneWaterGatesByDoorIDRes, error) {
		var waterGates model.WaterGateData

		err := db.
			Order("timestamp DESC").
			First(&waterGates, "water_channel_door_id = ? AND device_id = ?", req.WaterChannelDoorID, req.DeviceID).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, core.NewInternalServerError(err)
		}

		return &GetOneWaterGatesByDoorIDRes{
			WaterGates: &waterGates,
		}, nil
	}
}
