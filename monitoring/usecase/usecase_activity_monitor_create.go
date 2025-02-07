package usecase

import (
	"context"
	"github.com/google/uuid"
	"shared/core"
	"shared/gateway"
	sharedModel "shared/model"
	"time"
)

type ActivityMonitorCreateReq struct {
	UserName    string `json:"user_name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type ActivityMonitorCreateRes struct {
}

type ActivityMonitorCreateUseCase = core.ActionHandler[ActivityMonitorCreateReq, ActivityMonitorCreateRes]

func ImplActivityMonitorCreateUseCase(
	createActivityMonitoring gateway.CreateActivityMonitoringGateway,
) ActivityMonitorCreateUseCase {
	return func(ctx context.Context, req ActivityMonitorCreateReq) (*ActivityMonitorCreateRes, error) {

		_, err := createActivityMonitoring(ctx, gateway.CreateActivityMonitoringReq{
			ActivityMonitor: sharedModel.ActivityMonitor{
				ID:           uuid.NewString(),
				UserName:     req.UserName,
				Category:     req.Category,
				ActivityTime: time.Now(),
				Description:  req.Description,
			},
		})

		if err != nil {
			return nil, err
		}

		return &ActivityMonitorCreateRes{}, nil
	}
}
