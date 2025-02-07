package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type AlarmConfigGetOneReq struct {
	ID model.AlarmConfigID
}

type AlarmConfigGetOneRes struct {
	Item *model.AlarmConfig
}

type AlarmConfigGetOne = core.ActionHandler[AlarmConfigGetOneReq, AlarmConfigGetOneRes]

func ImplAlarmConfigGetOne(db *gorm.DB) AlarmConfigGetOne {
	return func(ctx context.Context, req AlarmConfigGetOneReq) (*AlarmConfigGetOneRes, error) {

		var obj model.AlarmConfig

		if err := middleware.
			GetDBFromContext(ctx, db).
			First(&obj, "id = ?", req.ID).
			Error; err != nil {

			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("id %v is not found", req.ID)
			}

			return nil, core.NewInternalServerError(err)
		}

		return &AlarmConfigGetOneRes{Item: &obj}, nil
	}
}
