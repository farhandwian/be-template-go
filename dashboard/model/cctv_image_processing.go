package model

import "time"

type CctvImageProcessing struct {
	Timestamp          time.Time `json:"timestamp"`
	WaterChannelDoorId int       `json:"water_channel_door_id"`
	IP                 string    `json:"ip"`
	DetectedObject     string    `json:"detected_object"`
	ImagePath          string    `json:"image_path"`
}
