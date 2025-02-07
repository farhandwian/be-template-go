package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type GetRainfallHourlyReq struct {
	PostID    int64
	StartDate string
	EndDate   string
}

type GetRainfallHourlyResp struct {
	RainfallHourly []model.HydrologyRainHourly
}

type GetRainfallHourlyGateway = core.ActionHandler[GetRainfallHourlyReq, GetRainfallHourlyResp]

func ImplGetRainfallHourlyGateway(db *gorm.DB) GetRainfallHourlyGateway {
	return func(ctx context.Context, request GetRainfallHourlyReq) (*GetRainfallHourlyResp, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var RainfallHourly []model.HydrologyRainHourly

		err := query.Raw(`
			SELECT 
				((sampling::date + (hour || ' hour')::interval) AT TIME ZONE '-7') AS sampling,
				rain_post_id,
				count,
				rain,
				hour
			FROM 
				hydrology_rain_hourlies
			WHERE 
				rain_post_id = ?
				AND sampling >= ?
				AND sampling <= ?
			ORDER BY 
				updated_at ASC;
		`, request.PostID, request.StartDate, request.EndDate).
			Scan(&RainfallHourly).Error

		if err != nil {
			return nil, core.NewInternalServerError(fmt.Errorf("error fetching rainfall post hourly list: %w", err))
		}

		return &GetRainfallHourlyResp{
			RainfallHourly: RainfallHourly,
		}, nil
	}
}
