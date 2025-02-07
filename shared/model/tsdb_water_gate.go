package model

import "time"

type WaterGateData struct {
	Timestamp          time.Time
	WaterChannelDoorID int
	DeviceID           int
	GateLevel          float32
	SecurityRelay      bool
	Run                bool
	Sensor             bool
	Status             bool
}

func (WaterGateData) TableName() string {
	return "water_gate_data"
}
