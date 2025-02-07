package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
	"sort"
)

type DeviceByWaterChannelDoorIdReq struct {
	WaterChannelDoorId int `json:"water_channel_door_id"`
}

type DeviceByWaterChannelDoorIdRes struct {
	Doors []gateway.WaterChannelDevice `json:"doors"`
}

type DeviceByWaterChannelDoorId = core.ActionHandler[DeviceByWaterChannelDoorIdReq, DeviceByWaterChannelDoorIdRes]

func ImplDeviceByWaterChannelDoorId(getAllDoor gateway.DeviceByWaterChannelDoorId) DeviceByWaterChannelDoorId {
	return func(ctx context.Context, req DeviceByWaterChannelDoorIdReq) (*DeviceByWaterChannelDoorIdRes, error) {

		res, err := getAllDoor(ctx, gateway.DeviceByWaterChannelDoorIdReq{WaterChannelDoorId: req.WaterChannelDoorId})
		if err != nil {
			return nil, err
		}

		sort.Slice(res.Items, func(i, j int) bool {
			return res.Items[i].Name < res.Items[j].Name
		})

		return &DeviceByWaterChannelDoorIdRes{Doors: res.Items}, nil
	}
}
