package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetCCTVDevicesReq struct {
}

type GetCCTVDevicesRes struct {
	Devices []model.WaterChannelDevice
}

type GetCCTVDevices = core.ActionHandler[GetCCTVDevicesReq, GetCCTVDevicesRes]

func ImplGetCCTVDevices(db *gorm.DB) GetCCTVDevices {
	return func(ctx context.Context, request GetCCTVDevicesReq) (*GetCCTVDevicesRes, error) {
		var devices []model.WaterChannelDevice

		result := db.
			Where("category = ?", "cctv").
			Find(&devices)
		if result.Error != nil {
			return nil, core.NewInternalServerError(result.Error)
		}

		return &GetCCTVDevicesRes{
			Devices: devices,
		}, nil
	}
}
