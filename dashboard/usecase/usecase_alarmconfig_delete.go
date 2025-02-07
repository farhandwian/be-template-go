package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
	sg "shared/gateway"
	"shared/model"
)

type AlarmConfigDeleteReq struct {
	ID model.AlarmConfigID `json:"id"`
}

type AlarmConfigDeleteRes struct{}

type AlarmConfigDelete = core.ActionHandler[AlarmConfigDeleteReq, AlarmConfigDeleteRes]

func ImplAlarmConfigDelete(
	deleteDB sg.AlarmConfigDelete,
	deleteGrafana gateway.AlarmConfigXDelete,
) AlarmConfigDelete {
	return func(ctx context.Context, req AlarmConfigDeleteReq) (*AlarmConfigDeleteRes, error) {

		if _, err := deleteDB(ctx, sg.AlarmConfigDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		if _, err := deleteGrafana(ctx, gateway.AlarmConfigXDeleteReq{UID: req.ID}); err != nil {
			return nil, err
		}

		return &AlarmConfigDeleteRes{}, nil
	}
}
