package gateway

import (
	"context"
	"shared/core"
	"time"

	"gorm.io/gorm"
)

type GetGateReq struct {
	WaterChannelDoorID int
	DeviceID           int
	MinTime            time.Time
	MaxTime            time.Time
}

type WaterGate struct {
	GateLevel          float64   `gorm:"column:gate_level"`
	WaterChannelDoorID int       `gorm:"column:water_channel_door_id"`
	DeviceID           int       `gorm:"column:device_id"`
	Time               time.Time `gorm:"column:time"`
}

type GetGateRes struct {
	DeviceID int
	Gates    []WaterGate
}

type GetGate = core.ActionHandler[GetGateReq, GetGateRes]

func ImplGetGate(tsDB *gorm.DB) GetGate {
	return func(ctx context.Context, req GetGateReq) (*GetGateRes, error) {
		var gates []WaterGate

		query := `
			SELECT
				timestamp AS time,
				gate_level,
				device_id
			FROM
				water_gate_data
			WHERE
				water_channel_door_id = ?
				AND device_id = ?
				AND status = true
				AND timestamp BETWEEN ? AND ?
			order by
				time ASC;
		`
		err := tsDB.Raw(query, req.WaterChannelDoorID, req.DeviceID, req.MinTime, req.MaxTime).Scan(&gates).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetGateRes{
			DeviceID: req.DeviceID,
			Gates:    gates,
		}, nil
	}
}
