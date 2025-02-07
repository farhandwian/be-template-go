package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"

	"gorm.io/gorm"
)

type GetListWaterChannelDeviceStatusReq struct {
	WaterChannelDoorIDs []int // New field to hold the list of IDs
}

type GetListWaterChannelDeviceStatusRes struct {
	Devices []model.WaterChannelDoorDevice `json:"devices"`
}

type GetListWaterChannelDeviceGateway = core.ActionHandler[GetListWaterChannelDeviceStatusReq, GetListWaterChannelDeviceStatusRes]

func ImplGetListWaterChannelDeviceGateway(tsDB *gorm.DB) GetListWaterChannelDeviceGateway {
	return func(ctx context.Context, request GetListWaterChannelDeviceStatusReq) (*GetListWaterChannelDeviceStatusRes, error) {
		var waterChannelDevices []model.WaterChannelDoorDevice

		query := tsDB.Table("water_gates").
			Select("DISTINCT ON (water_channel_door_id, device_id) *").
			Order("water_channel_door_id, device_id, updated_at DESC")

		// Add WHERE clause if WaterChannelDoorIDs are provided
		if len(request.WaterChannelDoorIDs) > 0 {
			query = query.Where("water_channel_door_id IN ?", request.WaterChannelDoorIDs)
		}

		err := query.Find(&waterChannelDevices).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListWaterChannelDeviceStatusRes{
			Devices: waterChannelDevices,
		}, nil
	}
}
