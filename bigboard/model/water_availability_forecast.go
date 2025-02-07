package model

import "time"

type WaterAvailabilityForecast struct {
	ID                    string `gorm:"type:char(36);primary_key;"`
	RequiredWater         float64
	AverageAvailableWater float64
	MaxAvailableWater     float64
	MinAvailableWater     float64
	Date                  string `gorm:"unique"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
