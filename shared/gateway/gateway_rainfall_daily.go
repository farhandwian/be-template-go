package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetRainfallDailyReq struct {
	PostIDs []int64
	Date    string
}

type GetRainfallDailyResp struct {
	RainfallDaily []model.HydrologyRainDaily
}

type GetRainfallDailyByPostID = core.ActionHandler[GetRainfallDailyReq, GetRainfallDailyResp]

func ImplGetRainfallDailyByPostIDGateway(db *gorm.DB) GetRainfallDailyByPostID {
	return func(ctx context.Context, request GetRainfallDailyReq) (*GetRainfallDailyResp, error) {
		query := middleware.GetDBFromContext(ctx, db).Session(&gorm.Session{
			PrepareStmt: true, // Enable prepared statements for better performance
		})

		var rainfallsTelemetry []model.HydrologyRainDaily

		// reviewed
		result := query.Raw(`
		WITH LatestUpdates AS (
			SELECT 
				rain_post_id, 
				MAX(updated_at) AS latest_updated_at
			FROM 
				hydrology_rain_dailies
			WHERE 
				rain_post_id IN (?) 
				AND sampling = ?
			GROUP BY 
				rain_post_id
		)
		SELECT 
			d.rain_post_id, 
			d.sampling, 
			d.count, 
			d.rain, 
			d.manual, 
			d.updated_at
		FROM 
			hydrology_rain_dailies d
		JOIN 
			LatestUpdates l
		ON 
			d.rain_post_id = l.rain_post_id 
			AND d.updated_at = l.latest_updated_at;
        `, request.PostIDs, request.Date).
			Scan(&rainfallsTelemetry)

		if result.Error != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error fetching rainfall telemetry: %w", result.Error))
		}

		return &GetRainfallDailyResp{
			RainfallDaily: rainfallsTelemetry,
		}, nil
	}
}
