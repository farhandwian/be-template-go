package gateway

import (
	"bigboard/model"
	"context"
	"shared/core"

	"gorm.io/gorm"
)

type GetListWaterChannelCCTVCountReq struct{}

type GetListWaterChannelCCTVCountRes struct {
	WaterCCTVCounts   []model.WaterChannelCCTVCount
	WaterCCTVCountMap map[int]model.WaterChannelCCTVCount
}

type GetListWaterChannelDoorCCTVCountGateway = core.ActionHandler[GetListWaterChannelCCTVCountReq, GetListWaterChannelCCTVCountRes]

func ImplGetListWaterChannelCCTVCount(db *gorm.DB) GetListWaterChannelDoorCCTVCountGateway {
	return func(ctx context.Context, request GetListWaterChannelCCTVCountReq) (*GetListWaterChannelCCTVCountRes, error) {

		var cctvCounts []model.WaterChannelCCTVCount

		// TODO Optimasi bisa dengan cache
		err := db.Table("water_channel_devices").
			Select("water_channel_door_id AS water_channel_door_id, COUNT(*) AS cctv_count").
			Where("category = ?", "cctv").
			Group("water_channel_door_id").
			Scan(&cctvCounts).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		return &GetListWaterChannelCCTVCountRes{
			WaterCCTVCounts: cctvCounts,
		}, nil
	}
}

func ImplGetListWaterChannelCCTVCountWithMap(db *gorm.DB) GetListWaterChannelDoorCCTVCountGateway {
	return func(ctx context.Context, request GetListWaterChannelCCTVCountReq) (*GetListWaterChannelCCTVCountRes, error) {

		cctvCounts := make(map[int]model.WaterChannelCCTVCount)

		// TODO Optimasi bisa dengan cache
		rows, err := db.Table("water_channel_devices").
			Select("water_channel_door_id AS water_channel_door_id, COUNT(*) AS cctv_count").
			Where("category = ?", "cctv").
			Group("water_channel_door_id").
			Rows()
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		defer rows.Close()

		for rows.Next() {
			var obj model.WaterChannelCCTVCount
			if err := db.ScanRows(rows, &obj); err != nil {
				return nil, err
			}
			cctvCounts[obj.WaterChannelDoorID] = obj
		}

		return &GetListWaterChannelCCTVCountRes{
			WaterCCTVCountMap: cctvCounts,
		}, nil
	}
}
