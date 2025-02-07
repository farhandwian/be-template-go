package gateway

import (
	"context"
	"shared/core"
	"shared/helper/cronjob"
	"shared/model"
)

type CronjobCancelDoorControlReq struct {
	DoorControlID model.DoorControlID
}

type CronjobCancelDoorControlRes struct {
}

type CronjobCancelDoorControlGateway = core.ActionHandler[CronjobCancelDoorControlReq, CronjobCancelDoorControlRes]

func ImplCronjobCancelDoorControl(cj *cronjob.CronJob) CronjobCancelDoorControlGateway {
	return func(ctx context.Context, req CronjobCancelDoorControlReq) (*CronjobCancelDoorControlRes, error) {

		if err := cj.CancelFunc(ctx, cronjob.FuncID(req.DoorControlID)); err != nil {
			return nil, err
		}

		return &CronjobCancelDoorControlRes{}, nil
	}
}
