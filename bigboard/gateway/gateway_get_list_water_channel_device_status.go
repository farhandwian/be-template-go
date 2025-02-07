package gateway

import (
	"bigboard/model"
	"context"
	"shared/core"

	"gorm.io/gorm"
)

type GetListWaterChannelDeviceStatusReq struct{}

type GetListWaterChannelDeviceStatusRes struct {
	Devices []model.WaterChannelDoorDevice `json:"devices"`
}

type GetListWaterChannelDeviceGateway = core.ActionHandler[GetListWaterChannelDeviceStatusReq, GetListWaterChannelDeviceStatusRes]

func ImplGetListWaterChannelDeviceGateway(tsDB *gorm.DB) GetListWaterChannelDeviceGateway {
	return func(ctx context.Context, request GetListWaterChannelDeviceStatusReq) (*GetListWaterChannelDeviceStatusRes, error) {

		var waterChannelDevices []model.WaterChannelDoorDevice

		err := tsDB.Raw(`
			SELECT DISTINCT ON (water_channel_door_id) 
			water_channel_door_id,
			actual_debit,
			timestamp AS latest_timestamp
		FROM actual_debits
		ORDER BY water_channel_door_id, timestamp DESC;
	`).Find(&waterChannelDevices).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListWaterChannelDeviceStatusRes{
			Devices: waterChannelDevices,
		}, nil
	}
}
