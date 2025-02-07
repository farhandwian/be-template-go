package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetDeviceReq struct {
	WaterChannelDoorID int
}

type GetDeviceRes struct {
	Devices []model.WaterChannelDevice
}

type GetDevice = core.ActionHandler[GetDeviceReq, GetDeviceRes]

func ImplGetDevice(db *gorm.DB) GetDevice {
	return func(ctx context.Context, request GetDeviceReq) (*GetDeviceRes, error) {
		var devices []model.WaterChannelDevice

		result := db.
			Where("water_channel_door_id = ? AND category = ?", request.WaterChannelDoorID, "controller").
			Find(&devices)
		if result.Error != nil {
			return nil, core.NewInternalServerError(result.Error)
		}

		return &GetDeviceRes{
			Devices: devices,
		}, nil
	}
}
