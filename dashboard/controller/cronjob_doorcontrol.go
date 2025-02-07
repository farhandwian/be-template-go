package controller

import (
	"context"
	"dashboard/usecase"
	"encoding/json"
	"shared/helper/cronjob"
	"shared/model"
	"time"
)

func RunDoorControl(cj *cronjob.CronJob, u usecase.DoorControlRunScheduled) {

	cj.RegisterFunction("door_control", func(data []byte) error {

		var dcp model.DoorControlPayload
		if err := json.Unmarshal(data, &dcp); err != nil {
			return err
		}

		if _, err := u(context.Background(), usecase.DoorControlRunScheduledReq{
			Now:                time.Now().In(time.Local),
			WaterChannelDoorID: dcp.WaterChannelDoorID,
			DeviceID:           dcp.DeviceID,
			OpenTarget:         dcp.OpenTarget,
			DoorControlID:      dcp.DoorControlID,
			IPAddress:          dcp.IPAddress,
		}); err != nil {
			return err
		}

		return nil
	})

}
