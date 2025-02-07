package gateway

import (
	"context"
	"gorm.io/gorm"
	"shared/core"
	"shared/middleware"
	sharedModel "shared/model"
)

type CreateActivityMonitoringReq struct {
	ActivityMonitor sharedModel.ActivityMonitor
}

type CreateActivityMonitoringRes struct {
}

type CreateActivityMonitoringGateway = func(ctx context.Context, req CreateActivityMonitoringReq) (*CreateActivityMonitoringRes, error)

func ImplCreateActivityMonitoringGateway(db *gorm.DB) CreateActivityMonitoringGateway {
	return func(ctx context.Context, req CreateActivityMonitoringReq) (*CreateActivityMonitoringRes, error) {
		query := middleware.GetDBFromContext(ctx, db)

		if req.ActivityMonitor.Description == "" {
			return nil, nil
		}

		err := query.Save(&req.ActivityMonitor).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return nil, nil
	}
}
