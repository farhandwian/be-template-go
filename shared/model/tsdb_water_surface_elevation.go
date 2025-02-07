package model

import "time"

type WaterSurfaceElevationData struct {
	WaterLevel         float32   `gorm:"column:water_level"`
	Status             bool      `gorm:"column:status"`
	WaterChannelDoorID int       `gorm:"column:water_channel_door_id"`
	Timestamp          time.Time `gorm:"column:timestamp;not null"`
}

func (WaterSurfaceElevationData) TableName() string {
	return "water_surface_elevation_data"
}
