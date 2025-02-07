package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	sharedModel "shared/model"

	"gorm.io/gorm"
)

type ActivityMonitorDeleteReq struct {
	ID string
}

type ActivityMonitorDeleteRes struct{}

type ActivityMonitorDelete = core.ActionHandler[ActivityMonitorDeleteReq, ActivityMonitorDeleteRes]

func ImplActivityMonitorDelete(db *gorm.DB) ActivityMonitorDelete {
	return func(ctx context.Context, req ActivityMonitorDeleteReq) (*ActivityMonitorDeleteRes, error) {

		if err := middleware.
			GetDBFromContext(ctx, db).
			Delete(&sharedModel.ActivityMonitor{}, "id = ?", req.ID).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &ActivityMonitorDeleteRes{}, nil
	}
}
