package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type AlarmConfigSaveReq struct {
	AlarmConfig *model.AlarmConfig
}

type AlarmConfigSaveRes struct {
	ID model.AlarmConfigID
}

type AlarmConfigSave = core.ActionHandler[AlarmConfigSaveReq, AlarmConfigSaveRes]

func ImplAlarmConfigSave(db *gorm.DB) AlarmConfigSave {
	return func(ctx context.Context, req AlarmConfigSaveReq) (*AlarmConfigSaveRes, error) {

		if err := middleware.
			GetDBFromContext(ctx, db).
			Save(req.AlarmConfig).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &AlarmConfigSaveRes{ID: req.AlarmConfig.ID}, nil
	}
}
