package gateway

import (
	"context"
	"gorm.io/gorm"
	"shared/core"
	"shared/middleware"
	"shared/model"
)

type GetListLatestManualWaterPostReq struct {
	Date string
}

type GetListLatestManualWaterPostResp struct {
	LatestManualWaterPost []model.HydrologyWaterLevelManual
}

type GetListLatestManualWaterPostGateway = core.ActionHandler[GetListLatestManualWaterPostReq, GetListLatestManualWaterPostResp]

func ImplGetListLatestManualWaterPostHandler(db *gorm.DB) GetListLatestManualWaterPostGateway {
	return func(ctx context.Context, req GetListLatestManualWaterPostReq) (*GetListLatestManualWaterPostResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		queryStr := `WITH RankedData AS (
			SELECT 
				water_level_post_id,
				tma,
				sampling,
				timestamp,
				ROW_NUMBER() OVER (
					PARTITION BY water_level_post_id 
					ORDER BY timestamp DESC, created_at DESC
				) AS row_num
			FROM hydrology_water_level_manuals
			WHERE DATE(sampling) = ?
		)
		SELECT 
			water_level_post_id,
			tma,
			sampling,
			timestamp
		FROM RankedData
		WHERE row_num = 1;`

		var waterLevelManual []model.HydrologyWaterLevelManual

		err := query.Raw(queryStr, req.Date).Scan(&waterLevelManual).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListLatestManualWaterPostResp{
			LatestManualWaterPost: waterLevelManual,
		}, nil
	}
}
