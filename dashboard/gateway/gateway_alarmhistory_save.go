package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type AlarmHistorySaveReq struct {
	AlarmHistory []model.AlarmHistory
}

type AlarmHistorySaveRes struct {
}

type AlarmHistorySave = core.ActionHandler[AlarmHistorySaveReq, AlarmHistorySaveRes]

func ImplAlarmHistorySave(db *gorm.DB) AlarmHistorySave {
	return func(ctx context.Context, req AlarmHistorySaveReq) (*AlarmHistorySaveRes, error) {

		if err := middleware.
			GetDBFromContext(ctx, db).
			Save(req.AlarmHistory).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &AlarmHistorySaveRes{}, nil
	}
}
