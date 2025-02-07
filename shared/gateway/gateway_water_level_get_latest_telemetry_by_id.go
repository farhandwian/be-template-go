package gateway

import (
	"context"
	"gorm.io/gorm"
	"shared/core"
	"shared/middleware"
	"shared/model"
	"time"
)

type GetLatestTelemetryWaterPostByIDReq struct {
	WaterLevelPostID int64
}

type GetLatestTelemetryWaterPostByIDResp struct {
	LatestManualWaterPost model.HydrologyWaterLevelTelemetry
}

type GetLatestTelemetryWaterPostByIDGateway = core.ActionHandler[GetLatestTelemetryWaterPostByIDReq, GetLatestTelemetryWaterPostByIDResp]

func ImplGetLatestTelemetryWaterPostByIDGateway(db *gorm.DB) GetLatestTelemetryWaterPostByIDGateway {
	return func(ctx context.Context, req GetLatestTelemetryWaterPostByIDReq) (*GetLatestTelemetryWaterPostByIDResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		queryStr := `WITH RankedData AS (
			SELECT 
				water_level_post_id,
				water_level,
				sampling,
				timestamp,
				ROW_NUMBER() OVER (
					PARTITION BY water_level_post_id 
					ORDER BY timestamp DESC, created_at DESC
				) AS row_num
			FROM hydrology_water_level_telemetries
			WHERE water_level_post_id = ? AND 
				  TO_DATE(sampling, 'YYYY-MM-DD') = ?
		)
		SELECT 
			water_level_post_id,
			water_level,
			sampling,
			timestamp
		FROM RankedData
		WHERE row_num = 1;`

		var waterLevelTelemetry model.HydrologyWaterLevelTelemetry

		err := query.Raw(queryStr, req.WaterLevelPostID, time.Now().Format(time.DateOnly)).Scan(&waterLevelTelemetry).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetLatestTelemetryWaterPostByIDResp{
			LatestManualWaterPost: waterLevelTelemetry,
		}, nil
	}
}
