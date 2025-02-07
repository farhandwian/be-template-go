package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetListWaterChannelDoorReq struct {
	Name string
}

type GetListWaterChannelDoorRes struct {
	WaterChannels   []model.WaterChannelDoor
	WaterChannelMap map[int]model.WaterChannelDoor
}

type GetListWaterChannelDoorGateway = core.ActionHandler[GetListWaterChannelDoorReq, GetListWaterChannelDoorRes]

func ImplGetListWaterChannelDoor(db *gorm.DB) GetListWaterChannelDoorGateway {
	return func(ctx context.Context, request GetListWaterChannelDoorReq) (*GetListWaterChannelDoorRes, error) {

		var waterChannels []model.WaterChannelDoor

		query := db.Select("*,water_channel_doors.external_id as external_id, water_channel_doors.name as name, wc.name as water_channel_name,wc.external_id as water_channel_id").Table("water_channel_doors").
			Joins("LEFT JOIN water_channels wc ON water_channel_doors.water_channel_id = wc.external_id")
		if request.Name != "" {
			query = query.Where("water_channel_doors.name LIKE ?", "%"+request.Name+"%")
		}
		err := query.Find(&waterChannels).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListWaterChannelDoorRes{
			WaterChannels: waterChannels,
		}, nil
	}
}

func ImplGetListWaterChannelDoorWithMap(db *gorm.DB) GetListWaterChannelDoorGateway {
	return func(ctx context.Context, request GetListWaterChannelDoorReq) (*GetListWaterChannelDoorRes, error) {

		waterChannelMap := make(map[int]model.WaterChannelDoor)

		query := db.Table("water_channel_doors")
		if request.Name != "" {
			query = query.Where("name LIKE ?", "%"+request.Name+"%")
		}
		rows, err := query.Rows()
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}
		defer rows.Close()

		for rows.Next() {
			var obj model.WaterChannelDoor
			if err := db.ScanRows(rows, &obj); err != nil {
				return nil, err
			}
			waterChannelMap[obj.ExternalID] = obj
		}

		return &GetListWaterChannelDoorRes{
			WaterChannelMap: waterChannelMap,
		}, nil
	}
}
