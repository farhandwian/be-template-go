package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type DeviceByWaterChannelDoorIdReq struct {
	WaterChannelDoorId int
}

type WaterChannelDevice struct {
	ExternalID int    `json:"door_id"`
	Name       string `json:"name"`
}

type DeviceByWaterChannelDoorIdRes struct {
	Items []WaterChannelDevice
}

type DeviceByWaterChannelDoorId = core.ActionHandler[DeviceByWaterChannelDoorIdReq, DeviceByWaterChannelDoorIdRes]

func ImplDeviceByWaterChannelDoorId(db *gorm.DB) DeviceByWaterChannelDoorId {
	return func(ctx context.Context, req DeviceByWaterChannelDoorIdReq) (*DeviceByWaterChannelDoorIdRes, error) {

		var waterChannelDevices []WaterChannelDevice

		query := middleware.GetDBFromContext(ctx, db)

		if err := query.
			Where("category = ?", "controller").
			Where("water_channel_door_id = ?", req.WaterChannelDoorId).
			Find(&waterChannelDevices).
			Error; err != nil {

			if err == gorm.ErrRecordNotFound {
				return &DeviceByWaterChannelDoorIdRes{Items: []WaterChannelDevice{}}, nil
			}

			return nil, core.NewInternalServerError(err)
		}

		return &DeviceByWaterChannelDoorIdRes{Items: waterChannelDevices}, nil
	}
}
