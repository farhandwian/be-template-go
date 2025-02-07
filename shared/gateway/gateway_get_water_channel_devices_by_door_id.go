package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetWaterChannelDevicesByDoorIDReq struct {
	WaterChannelDoorID int
}

type GetWaterChannelDevicesByDoorIDRes struct {
	Devices []model.WaterChannelDevice
}

type GetWaterChannelDevicesByDoorID = core.ActionHandler[GetWaterChannelDevicesByDoorIDReq, GetWaterChannelDevicesByDoorIDRes]

func ImplGetWaterChannelDevicesByDoorID(db *gorm.DB) GetWaterChannelDevicesByDoorID {
	return func(ctx context.Context, request GetWaterChannelDevicesByDoorIDReq) (*GetWaterChannelDevicesByDoorIDRes, error) {
		var devices []model.WaterChannelDevice

		result := db.
			Where("water_channel_door_id = ?", request.WaterChannelDoorID).
			Find(&devices)
		if result.Error != nil {
			return nil, core.NewInternalServerError(result.Error)
		}

		return &GetWaterChannelDevicesByDoorIDRes{
			Devices: devices,
		}, nil
	}
}
