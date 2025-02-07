package gateway

import (
	"context"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type WaterChannelDoorByKeywordReq struct {
	Keyword string
}

type WaterChannelDoor struct {
	ExternalID int    `json:"water_channel_door_id"`
	Name       string `json:"name"`
}

type WaterChannelDoorByKeywordRes struct {
	Items []WaterChannelDoor
}

type WaterChannelDoorByKeyword = core.ActionHandler[WaterChannelDoorByKeywordReq, WaterChannelDoorByKeywordRes]

func ImplWaterChannelDoorByKeyword(db *gorm.DB) WaterChannelDoorByKeyword {
	return func(ctx context.Context, req WaterChannelDoorByKeywordReq) (*WaterChannelDoorByKeywordRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var waterChannelDoors []WaterChannelDoor

		if err := query.
			Where("LOWER(name) LIKE LOWER(?)", req.Keyword+"%").
			Find(&waterChannelDoors).
			Error; err != nil {

			if err == gorm.ErrRecordNotFound {
				return &WaterChannelDoorByKeywordRes{Items: []WaterChannelDoor{}}, nil
			}

			return nil, core.NewInternalServerError(err)
		}

		return &WaterChannelDoorByKeywordRes{Items: waterChannelDoors}, nil
	}
}
