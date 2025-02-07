package gateway

import (
	"bigboard/model"
	"context"
	"shared/core"

	"gorm.io/gorm"
)

type GetListWaterChannelOfficerCountReq struct{}

type GetListWaterChannelOfficerCountRes struct {
	OfficerCounts   []model.WaterChannelOfficerCount
	OfficerCountMap map[int]model.WaterChannelOfficerCount
}

type GetListWaterChannelDoorOfficerCountGateway = core.ActionHandler[GetListWaterChannelOfficerCountReq, GetListWaterChannelOfficerCountRes]

func ImplGetListWaterChannelOfficerCount(db *gorm.DB) GetListWaterChannelDoorOfficerCountGateway {
	return func(ctx context.Context, request GetListWaterChannelOfficerCountReq) (*GetListWaterChannelOfficerCountRes, error) {

		var officerCounts []model.WaterChannelOfficerCount

		err := db.Table("water_channel_officers").
			Select("water_channel_door_id AS water_channel_door_id, COUNT(*) AS officer_count").
			Group("water_channel_door_id").
			Scan(&officerCounts).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetListWaterChannelOfficerCountRes{
			OfficerCounts: officerCounts,
		}, nil
	}
}

func ImplGetListWaterChannelOfficerCountWithMap(db *gorm.DB) GetListWaterChannelDoorOfficerCountGateway {
	return func(ctx context.Context, request GetListWaterChannelOfficerCountReq) (*GetListWaterChannelOfficerCountRes, error) {

		officerCounts := make(map[int]model.WaterChannelOfficerCount)

		rows, err := db.Table("water_channel_officers").
			Select("water_channel_door_id AS water_channel_door_id, COUNT(*) AS officer_count").
			Group("water_channel_door_id").
			Scan(&officerCounts).
			Rows()
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		defer rows.Close()

		for rows.Next() {
			var obj model.WaterChannelOfficerCount
			if err := db.ScanRows(rows, &obj); err != nil {
				return nil, err
			}
			officerCounts[obj.WaterChannelDoorID] = obj
		}

		return &GetListWaterChannelOfficerCountRes{
			OfficerCountMap: officerCounts,
		}, nil
	}
}
