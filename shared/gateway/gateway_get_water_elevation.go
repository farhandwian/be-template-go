package gateway

import (
	"context"
	"shared/core"
	"time"

	"gorm.io/gorm"
)

type GetWaterSurfaceElevationReq struct {
	WaterChannelDoorID int
	MinTime            time.Time `json:"min_time"`
	MaxTime            time.Time `json:"max_time"`
}

type WaterSurfaceElevation struct {
	WaterLevel float64   `json:"tma"`
	Time       time.Time `json:"time"`
}

type GetWaterSurfaceElevationRes struct {
	Results []WaterSurfaceElevation
}

type GetWaterSurfaceElevation = core.ActionHandler[GetWaterSurfaceElevationReq, GetWaterSurfaceElevationRes]

func ImplGetWaterSurfaceElevation(tsDB *gorm.DB) GetWaterSurfaceElevation {
	return func(ctx context.Context, req GetWaterSurfaceElevationReq) (*GetWaterSurfaceElevationRes, error) {

		var results []WaterSurfaceElevation
		timeRange := req.MaxTime.Sub(req.MinTime).Hours() / 24

		bucket := "5 minutes"
		if timeRange > 7 {
			bucket = "1 day"
		} else if timeRange > 1 {
			bucket = "1 hour"
		}

		query := `
			SELECT
			time_bucket(?, timestamp) AS time,
			AVG(water_level) AS water_level
			FROM water_surface_elevation_data
			WHERE water_channel_door_id = ? AND status = true
			AND timestamp BETWEEN ? AND ?
			GROUP BY time ORDER BY time ASC;
		`
		err := tsDB.Raw(query, bucket, req.WaterChannelDoorID, req.MinTime, req.MaxTime).Scan(&results).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetWaterSurfaceElevationRes{
			Results: results,
		}, nil
	}
}
