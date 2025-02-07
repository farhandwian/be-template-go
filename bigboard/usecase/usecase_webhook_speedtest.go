package usecase

import (
	"context"
	"github.com/google/uuid"
	"shared/core"
	"shared/gateway"
	sharedModel "shared/model"
	"time"
)

type WebhookSpeedTestReq struct {
	Message string `json:"message"`
}

type WebhookSpeedTestRes struct {
}

type WebhookSpeedTestUseCase = core.ActionHandler[WebhookSpeedTestReq, WebhookSpeedTestRes]

func ImplWebhookSpeedTestUseCase(createActivityMonitoring gateway.CreateActivityMonitoringGateway) WebhookSpeedTestUseCase {
	return func(ctx context.Context, req WebhookSpeedTestReq) (*WebhookSpeedTestRes, error) {
		_, err := createActivityMonitoring(ctx, gateway.CreateActivityMonitoringReq{
			ActivityMonitor: sharedModel.ActivityMonitor{
				ID:           uuid.NewString(),
				UserName:     "System",
				Category:     "System",
				ActivityTime: time.Now(),
				Description:  req.Message,
			},
		})

		if err != nil {
			return nil, err
		}

		return &WebhookSpeedTestRes{}, nil
	}
}
