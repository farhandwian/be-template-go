package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetListWaterChannelDoorWithCCTVAndOfficerReq struct{}

type GetListWaterChannelDoorWithCCTVAndOfficerRes struct {
	WaterChannels []model.WaterChannelDoorWithAdditionalInfo `json:"water_channels"`
}

type GetListWaterChannelDoorWithCCTVAndOfficerGateway = core.ActionHandler[GetListWaterChannelDoorWithCCTVAndOfficerReq, GetListWaterChannelDoorWithCCTVAndOfficerRes]

func ImplGetListWaterChannelDoorWithCCTVAndOfficer(db *gorm.DB) GetListWaterChannelDoorWithCCTVAndOfficerGateway {
	return func(ctx context.Context, request GetListWaterChannelDoorWithCCTVAndOfficerReq) (*GetListWaterChannelDoorWithCCTVAndOfficerRes, error) {

		var waterChannels []model.WaterChannelDoorWithAdditionalInfo

		err := db.Table("water_channel_doors").
			Select("water_channel_doors.*, COALESCE(device_counts.cctv_count, 0) as cctv_count, COALESCE(officer_counts.officer_count, 0) as officer_count, wc.name as water_channel_name, wc.address as water_channel_address").
			Joins("LEFT JOIN (SELECT water_channel_door_id, COUNT(*) as cctv_count FROM water_channel_devices WHERE category = 'cctv' GROUP BY water_channel_door_id) as device_counts ON water_channel_doors.external_id = device_counts.water_channel_door_id").
			Joins("LEFT JOIN (SELECT water_channel_door_id, COUNT(*) as officer_count FROM water_channel_officers GROUP BY water_channel_door_id) as officer_counts ON water_channel_doors.external_id = officer_counts.water_channel_door_id").
			Joins("LEFT JOIN water_channels wc ON water_channel_doors.water_channel_id = wc.external_id").
			Find(&waterChannels).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListWaterChannelDoorWithCCTVAndOfficerRes{
			WaterChannels: waterChannels,
		}, nil
	}
}
