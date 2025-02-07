package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type WaterLevelPostCalculationReq struct {
	Date string
}
type WaterLevelPostCalculationResp struct {
	WaterLevelPost []model.WaterLevelCalculation
}

type WaterLevelPostCalculationGateway = core.ActionHandler[WaterLevelPostCalculationReq, WaterLevelPostCalculationResp]

func ImplWaterLevelPostCalculationGateway(db *gorm.DB) WaterLevelPostCalculationGateway {
	return func(ctx context.Context, request WaterLevelPostCalculationReq) (*WaterLevelPostCalculationResp, error) {

		query := middleware.GetDBFromContext(ctx, db)

		queryStr := `
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
			AND DATE(CAST(NULLIF(hwlt.sampling, '') AS timestamp)) = ?
		GROUP BY 
			hwlt.water_level_post_id,
			wlp.name,
			wlp.river,
			wlp.elevation
		ORDER BY 
			hwlt.water_level_post_id;		
		`

		var waterLevelPostList []model.WaterLevelCalculation

		err := query.Raw(queryStr, request.Date, request.Date).Scan(&waterLevelPostList).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &WaterLevelPostCalculationResp{
			WaterLevelPost: waterLevelPostList,
		}, err
	}
}
