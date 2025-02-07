package gateway

import (
	"bigboard/model"
	"context"
	"gorm.io/gorm"
	"shared/core"
	"time"
)

type GetWaterAvailabilityForecastReq struct {
}

type GetWaterAvailabilityForecastResp struct {
	WaterAvailabilityForecast []model.WaterAvailabilityForecast
}

type GetWaterAvailabilityForecastGateway = core.ActionHandler[GetWaterAvailabilityForecastReq, GetWaterAvailabilityForecastResp]

func ImplGetWaterAvailabilityForecastGateway(db *gorm.DB) GetWaterAvailabilityForecastGateway {
	return func(ctx context.Context, request GetWaterAvailabilityForecastReq) (*GetWaterAvailabilityForecastResp, error) {
		var waterAvailabilityForecast []model.WaterAvailabilityForecast
		currentDate := time.Now().Format(time.DateOnly)
		err := db.Where("date >= ?", currentDate).Order("date ASC").Limit(15).Find(&waterAvailabilityForecast).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetWaterAvailabilityForecastResp{
			WaterAvailabilityForecast: waterAvailabilityForecast,
		}, nil
	}
}
