package gateway

import (
	"context"
	"shared/core"
	"time"

	"gorm.io/gorm"
)

type GetDebitReq struct {
	WaterChannelDoorID int
	MinTime            time.Time `json:"min_time"`
	MaxTime            time.Time `json:"max_time"`
}

type ActualDebit struct {
	ActualDebit float64   `json:"debit"`
	Time        time.Time `json:"time"`
}

type GetDebitRes struct {
	Debits []ActualDebit
}

type GetDebit = core.ActionHandler[GetDebitReq, GetDebitRes]

func ImplGetDebit(tsDB *gorm.DB) GetDebit {
	return func(ctx context.Context, req GetDebitReq) (*GetDebitRes, error) {

		var debits []ActualDebit

		timeRange := req.MaxTime.Sub(req.MinTime).Hours() / 24

		bucket := "5 minutes"
		if timeRange > 7 {
			bucket = "1 day"
		} else if timeRange > 1 {
			bucket = "1 hour"
		}

		query := `
			SELECT
			time_bucket(?, timestamp) AS time,
			AVG(actual_debit) AS actual_debit
			FROM actual_debit_data
			WHERE water_channel_door_id = ? 
			AND timestamp BETWEEN ? AND ?
			GROUP BY time ORDER BY time ASC;
		`
		err := tsDB.Raw(query, bucket, req.WaterChannelDoorID, req.MinTime, req.MaxTime).Scan(&debits).Error

		if err != nil {
			return nil, core.NewInternalServerError(err)
		}

		return &GetDebitRes{
			Debits: debits,
		}, nil
	}
}
