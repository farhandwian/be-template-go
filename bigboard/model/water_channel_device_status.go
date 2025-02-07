package model

import "time"

type WaterChannelDoorDevice struct {
	ID                 string
	WaterChannelDoorID int
	DeviceID           int
	Status             bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
