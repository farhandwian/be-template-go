package model

import "time"

type WaterChannelDeviceCategory string

const (
	CategoryController WaterChannelDeviceCategory = "controller"
	CategoryTelemetri  WaterChannelDeviceCategory = "telemetri"
	CategoryCCTV       WaterChannelDeviceCategory = "cctv"
)

type WaterChannelDevice struct {
	ID                 string                     `gorm:"type:char(36);primary_key;"`
	ExternalID         int                        `gorm:"uniqueIndex;column:external_id"`
	Category           WaterChannelDeviceCategory `gorm:"type:varchar(255);"`
	Name               string                     `gorm:"type:varchar(255);"`
	IPAddress          string
	GroupRelay         int
	Type               string
	FullTime           float64
	MaxHeightSensor    float32
	UpperLimit         *int
	LowerLimit         *int
	MeasurementScale   int
	WaterChannelDoorID int
	ApiCreatedAt       *time.Time
	ApiUpdatedAt       *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	DetectedObject     string
}
