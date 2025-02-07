package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetListWaterChannelTMAReq struct{}

type GetListWaterChannelTMARes struct {
	WaterChannelTMA []model.WaterSurfaceElevationList
}

type GetListWaterChannelTMAGateway = core.ActionHandler[GetListWaterChannelTMAReq, GetListWaterChannelTMARes]

func ImplGetListWaterChannelTMA(tsDB *gorm.DB) GetListWaterChannelTMAGateway {
	return func(ctx context.Context, request GetListWaterChannelTMAReq) (*GetListWaterChannelTMARes, error) {

		var tma []model.WaterSurfaceElevationList

		err := tsDB.Raw(`
			SELECT timestamp,
				water_channel_door_id,
				latest_water_levels.latest_level as water_surface_elevation,
				latest_status as status,
				water_channel_door_id
			FROM latest_water_levels;
		`).Scan(&tma).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListWaterChannelTMARes{
			WaterChannelTMA: tma,
		}, nil
	}
}
