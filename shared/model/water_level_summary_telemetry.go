package model

type HydrologyWaterLevelSummaryTelemetry struct {
	ID               string `gorm:"primaryKey;length:36"`
	TMAMaximum       float64
	TMAMinimum       float64
	LatestUpdate     string
	WaterLevelPostID int64
}
