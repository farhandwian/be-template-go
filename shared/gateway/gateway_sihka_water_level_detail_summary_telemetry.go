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

type GetWaterLevelDetailSummaryTelemetryReq struct {
	ID        int64
	StartDate string
	EndDate   string
}

type GetWaterLevelDetailSummaryTelemetryRes struct {
	WaterLevelDetailTelemetry model.HydrologyWaterLevelSummaryTelemetry
}

type GetWaterLevelDetailSummaryTelemetryGateway = core.ActionHandler[GetWaterLevelDetailSummaryTelemetryReq, GetWaterLevelDetailSummaryTelemetryRes]

func ImplGetWaterLevelDetailSummaryTelemetryGateway(db *gorm.DB) GetWaterLevelDetailSummaryTelemetryGateway {
	return func(ctx context.Context, req GetWaterLevelDetailSummaryTelemetryReq) (*GetWaterLevelDetailSummaryTelemetryRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var detail model.HydrologyWaterLevelSummaryTelemetry

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

		err = query.Table("hydrology_water_level_telemetries").
			Select("MAX(water_level) as tma_maximum, MIN(water_level) as tma_minimum, max(sampling) as latest_update,water_level_post_id").
			Where("water_level_post_id = ?", req.ID).
			Where("timestamp BETWEEN ? AND ?", startDateTime, endDateTime).
			Group("water_level_post_id").
			Order("water_level_post_id").
			Limit(1).
			Find(&detail).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetWaterLevelDetailSummaryTelemetryRes{
			WaterLevelDetailTelemetry: detail,
		}, nil
	}
}
