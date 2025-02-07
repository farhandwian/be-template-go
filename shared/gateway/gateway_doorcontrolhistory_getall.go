package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/helper"
	"shared/middleware"
	"shared/model"

	"gorm.io/gorm"
)

type DoorControlHistoryGetAllReq struct {
	Page               int
	Size               int
	WaterChannelDoorID int
	DeviceID           int
}

type DoorControlHistoryGetAllRes struct {
	Items []model.DoorControlHistory
	Count int64
}

type DoorControlHistoryGetAll = core.ActionHandler[DoorControlHistoryGetAllReq, DoorControlHistoryGetAllRes]

func ImplDoorControlHistoryGetAll(db *gorm.DB) DoorControlHistoryGetAll {
	return func(ctx context.Context, req DoorControlHistoryGetAllReq) (*DoorControlHistoryGetAllRes, error) {

		query := middleware.
			GetDBFromContext(ctx, db).
			Model(&model.DoorControlHistory{})

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
		allowerdSortBy := map[string]bool{
			"date": true,
		}

		sortBy, sortOrder, err := helper.ValidateSortParams(allowerdSortBy, "date", "desc", "date")
		if err != nil {
			return nil, err
		}

		orderCaluse := fmt.Sprintf("%s %s", sortBy, sortOrder)

		var objs []model.DoorControlHistory

		if err := query.
			Offset((page - 1) * size).
			Limit(size).
			Order(orderCaluse).
			Find(&objs).
			Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &DoorControlHistoryGetAllRes{
			Items: objs,
			Count: count,
		}, nil
	}
}
