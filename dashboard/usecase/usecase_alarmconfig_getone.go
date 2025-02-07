package usecase

import (
	"context"
	"shared/core"
	sg "shared/gateway"
	"shared/model"
)

type AlarmConfigGetOneReq struct {
	ID model.AlarmConfigID `json:"id"`
}

type AlarmConfigGetOneRes struct {
	AlarmConfig *model.AlarmConfig `json:"alarm_config"`
}

type AlarmConfigGetOne = core.ActionHandler[AlarmConfigGetOneReq, AlarmConfigGetOneRes]

func ImplAlarmConfigGetOne(getOne sg.AlarmConfigGetOne) AlarmConfigGetOne {
	return func(ctx context.Context, req AlarmConfigGetOneReq) (*AlarmConfigGetOneRes, error) {
		res, err := getOne(ctx, sg.AlarmConfigGetOneReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &AlarmConfigGetOneRes{AlarmConfig: res.Item}, nil
	}
}
