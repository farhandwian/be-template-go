package usecase

import (
	"context"
	"errors"
	"shared/core"
	"shared/gateway"
	"strconv"
)

type GetDeviceStatusReq struct {
}

type GetDeviceStatusRes struct {
	Up    int `json:"up"`
	Down  int `json:"down"`
	Total int `json:"total"`
}

type GetDeviceStatusUseCase = core.ActionHandler[GetDeviceStatusReq, GetDeviceStatusRes]

func ImplGetDeviceStatusUseCase(
	getDeviceStatusGateway gateway.GetDeviceStatusGateway,
	getTmaDeviceCountGateway gateway.GetTmaDeviceCountGateway,
	getConfig gateway.GetConfigGateway,
) GetDeviceStatusUseCase {
	return func(ctx context.Context, req GetDeviceStatusReq) (*GetDeviceStatusRes, error) {
		// check if guest mode
		isGuestMode := false
		config, err := getConfig(ctx, gateway.GetConfigReq{Name: "GUEST_MODE"})
		if err == nil {
			isGuestMode = config.Config.Value == "1"
		}

		res, err := getDeviceStatusGateway(ctx, gateway.GetDeviceStatusReq{})
		if err != nil {
			return nil, err
		}

		var upDeviceCountInt, downDeviceCountInt int

		// Handle Up Device Count
		if len(res.UpDeviceResponse.PrometheusData.Result) > 0 && len(res.UpDeviceResponse.PrometheusData.Result[0].Value) > 1 {
			upDeviceCount := res.UpDeviceResponse.PrometheusData.Result[0].Value[1].(string)
			upDeviceCountInt, err = strconv.Atoi(upDeviceCount)
			if err != nil {
				return nil, errors.New("failed to parse up device count")
			}
		}

		// Handle Down Device Count
		if len(res.DownDeviceResponse.PrometheusData.Result) > 0 && len(res.DownDeviceResponse.PrometheusData.Result[0].Value) > 1 {
			downDeviceCount := res.DownDeviceResponse.PrometheusData.Result[0].Value[1].(string)
			downDeviceCountInt, err = strconv.Atoi(downDeviceCount)
			if err != nil {
				return nil, errors.New("failed to parse down device count")
			}
		}

		// Get TMA Device Count
		tmaDeviceCount, err := getTmaDeviceCountGateway(ctx, gateway.GetTmaDeviceCountReq{})
		if err != nil {
			return nil, err
		}

		// Calculate totals
		totalUpDevice := upDeviceCountInt + tmaDeviceCount.TMADeviceCount.DeviceOn
		totalDownDevice := downDeviceCountInt + tmaDeviceCount.TMADeviceCount.DeviceOff
		totalDevices := totalUpDevice + totalDownDevice

		if isGuestMode {
			return &GetDeviceStatusRes{
				Up:    totalUpDevice + totalDownDevice,
				Down:  0,
				Total: totalDevices,
			}, nil
		} else {
			return &GetDeviceStatusRes{
				Up:    totalUpDevice,
				Down:  totalDownDevice,
				Total: totalDevices,
			}, nil
		}
	}
}
