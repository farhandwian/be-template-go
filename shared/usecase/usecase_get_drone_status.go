package usecase

import (
	"context"
	"shared/core"
	"time"
)

type GetDroneStatusReq struct {
}

type GetDroneStatusResp struct {
	LatestInspections   string `json:"latest_inspections"`
	DetectedObjectCount int    `json:"detected_object_count"`
}

type GetDroneStatusUseCase = core.ActionHandler[GetDroneStatusReq, GetDroneStatusResp]

func ImplGetDroneStatusUseCase() GetDroneStatusUseCase {
	return func(ctx context.Context, req GetDroneStatusReq) (*GetDroneStatusResp, error) {
		desiredDate := time.Date(2024, 12, 12, 0, 0, 0, 0, time.UTC)

		return &GetDroneStatusResp{
			LatestInspections:   desiredDate.Format(time.RFC3339),
			DetectedObjectCount: 10,
		}, nil
	}
}
