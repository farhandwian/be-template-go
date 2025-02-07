package usecase

import (
	"context"
	"shared/core"
)

type HealthCheckReq struct {
}

type HealthCheckRes struct {
	Status string `json:"status"`
}

type HealthCheckUseCase = core.ActionHandler[HealthCheckReq, HealthCheckRes]

func ImplHealthCheckUseCase() HealthCheckUseCase {
	return func(ctx context.Context, req HealthCheckReq) (*HealthCheckRes, error) {

		return &HealthCheckRes{Status: "OK"}, nil
	}
}
