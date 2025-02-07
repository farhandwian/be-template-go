package usecase

import (
	"context"
	"shared/core"
	"shared/usecase"
	"time"
)

type AISystemSummaryReq struct{}

type AISystemSummaryRes struct {
	GeneralInfo     usecase.GeneralInfoRes          `json:"general_info"`
	DeviceStatus    usecase.GetDeviceStatusRes      `json:"device_status"`
	DroneStatus     usecase.GetDroneStatusResp      `json:"drone_status"`
	MonitorStatus   usecase.GetMonitoringStatusResp `json:"monitor_status"`
	PerizinanStatus usecase.GetPerizinanStatusResp  `json:"perizinan_status"`
	ServiceStatus   usecase.GetServiceStatusResp    `json:"service_status"`
	SpeedStatus     usecase.GetSpeedtestStatusRes   `json:"speed_status"`
	WeatherForecast usecase.GetWeatherForecastRes   `json:"weather_forecast"`
	TMAHuluBendung  float32                         `json:"tma_hulu_bendung"`
}

type AISystemSummaryUseCase = core.ActionHandler[AISystemSummaryReq, AISystemSummaryRes]

func ImplAISystemSummary(
	getGeneralInfoUseCase usecase.GeneralInfoUseCase,
	getDeviceStatusUseCase usecase.GetDeviceStatusUseCase,
	getDroneStatusUseCase usecase.GetDroneStatusUseCase,
	getMonitorStatusUseCase usecase.GetMonitoringStatusUseCase,
	getPerizinanStatusUseCase usecase.GetPerizinanStatusUseCase,
	getServiceStatusUseCase usecase.GetServiceStatusUseCase,
	getSpeedStatusUseCase usecase.GetSpeedtestStatusUseCase,
	getWeatherForecastUseCase usecase.GetWeatherForecastUseCase,
	getLatestTMAHuluBendung usecase.ChartTMAUseCase,
) AISystemSummaryUseCase {
	return func(ctx context.Context, request AISystemSummaryReq) (*AISystemSummaryRes, error) {
		generalInfo, err := getGeneralInfoUseCase(ctx, usecase.GeneralInfoReq{})
		if err != nil {
			return nil, err
		}

		deviceStatus, err := getDeviceStatusUseCase(ctx, usecase.GetDeviceStatusReq{})
		if err != nil {
			return nil, err
		}

		droneStatus, err := getDroneStatusUseCase(ctx, usecase.GetDroneStatusReq{})
		if err != nil {
			return nil, err
		}

		monitorStatus, err := getMonitorStatusUseCase(ctx, usecase.GetMonitoringStatusReq{})
		if err != nil {
			return nil, err
		}

		perizinanStatus, err := getPerizinanStatusUseCase(ctx, usecase.GetPerizinanStatusReq{})
		if err != nil {
			return nil, err
		}

		serviceStatus, err := getServiceStatusUseCase(ctx, usecase.GetServiceStatusReq{})
		if err != nil {
			return nil, err
		}

		speedStatus, err := getSpeedStatusUseCase(ctx, usecase.GetSpeedtestStatusReq{})
		if err != nil {
			return nil, err
		}

		weatherForecast, err := getWeatherForecastUseCase(ctx, usecase.GetWeatherForecastReq{
			ADM4: "32.79.03.1004",
		})
		if err != nil {
			return nil, err
		}

		latestHuluBendungTMA, err := getLatestTMAHuluBendung(ctx, usecase.ChartTMAReq{
			WaterChannelDoorID: 7,
			MinTime:            time.Now().Add(-5 * time.Minute),
			MaxTime:            time.Now(),
		})
		if err != nil {
			return nil, err
		}

		return &AISystemSummaryRes{
			DeviceStatus:    *deviceStatus,
			DroneStatus:     *droneStatus,
			MonitorStatus:   *monitorStatus,
			PerizinanStatus: *perizinanStatus,
			ServiceStatus:   *serviceStatus,
			SpeedStatus:     *speedStatus,
			WeatherForecast: *weatherForecast,
			GeneralInfo:     *generalInfo,
			TMAHuluBendung:  float32(latestHuluBendungTMA.Charts[len(latestHuluBendungTMA.Charts)-1].WaterLevel),
		}, nil
	}
}
