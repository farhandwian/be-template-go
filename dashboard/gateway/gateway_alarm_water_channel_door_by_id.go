package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type WaterChannelDoorByIDReq struct {
	ID int
}

type WaterChannelDoorByIDRes struct {
	Items []WaterChannelDoor
}

type WaterChannelDoorByID = core.ActionHandler[WaterChannelDoorByIDReq, WaterChannelDoorByIDRes]

func ImplWaterChannelDoorByID(db *gorm.DB) WaterChannelDoorByID {
	return func(ctx context.Context, req WaterChannelDoorByIDReq) (*WaterChannelDoorByIDRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var waterChannelDoors []WaterChannelDoor

		if err := query.
			Where("external_id", req.ID).
			Find(&waterChannelDoors).
			Error; err != nil {

			if err == gorm.ErrRecordNotFound {
				return &WaterChannelDoorByIDRes{Items: []WaterChannelDoor{}}, nil
			}

			return nil, core.NewInternalServerError(err)
		}

		return &WaterChannelDoorByIDRes{Items: waterChannelDoors}, nil
	}
}
