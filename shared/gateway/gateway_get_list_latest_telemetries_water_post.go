package gateway

import (
	"context"
	"gorm.io/gorm"
	"shared/core"
	"shared/middleware"
	"shared/model"
)

type GetListLatestTelemetryWaterPostReq struct {
	Date string
}

type GetListLatestTelemetryWaterPostResp struct {
	LatestManualWaterPost []model.HydrologyWaterLevelTelemetry
}

type GetListLatestTelemetryWaterPostGateway = core.ActionHandler[GetListLatestTelemetryWaterPostReq, GetListLatestTelemetryWaterPostResp]

func ImplGetListLatestTelemetryWaterPostGateway(db *gorm.DB) GetListLatestTelemetryWaterPostGateway {
	return func(ctx context.Context, req GetListLatestTelemetryWaterPostReq) (*GetListLatestTelemetryWaterPostResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		queryStr := `WITH RankedData AS (
			SELECT 
				water_level_post_id,
				water_level,
				sampling,
				timestamp,
				created_at,
				ROW_NUMBER() OVER (
					PARTITION BY water_level_post_id 
					ORDER BY timestamp DESC, created_at DESC
				) AS row_num
			FROM hydrology_water_level_telemetries
			WHERE DATE(created_at) = ? 
			)
			SELECT 
				water_level_post_id,
				water_level,
				sampling,
				timestamp,
				created_at
			FROM RankedData
			WHERE row_num = 1;`

		var waterLevelTelemetry []model.HydrologyWaterLevelTelemetry

		err := query.Debug().Raw(queryStr, req.Date).Scan(&waterLevelTelemetry).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListLatestTelemetryWaterPostResp{
			LatestManualWaterPost: waterLevelTelemetry,
		}, nil
	}
}
