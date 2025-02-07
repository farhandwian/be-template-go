package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"
	"time"

	"gorm.io/gorm"
)

type GetWaterLevelDetailTelemetryReq struct {
	ID int64
}

type GetWaterLevelDetailTelemetryRes struct {
	WaterLevelDetailTelemetry model.WaterLevelCalculation
}

type GetWaterLevelDetailCalculationGateway = core.ActionHandler[GetWaterLevelDetailTelemetryReq, GetWaterLevelDetailTelemetryRes]

func ImplGetWaterLevelPostCalculationDetailGateway(db *gorm.DB) GetWaterLevelDetailCalculationGateway {
	return func(ctx context.Context, req GetWaterLevelDetailTelemetryReq) (*GetWaterLevelDetailTelemetryRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		date := time.Now().Format(time.DateOnly)

		var telemetry model.WaterLevelCalculation

		err := query.Raw(`
			SELECT 
                        hwlt.water_level_post_id as id,
                        wlp.name,
                        AVG(hwlt.water_level) as water_level_telemetry,
                        (
                                SELECT tma
                                FROM hydrology_water_level_manuals mwl
                                WHERE mwl.water_level_post_id = hwlt.water_level_post_id
                                AND mwl.sampling != ''  
                                AND DATE(CAST(NULLIF(mwl.sampling, '') AS timestamp)) = ?
                                ORDER BY sampling DESC, created_at DESC
                                LIMIT 1
                        ) as water_level_manual,
                        MAX(hwlt.sampling) as latest_update 
                FROM 
                        hydrology_water_level_telemetries hwlt
                        LEFT JOIN hydrology_water_level_posts wlp ON hwlt.water_level_post_id = wlp.water_level_post_id
                WHERE 
                        hwlt.sampling != ''  
                        AND DATE(CAST(NULLIF(hwlt.sampling, '') AS timestamp)) = ? and hwlt.water_level_post_id =?
                GROUP BY 
                        hwlt.water_level_post_id,
                        wlp.name,
                        wlp.river,
                        wlp.elevation
                ORDER BY 
                        hwlt.water_level_post_id;       
        `, date, date, req.ID).Scan(&telemetry).Error

		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error fetching latest water level telemetry: %v", err))
		}

		return &GetWaterLevelDetailTelemetryRes{
			WaterLevelDetailTelemetry: telemetry,
		}, nil
	}
}
