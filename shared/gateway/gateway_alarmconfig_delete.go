package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type AlarmConfigDeleteReq struct {
	ID model.AlarmConfigID
}

type AlarmConfigDeleteRes struct{}

type AlarmConfigDelete = core.ActionHandler[AlarmConfigDeleteReq, AlarmConfigDeleteRes]

func ImplAlarmConfigDelete(db *gorm.DB) AlarmConfigDelete {
	return func(ctx context.Context, req AlarmConfigDeleteReq) (*AlarmConfigDeleteRes, error) {

		if err := middleware.
			GetDBFromContext(ctx, db).
			Delete(&model.AlarmConfig{}, "id = ?", req.ID).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &AlarmConfigDeleteRes{}, nil
	}
}
