package gateway

import (
	"context"
	"fmt"
	"shared/core"
	"shared/model"
	"strings"

	"gorm.io/gorm"
)

type GetLatestWaterGatesByDoorIDReq struct {
	WaterChannelDoorID int
	DeviceIDs          []int
}

type GetLatestWaterGatesByDoorIDRes struct {
	WaterGates []model.WaterGateData
}

type GetLatestWaterGatesByDoorID = core.ActionHandler[GetLatestWaterGatesByDoorIDReq, GetLatestWaterGatesByDoorIDRes]

func ImplGetLatestWaterGatesByDoorID(db *gorm.DB) GetLatestWaterGatesByDoorID {
	return func(ctx context.Context, req GetLatestWaterGatesByDoorIDReq) (*GetLatestWaterGatesByDoorIDRes, error) {
		var waterGates []model.WaterGateData

		// Buat placeholder untuk IN clause
		deviceIDPlaceholders := make([]string, len(req.DeviceIDs))
		deviceIDValues := make([]interface{}, len(req.DeviceIDs)+1) // +1 untuk water_channel_door_id

		deviceIDValues[0] = req.WaterChannelDoorID

		for i := range req.DeviceIDs {
			deviceIDPlaceholders[i] = "?"
			deviceIDValues[i+1] = req.DeviceIDs[i]
		}

		query := fmt.Sprintf(`
			SELECT DISTINCT ON (device_id) *
			FROM water_gate_data
			WHERE water_channel_door_id = ?
			AND device_id IN (%s)
			AND timestamp >= NOW() - INTERVAL '5 minutes'
			ORDER BY device_id, timestamp DESC
		`, strings.Join(deviceIDPlaceholders, ","))

		if err := db.Raw(query, deviceIDValues...).
			Find(&waterGates).Error; err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetLatestWaterGatesByDoorIDRes{
			WaterGates: waterGates,
		}, nil
	}
}
