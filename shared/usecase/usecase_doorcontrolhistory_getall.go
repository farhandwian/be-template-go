package usecase

import (
	"context"
	"shared/core"
	"shared/gateway"
	"shared/model"
)

type DoorControlHistoryGetAllReq struct {
	Page               int `json:"page"`
	Size               int `json:"size"`
	WaterChannelDoorID int `json:"water_channel_door_id"` // pintu
	DeviceID           int `json:"device_id"`             // perangkat
}

type DoorControlHistoryGetAllRes struct {
	DoorControlHistories []model.DoorControlHistory `json:"door_control_histories"`
	Metadata             *Metadata                  `json:"metadata"`
}

type DoorControlHistoryGetAll = core.ActionHandler[DoorControlHistoryGetAllReq, DoorControlHistoryGetAllRes]

func ImplDoorControlHistoryGetAll(
	getAll gateway.DoorControlHistoryGetAll,
) DoorControlHistoryGetAll {
	return func(ctx context.Context, req DoorControlHistoryGetAllReq) (*DoorControlHistoryGetAllRes, error) {

		res, err := getAll(ctx, gateway.DoorControlHistoryGetAllReq{
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

		return &DoorControlHistoryGetAllRes{
			DoorControlHistories: res.Items,
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
