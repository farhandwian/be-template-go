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
	ID        int64
	StartDate string
	EndDate   string
}

type GetWaterLevelDetailTelemetryRes struct {
	WaterLevelDetailTelemetry []model.HydrologyWaterLevelTelemetry
}

type GetWaterLevelDetailTelemetryGateway = core.ActionHandler[GetWaterLevelDetailTelemetryReq, GetWaterLevelDetailTelemetryRes]

func ImplGetWaterLevelDetailTelemetryGateway(db *gorm.DB) GetWaterLevelDetailTelemetryGateway {
	return func(ctx context.Context, req GetWaterLevelDetailTelemetryReq) (*GetWaterLevelDetailTelemetryRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var telemetries []model.HydrologyWaterLevelTelemetry

		startDate, err := time.Parse(time.DateOnly, req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start date format: %v", err)
		}

		endDate, err := time.Parse(time.DateOnly, req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end date format: %v", err)
		}

		startDateTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
		endDateTime := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, endDate.Location())

		// Perform hourly grouping based on `Timestamp`
		err = query.Table("hydrology_water_level_telemetries").
			Select(`
        DATE_TRUNC('hour', sampling::timestamp) AS sampling, 
        AVG(battery) AS battery, 
		water_level_post_id,
        CASE 
            WHEN AVG(battery) != 0 THEN AVG(water_level) * 100 
            ELSE AVG(water_level) 
        END AS water_level
    `).
			Where("water_level_post_id = ?", req.ID).
			Where("sampling::timestamp BETWEEN ? AND ?", startDateTime, endDateTime).
			Group("DATE_TRUNC('hour', sampling::timestamp), water_level_post_id"). // Explicit grouping by truncated timestamp
			Order("sampling ASC").
			Find(&telemetries).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetWaterLevelDetailTelemetryRes{
			WaterLevelDetailTelemetry: telemetries,
		}, nil
	}
}
