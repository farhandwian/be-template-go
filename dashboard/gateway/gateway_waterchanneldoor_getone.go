package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"
	"shared/middleware"

	"gorm.io/gorm"
)

type WaterChannelDoorGetOneReq struct {
	ID int
}
type WaterChannelDoorGetOneRes struct {
	WaterChannelDoor model.WaterChannelDoor
}

type WaterChannelDoorGetOneGateway = core.ActionHandler[WaterChannelDoorGetOneReq, WaterChannelDoorGetOneRes]

func ImplWaterChannelDoorGetOneGateway(db *gorm.DB) WaterChannelDoorGetOneGateway {
	return func(ctx context.Context, request WaterChannelDoorGetOneReq) (*WaterChannelDoorGetOneRes, error) {

		query := middleware.GetDBFromContext(ctx, db)

		var obj model.WaterChannelDoor

		if err := query.Where("external_id=?", request.ID).First(&obj).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, core.NewInternalServerError(err)
		}
		return &WaterChannelDoorGetOneRes{WaterChannelDoor: obj}, nil
	}
}
