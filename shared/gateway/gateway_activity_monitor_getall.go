package gateway

import (
	"context"
	"dashboard/gateway"
	"gorm.io/gorm"
	"shared/core"
	"shared/middleware"
	sharedModel "shared/model"
)

type GetActivityMonitoringReq struct {
	Page int
	Size int
}

type GetActivityMonitoringRes struct {
	ActivityMonitorings []sharedModel.ActivityMonitor
	Count               int64
}

type GetActivityMonitoringGateway = func(ctx context.Context, req GetActivityMonitoringReq) (*GetActivityMonitoringRes, error)

func ImplGetActivityMonitoringGateway(db *gorm.DB) GetActivityMonitoringGateway {
	return func(ctx context.Context, req GetActivityMonitoringReq) (*GetActivityMonitoringRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		var count int64

		if err := query.
			Model(&sharedModel.ActivityMonitor{}).
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := gateway.ValidatePageSize(req.Page, req.Size)

		var activityMonitorings []sharedModel.ActivityMonitor

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order("created_at desc").
			Find(&activityMonitorings).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetActivityMonitoringRes{
			ActivityMonitorings: activityMonitorings,
			Count:               count,
		}, nil
	}
}
