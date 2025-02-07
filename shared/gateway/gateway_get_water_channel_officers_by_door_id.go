package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetWaterChannelOfficersByDoorIDReq struct {
	WaterChannelDoorID int
}

type GetWaterChannelOfficersByDoorIDRes struct {
	Officers []model.WaterChannelOfficer
}

type GetWaterChannelOfficersByDoorID = core.ActionHandler[GetWaterChannelOfficersByDoorIDReq, GetWaterChannelOfficersByDoorIDRes]

func ImplGetWaterChannelOfficersByDoorID(db *gorm.DB) GetWaterChannelOfficersByDoorID {
	return func(ctx context.Context, request GetWaterChannelOfficersByDoorIDReq) (*GetWaterChannelOfficersByDoorIDRes, error) {
		var officers []model.WaterChannelOfficer

		result := db.Where("water_channel_door_id = ?", request.WaterChannelDoorID).Find(&officers)
		if result.Error != nil {
			return nil, core.NewInternalServerError(result.Error)
		}

		return &GetWaterChannelOfficersByDoorIDRes{
			Officers: officers,
		}, nil
	}
}
