package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetOneWaterChannelDeviceByDoorIDReq struct {
	WaterChannelDoorID int
	DeviceID           int
}

type GetOneWaterChannelDeviceByDoorIDRes struct {
	WaterChannelDevice *model.WaterChannelDevice
}

type GetOneWaterChannelDeviceByDoorID = core.ActionHandler[GetOneWaterChannelDeviceByDoorIDReq, GetOneWaterChannelDeviceByDoorIDRes]

func ImplGetOneWaterChannelDeviceByDoorID(db *gorm.DB) GetOneWaterChannelDeviceByDoorID {
	return func(ctx context.Context, req GetOneWaterChannelDeviceByDoorIDReq) (*GetOneWaterChannelDeviceByDoorIDRes, error) {
		var waterChannelDevice model.WaterChannelDevice

		err := db.First(&waterChannelDevice, "water_channel_door_id = ? AND external_id = ? AND category = ?", req.WaterChannelDoorID, req.DeviceID, "controller").Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, core.NewInternalServerError(err)
		}

		return &GetOneWaterChannelDeviceByDoorIDRes{
			WaterChannelDevice: &waterChannelDevice,
		}, nil
	}
}
