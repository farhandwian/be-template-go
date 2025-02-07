package gateway

import (
	"context"
	"shared/core"
	"shared/helper"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type DoorControlGetAllReq struct {
	Page               int
	Size               int
	WaterChannelDoorID int
	DeviceID           int
}

type DoorControlGetAllRes struct {
	Items []model.DoorControl
	Count int64
}

type DoorControlGetAll = core.ActionHandler[DoorControlGetAllReq, DoorControlGetAllRes]

func ImplDoorControlGetAll(db *gorm.DB) DoorControlGetAll {
	return func(ctx context.Context, req DoorControlGetAllReq) (*DoorControlGetAllRes, error) {

		query := middleware.
			GetDBFromContext(ctx, db).
			Model(&model.DoorControl{})

		var count int64

		if req.WaterChannelDoorID != 0 {
			query = query.Where("water_channel_door_id = ?", req.WaterChannelDoorID)
		}

		if req.DeviceID != 0 {
			query = query.Where("device_id = ?", req.DeviceID)
		}

		if err := query.
			Count(&count).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		page, size := helper.ValidatePageSize(req.Page, req.Size)

		var objs []model.DoorControl

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DoorControlGetAllRes{
			Items: objs,
			Count: count,
		}, nil
	}
}
