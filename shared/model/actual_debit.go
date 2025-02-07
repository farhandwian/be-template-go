package model

import "time"

type ActualDebitData struct {
	WaterChannelDoorID int
	ActualDebit        float64
	Timestamp          time.Time
}

func (a *ActualDebitData) TableName() string {
	return "actual_debit_data"
}
