package usecase

import (
	"context"
	"monitoring/gateway"
	"shared/core"
)

type ActivityMonitorDeleteReq struct {
	ID string `json:"id"`
}

type ActivityMonitorDeleteRes struct{}

type ActivityMonitorDelete = core.ActionHandler[ActivityMonitorDeleteReq, ActivityMonitorDeleteRes]

func ImplActivityMonitorDelete(delete gateway.ActivityMonitorDelete) ActivityMonitorDelete {
	return func(ctx context.Context, req ActivityMonitorDeleteReq) (*ActivityMonitorDeleteRes, error) {

		if _, err := delete(ctx, gateway.ActivityMonitorDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &ActivityMonitorDeleteRes{}, nil
	}
}
