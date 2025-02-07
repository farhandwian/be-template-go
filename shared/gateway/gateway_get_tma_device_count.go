package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetTmaDeviceCountReq struct{}

type GetTmaDeviceCountRes struct {
	TMADeviceCount model.TmaDevice
}

type GetTmaDeviceCountGateway = core.ActionHandler[GetTmaDeviceCountReq, GetTmaDeviceCountRes]

func ImplGetTmaDeviceCount(tsDB *gorm.DB) GetTmaDeviceCountGateway {
	return func(ctx context.Context, request GetTmaDeviceCountReq) (*GetTmaDeviceCountRes, error) {

		var tmaDeviceCount model.TmaDevice

		err := tsDB.Raw(`
		SELECT 
			COUNT(*) FILTER (WHERE latest_status = false) AS device_off,
			COUNT(*) FILTER (WHERE latest_status = true) AS device_on
		FROM latest_water_levels;
		`).Scan(&tmaDeviceCount).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetTmaDeviceCountRes{
			TMADeviceCount: model.TmaDevice{
				DeviceOff: tmaDeviceCount.DeviceOff,
				DeviceOn:  tmaDeviceCount.DeviceOn,
			},
		}, nil
	}
}
