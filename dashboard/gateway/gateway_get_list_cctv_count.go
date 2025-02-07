package gateway

import (
	"bigboard/model"
	"context"
	"errors"
	"shared/core"

	"gorm.io/gorm"
)

type GetListCCTVCountReq struct {
	WaterChannelDoorID int
	DeviceID           int
}

type GetListCCTVCountRes struct {
	CCTVDevices []*model.WaterChannelCCTVCount
}

type GetListCCTVCountGateway = core.ActionHandler[GetListCCTVCountReq, GetListCCTVCountRes]

func ImplGetListCCTVCountGateway(db *gorm.DB) GetListCCTVCountGateway {
	return func(ctx context.Context, req GetListCCTVCountReq) (*GetListCCTVCountRes, error) {
		var cctvCount []*model.WaterChannelCCTVCount

		rawQuery := `
            SELECT 
                wcd.external_id AS water_channel_door_id, 
                COALESCE(device_counts.cctv_count, 0) AS cctv_count
            FROM water_channel_doors wcd
            LEFT JOIN (
                SELECT 
                    water_channel_door_id, 
                    COUNT(*) AS cctv_count 
                FROM water_channel_devices 
                WHERE category = 'cctv'
                GROUP BY water_channel_door_id
            ) AS device_counts 
            ON wcd.external_id = device_counts.water_channel_door_id
        `

		err := db.Raw(rawQuery).Scan(&cctvCount).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, core.NewInternalServerError(err)
		}

		return &GetListCCTVCountRes{
			CCTVDevices: cctvCount,
		}, nil
	}
}
