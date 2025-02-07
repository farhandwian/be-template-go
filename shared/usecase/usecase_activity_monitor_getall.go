package usecase

import (
	"shared/core"
	"shared/gateway"
	sharedModel "shared/model"
)

import (
	"context"
)

type ActivityMonitorGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type ActivityMonitorGetAllUseCaseRes struct {
	ActivityMonitors []sharedModel.ActivityMonitor `json:"activity_monitors"`
	Metadata         *Metadata                     `json:"metadata"`
}

type ActivityMonitorGetAllUseCase = core.ActionHandler[ActivityMonitorGetAllUseCaseReq, ActivityMonitorGetAllUseCaseRes]

func ImpActivityMonitorGetAllUseCase(getAllActivityMonitors gateway.GetActivityMonitoringGateway) ActivityMonitorGetAllUseCase {
	return func(ctx context.Context, req ActivityMonitorGetAllUseCaseReq) (*ActivityMonitorGetAllUseCaseRes, error) {

		res, err := getAllActivityMonitors(ctx, gateway.GetActivityMonitoringReq{
			Page: req.Page,
			Size: req.Size,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &ActivityMonitorGetAllUseCaseRes{
			ActivityMonitors: res.ActivityMonitorings,
			Metadata: &Metadata{
				Pagination: Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
