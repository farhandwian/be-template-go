package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"time"
)

type WaterAvailabilityForecastReq struct{}

type WaterAvailabilityForecastResp struct {
	WaterAvailabilityForecast []WaterAvailabilityItem `json:"water_availability_forecast"`
}

type WaterAvailabilityItem struct {
	ID                    string    `json:"id"`
	RequiredWater         float64   `json:"required_water"`
	AverageAvailableWater float64   `json:"average_available_water"`
	MaxAvailableWater     float64   `json:"max_available_water"`
	MinAvailableWater     float64   `json:"min_available_water"`
	Date                  string    `json:"date"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type WaterAvailabilityForecastUseCase = core.ActionHandler[WaterAvailabilityForecastReq, WaterAvailabilityForecastResp]

func ImplWaterAvailabilityForecastUseCase(waterAvailabilityForecastGateway gateway.GetWaterAvailabilityForecastGateway) WaterAvailabilityForecastUseCase {
	return func(ctx context.Context, request WaterAvailabilityForecastReq) (*WaterAvailabilityForecastResp, error) {

		waterAvailabilityForecastData, err := waterAvailabilityForecastGateway(ctx, gateway.GetWaterAvailabilityForecastReq{})
		if err != nil {
			return nil, err
		}

		var waterAvailabilityForecast []WaterAvailabilityItem
		for _, waterAvailabilityItem := range waterAvailabilityForecastData.WaterAvailabilityForecast {
			waterAvailabilityForecast = append(waterAvailabilityForecast, WaterAvailabilityItem{
				ID:                    waterAvailabilityItem.ID,
				RequiredWater:         waterAvailabilityItem.RequiredWater,
				AverageAvailableWater: waterAvailabilityItem.AverageAvailableWater,
				MaxAvailableWater:     waterAvailabilityItem.MaxAvailableWater,
				MinAvailableWater:     waterAvailabilityItem.MinAvailableWater,
				Date:                  waterAvailabilityItem.Date,
				CreatedAt:             waterAvailabilityItem.CreatedAt,
				UpdatedAt:             waterAvailabilityItem.UpdatedAt,
			})
		}
		return &WaterAvailabilityForecastResp{
			WaterAvailabilityForecast: waterAvailabilityForecast,
		}, nil
	}
}
