package gateway

import (
	"context"
	"encoding/json"
	"shared/core"
	"shared/helper/cronjob"
	"shared/model"
	"time"
)

type CronjobSetDoorControlReq struct {
	TargetDate         time.Time
	DoorControlID      model.DoorControlID
	OpenTarget         float32
	WaterChannelDoorID int
	DeviceID           int
	IPAddress          string
}

type CronjobSetDoorControlRes struct {
}

type CronjobSetDoorControlGateway = core.ActionHandler[CronjobSetDoorControlReq, CronjobSetDoorControlRes]

func ImplCronjobSetDoorControl(cj *cronjob.CronJob) CronjobSetDoorControlGateway {
	return func(ctx context.Context, req CronjobSetDoorControlReq) (*CronjobSetDoorControlRes, error) {

		dcp := model.DoorControlPayload{
			DoorControlID:      req.DoorControlID,
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
			OpenTarget:         req.OpenTarget,
			IPAddress:          req.IPAddress,
		}

		data, err := json.Marshal(dcp)
		if err != nil {
			return nil, err
		}

		if err := cj.ExecuteLater(ctx, req.TargetDate, cronjob.FuncID(req.DoorControlID), "door_control", data); err != nil {
			return nil, err
		}

		return &CronjobSetDoorControlRes{}, nil
	}
}
