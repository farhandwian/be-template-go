package model

import (
	"iam/model"
	"time"
)

type DoorControlHistoryID string

type DoorControlHistory struct {
	ID                 DoorControlHistoryID   `json:"id"`
	Date               time.Time              `json:"date"`
	WaterChannelDoorID int                    `json:"water_channel_door_id"`
	DeviceID           int                    `json:"device_id"`
	DoorName           string                 `json:"door_name"`
	OpenCurrent        float32                `json:"open_current"`
	OpenTarget         float32                `json:"open_target"`
	Type               DoorControlHistoryType `json:"type"`
	Reason             string                 `json:"reason"`
	Status             DoorControlStatus      `json:"status"`
	OfficerName        string                 `json:"officer_name"`
	OfficerId          model.UserID           `json:"officer_id"`
	ErrorMessage       string                 `json:"error_message"`
}

type DoorControlHistoryType string

const (
	TypeTerjadwal DoorControlHistoryType = "Terjadwal"
	TypeLangsung  DoorControlHistoryType = "Langsung"
)
