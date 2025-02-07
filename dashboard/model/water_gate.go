package model

import "time"

type WaterGate struct {
	ID                 string `gorm:"type:char(36);primary_key;"`
	SecurityRelay      int
	GroupRelay         int
	GateLevel          float64
	WaterChannelDoorID int
	DeviceID           int
	Status             bool
	Timestamp          time.Time `gorm:"primaryKey;index"`
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}
