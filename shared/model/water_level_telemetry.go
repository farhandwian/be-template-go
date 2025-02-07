package model

import "time"

type HydrologyWaterLevelTelemetry struct {
	ID               string `gorm:"primaryKey;length:36"`
	WaterLevelPostID int64
	Battery          float64
	WaterLevel       float64
	Sampling         string
	Timestamp        time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
