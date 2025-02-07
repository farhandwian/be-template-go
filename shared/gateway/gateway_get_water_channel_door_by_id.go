package gateway

import (
	"context"
	"errors"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetWaterChannelDoorByIDReq struct {
	WaterChannelDoorID int
}

type GetWaterChannelDoorByIDRes struct {
	WaterChannelDoor model.WaterChannelDoor
}

type GetWaterChannelDoorByID = core.ActionHandler[GetWaterChannelDoorByIDReq, GetWaterChannelDoorByIDRes]

func ImplGetWaterChannelDoorByID(db *gorm.DB) GetWaterChannelDoorByID {
	return func(ctx context.Context, request GetWaterChannelDoorByIDReq) (*GetWaterChannelDoorByIDRes, error) {
		var waterChannelDoor model.WaterChannelDoor

		result := db.Where("external_id = ?", request.WaterChannelDoorID).First(&waterChannelDoor)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, errors.New("water channel door not found")
			}
			return nil, result.Error
		}

		return &GetWaterChannelDoorByIDRes{
			WaterChannelDoor: waterChannelDoor,
		}, nil
	}
}
