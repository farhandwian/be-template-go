package gateway

import (
	"context"
	"dashboard/model"
	"shared/core"

	"gorm.io/gorm"
)

type GetListWaterChannelDoorOfficerReq struct {
	Name string
}

type GetListWaterChannelDoorOfficerRes struct {
	WaterChannels []model.WaterChannelDoorOfficer
}

type GetListWaterChannelDoorOfficerGateway = core.ActionHandler[GetListWaterChannelDoorOfficerReq, GetListWaterChannelDoorOfficerRes]

func ImplGetListWaterChannelDoorOfficer(db *gorm.DB) GetListWaterChannelDoorOfficerGateway {
	return func(ctx context.Context, request GetListWaterChannelDoorOfficerReq) (*GetListWaterChannelDoorOfficerRes, error) {
		var waterChannelDoorOfficer []model.WaterChannelDoorOfficer

		query := db.Table("water_channel_doors").
			Select("water_channel_doors.external_id as water_channel_door_id,COALESCE(officer_counts.officer_count, 0) as officer_count").
			Joins("LEFT JOIN (SELECT water_channel_door_id, COUNT(*) as officer_count FROM water_channel_officers GROUP BY water_channel_door_id) as officer_counts ON water_channel_doors.external_id = officer_counts.water_channel_door_id")

		err := query.Find(&waterChannelDoorOfficer).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListWaterChannelDoorOfficerRes{
			WaterChannels: waterChannelDoorOfficer,
		}, nil
	}
}
