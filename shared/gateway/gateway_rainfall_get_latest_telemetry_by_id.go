package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetLatestTelemetryRainPostByIDReq struct {
	PostIDs []int64
	Date    string
}

type GetLatestTelemetryRainPostByIDResp struct {
	LatestTelemetry []model.HydrologyRainHourly
}

type GetLatestTelemetryRainPostByIDGateway = core.ActionHandler[GetLatestTelemetryRainPostByIDReq, GetLatestTelemetryRainPostByIDResp]

func ImplGetLatestTelemetryRainPostByIDGateway(db *gorm.DB) GetLatestTelemetryRainPostByIDGateway {
	return func(ctx context.Context, req GetLatestTelemetryRainPostByIDReq) (*GetLatestTelemetryRainPostByIDResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var rainfallsTelemetry []model.HydrologyRainHourly

		result := query.Raw(`
				SELECT DISTINCT ON (rain_post_id) *
				FROM hydrology_rain_hourlies
				WHERE 
					rain_post_id IN (?) 
					AND sampling = ?
					AND count > 0 
				ORDER BY rain_post_id, hour DESC;
			`, req.PostIDs, req.Date).
			Scan(&rainfallsTelemetry)

		if result.Error != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error fetching rainfall telemetry: %w", result.Error))
		}

		return &GetLatestTelemetryRainPostByIDResp{
			LatestTelemetry: rainfallsTelemetry,
		}, nil
	}
}
