package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
)

type GetMonitoringStatusReq struct {
}

type GetMonitoringStatusResp struct {
	WarningCount  int `json:"warning_count"`
	CriticalCount int `json:"critical_count"`
}

type GetMonitoringStatusUseCase = core.ActionHandler[GetMonitoringStatusReq, GetMonitoringStatusResp]

func ImplGetMonitoringStatusUseCase(GetAlarmCountGateway gateway.GetAlarmCountGateway) GetMonitoringStatusUseCase {
	return func(ctx context.Context, req GetMonitoringStatusReq) (*GetMonitoringStatusResp, error) {
		res, err := GetAlarmCountGateway(ctx, gateway.GetAlarmCountReq{})
		if err != nil {
			return nil, err
		}
		return &GetMonitoringStatusResp{
			WarningCount:  res.WarningCount,
			CriticalCount: res.CriticalCount,
		}, nil

	}
}
