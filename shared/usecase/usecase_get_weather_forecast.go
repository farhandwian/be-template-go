package usecase

import (
	"bigboard/model"
	"context"
	"fmt"
	"shared/core"
	"shared/gateway"
	"time"
)

type GetWeatherForecastReq struct {
	ADM4 string
}

type GetWeatherForecastRes struct {
	Province         string                      `json:"province"`
	City             string                      `json:"city"`
	District         string                      `json:"district"`
	SubDistrict      string                      `json:"sub_district"`
	WeatherForecasts []model.WeatherForecastItem `json:"weather_forecasts"`
}

type GetWeatherForecastUseCase = core.ActionHandler[GetWeatherForecastReq, GetWeatherForecastRes]

func ImplGetWeatherForecast(fetchBMKGWeatherForecast gateway.FetchBMKGWeatherForecast) GetWeatherForecastUseCase {
	return func(ctx context.Context, request GetWeatherForecastReq) (*GetWeatherForecastRes, error) {
		bmkgResp, err := fetchBMKGWeatherForecast(ctx, gateway.FetchBMKGWeatherForecastReq{ADM4: request.ADM4})
		if err != nil {
			return nil, err
		}

		weatherForecast := &GetWeatherForecastRes{}
		weatherForecast.Province = bmkgResp.BMKGResponse.BMKGLocation.Provinsi
		weatherForecast.City = bmkgResp.BMKGResponse.BMKGLocation.KotaKab
		weatherForecast.District = bmkgResp.BMKGResponse.BMKGLocation.Kecamatan
		weatherForecast.SubDistrict = bmkgResp.BMKGResponse.BMKGLocation.Desa

		for _, data := range bmkgResp.BMKGResponse.Data {
			for _, forecastGroup := range data.Cuaca {
				for _, forecast := range forecastGroup {
					localTime, err := time.Parse("2006-01-02 15:04:05", forecast.LocalDatetime)
					if err != nil {
						continue
					}

					item := model.WeatherForecastItem{
						Date:          localTime.Format("2006-01-02"),
						Time:          localTime.Format("15:04"),
						Status:        forecast.WeatherDesc,
						Temperature:   fmt.Sprintf("%.1f°C", forecast.T),
						Humidity:      fmt.Sprintf("%d%%", forecast.Hu),
						WindDirection: forecast.Wd,
						WindVelocity:  fmt.Sprintf("%.1f km/h", forecast.Ws),
						DamUpstream:   "", // This information is not available in the API response
					}

					weatherForecast.WeatherForecasts = append(weatherForecast.WeatherForecasts, item)
				}
			}
		}

		return weatherForecast, nil
	}
}
