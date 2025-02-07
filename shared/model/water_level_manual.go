package model

import "time"

type HydrologyWaterLevelManual struct {
	ID               string `gorm:"primaryKey;size:36"`
	WaterLevelPostID int64
	TMA              float64
	Sampling         string
	Timestamp        time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
