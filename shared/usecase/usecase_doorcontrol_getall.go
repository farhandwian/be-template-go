package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"shared/model"
)

type DoorControlGetAllReq struct {
	Page               int `json:"page"`
	Size               int `json:"size"`
	WaterChannelDoorID int `json:"water_channel_door_id"` // pintu
	DeviceID           int `json:"device_id"`             // perangkat
}

type DoorControlGetAllRes struct {
	DoorControls []model.DoorControl `json:"door_controls"`
	Metadata     *Metadata           `json:"metadata"`
}

type DoorControlGetAll = core.ActionHandler[DoorControlGetAllReq, DoorControlGetAllRes]

func ImplDoorControlGetAll(
	getAll gateway.DoorControlGetAll,
) DoorControlGetAll {
	return func(ctx context.Context, req DoorControlGetAllReq) (*DoorControlGetAllRes, error) {

		res, err := getAll(ctx, gateway.DoorControlGetAllReq{
			Page:               req.Page,
			Size:               req.Size,
			WaterChannelDoorID: req.WaterChannelDoorID,
			DeviceID:           req.DeviceID,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &DoorControlGetAllRes{
			DoorControls: res.Items,
			Metadata: &Metadata{
				Pagination: Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
