package gateway

import (
	"context"
	"errors"
	"fmt"
	"shared/core"
	"shared/middleware"
	sharedModel "shared/model"

	"gorm.io/gorm"
)

type ActivityMonitorGetDetailReq struct {
	ID string
}

type ActivityMonitorResp struct {
	ActivityMonitor *sharedModel.ActivityMonitor
}

type ActivityMonitorGetOneGateway = core.ActionHandler[ActivityMonitorGetDetailReq, ActivityMonitorResp]

func ImplActivityMonitorGetOneGateway(db *gorm.DB) ActivityMonitorGetOneGateway {
	return func(ctx context.Context, req ActivityMonitorGetDetailReq) (*ActivityMonitorResp, error) {

		var activityMonitor sharedModel.ActivityMonitor

		if err := middleware.
			GetDBFromContext(ctx, db).
			First(&activityMonitor, "id = ?", req.ID).
			Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("id %v is not found", req.ID)
			}

			return nil, core.NewInternalServerError(err)
		}

		return &ActivityMonitorResp{ActivityMonitor: &activityMonitor}, nil
	}
}
