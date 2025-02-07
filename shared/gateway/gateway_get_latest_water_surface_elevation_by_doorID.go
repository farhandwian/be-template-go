package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetLatestWaterSurfaceElevationByDoorIDReq struct {
	WaterChannelDoorID int
}

type GetLatestWaterSurfaceElevationByDoorIDRes struct {
	WaterSurfaceElevation *model.WaterSurfaceElevationData
}

type GetLatestWaterSurfaceElevationByDoorID = core.ActionHandler[GetLatestWaterSurfaceElevationByDoorIDReq, GetLatestWaterSurfaceElevationByDoorIDRes]

func ImplGetLatestWaterSurfaceElevationByDoorID(db *gorm.DB) GetLatestWaterSurfaceElevationByDoorID {
	return func(ctx context.Context, request GetLatestWaterSurfaceElevationByDoorIDReq) (*GetLatestWaterSurfaceElevationByDoorIDRes, error) {
		var waterSurfaceElevation model.WaterSurfaceElevationData

		result := db.Raw(`
			SELECT timestamp, water_channel_door_id, latest_water_levels.latest_level as water_level, latest_water_levels.latest_status as status
			FROM latest_water_levels
			WHERE water_channel_door_id = ?
			LIMIT 1
		`, request.WaterChannelDoorID).Scan(&waterSurfaceElevation)

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return &GetLatestWaterSurfaceElevationByDoorIDRes{
					WaterSurfaceElevation: nil,
				}, nil
			}
			return nil, core.NewInternalServerError(result.Error)
		}

		return &GetLatestWaterSurfaceElevationByDoorIDRes{
			WaterSurfaceElevation: &waterSurfaceElevation,
		}, nil
	}
}
