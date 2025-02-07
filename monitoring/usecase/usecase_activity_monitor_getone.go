package usecase

import (
	"context"
	"monitoring/gateway"
	"shared/core"
	"time"
)

type ActivityMonitorGetOneReq struct {
	ID string `json:"id"`
}

type ActivityMonitorGetOneRes struct {
	ActivityMonitor *ActivityMonitor `json:"activity_monitor"`
}
type ActivityMonitor struct {
	ID           string    `json:"id" `
	UserName     string    ` json:"user_name"`
	Category     string    ` json:"category"`
	ActivityTime time.Time `json:"activity_time"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type ActivityMonitorGetOneUseCase = core.ActionHandler[ActivityMonitorGetOneReq, ActivityMonitorGetOneRes]

func ImplActivityMonitorGetOne(getOne gateway.ActivityMonitorGetOneGateway) ActivityMonitorGetOneUseCase {
	return func(ctx context.Context, req ActivityMonitorGetOneReq) (*ActivityMonitorGetOneRes, error) {
		res, err := getOne(ctx, gateway.ActivityMonitorGetDetailReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &ActivityMonitorGetOneRes{ActivityMonitor: &ActivityMonitor{
			ID:           res.ActivityMonitor.ID,
			UserName:     res.ActivityMonitor.UserName,
			Category:     res.ActivityMonitor.Category,
			ActivityTime: res.ActivityMonitor.ActivityTime,
			Description:  res.ActivityMonitor.Description,
			CreatedAt:    res.ActivityMonitor.CreatedAt,
			UpdatedAt:    res.ActivityMonitor.UpdatedAt,
		}}, nil
	}
}
