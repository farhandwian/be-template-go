package gateway

import (
	"context"
	"shared/core"
	"shared/model"

	"gorm.io/gorm"
)

type GetListWaterGateLevelReq struct {
}

type GetListWaterGateLevelResp struct {
	WaterGates []model.WaterGateData
}

type GetListWaterGateLevelGateway = core.ActionHandler[GetListWaterGateLevelReq, GetListWaterGateLevelResp]

func ImplGetListWaterGateLevelGateway(tsDB *gorm.DB) GetListWaterGateLevelGateway {
	return func(ctx context.Context, request GetListWaterGateLevelReq) (*GetListWaterGateLevelResp, error) {
		var waterGates []model.WaterGateData

		query := tsDB.Raw(`
				WITH latest_gate AS (
			SELECT DISTINCT ON (water_channel_door_id, device_id)
				   water_channel_door_id,
				   device_id,
				   gate_level,
				   timestamp
			FROM water_gates
			ORDER BY water_channel_door_id, device_id, timestamp DESC
		)
		SELECT water_channel_door_id, device_id, gate_level
		FROM latest_gate;
		`)

		err := query.Find(&waterGates).Error
		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetListWaterGateLevelResp{
			WaterGates: waterGates,
		}, nil
	}
}
