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

type GetWaterLevelDetailManualReq struct {
	ID        int64
	StartDate string
	EndDate   string
}

type GetWaterLevelDetailManualRes struct {
	WaterLevelDetailTelemetry []model.HydrologyWaterLevelManual
}

type GetWaterLevelDetailManualGateway = core.ActionHandler[GetWaterLevelDetailManualReq, GetWaterLevelDetailManualRes]

func ImplGetWaterLevelDetailManualGateway(db *gorm.DB) GetWaterLevelDetailManualGateway {
	return func(ctx context.Context, req GetWaterLevelDetailManualReq) (*GetWaterLevelDetailManualRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var detail []model.HydrologyWaterLevelManual

		// Parse dates
		startDate, err := time.Parse(time.DateOnly, req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start date format: %v", err)
		}

		endDate, err := time.Parse(time.DateOnly, req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end date format: %v", err)
		}

		startDateTime := startDate.Format("2006-01-02 00:00:00")
		endDateTime := endDate.Format("2006-01-02 23:59:59.999")

		// Execute query with formatted timestamps
		err = query.Table("hydrology_water_level_manuals").
			Select("*").
			Where("water_level_post_id = ?", req.ID).
			Where("sampling BETWEEN ? AND ?", startDateTime, endDateTime).
			Order("sampling ASC").
			Find(&detail).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetWaterLevelDetailManualRes{
			WaterLevelDetailTelemetry: detail,
		}, nil
	}
}
