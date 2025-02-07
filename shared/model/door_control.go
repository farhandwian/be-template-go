package model

import (
	iammodel "iam/model"
	"time"
)

type DoorControlID string

type DoorControl struct {
	ID                 DoorControlID     `json:"id"`
	Date               time.Time         `json:"date"`
	WaterChannelDoorID int               `json:"water_channel_door_id"`
	DeviceID           int               `json:"device_id"`
	DoorName           string            `json:"door_name"`
	OpenTarget         float32           `json:"open_target"`
	Reason             string            `json:"reason"`
	ErrorMessage       string            `json:"error_message"`
	Status             DoorControlStatus `json:"status"`
	OfficerId          iammodel.UserID   `json:"officer_id"`
	OfficerName        string            `json:"officer_name"`
}

type DoorControlStatus string

const (
	StatusMenunggu   DoorControlStatus = "Menunggu"
	StatusGagal      DoorControlStatus = "Gagal"
	StatusDieksekusi DoorControlStatus = "Dieksekusi"
	StatusDibatalkan DoorControlStatus = "Dibatalkan"
)

type DoorControlPayload struct {
	DoorControlID      DoorControlID
	OpenTarget         float32
	WaterChannelDoorID int
	DeviceID           int
	IPAddress          string
}
