package model

import (
	"time"

	"gorm.io/datatypes"
)

type FuncName string
type SensitiveJobsID string

const (
	FuncTypeDoorControl       FuncName = "door_control"
	FuncTypeDoorControlSensor FuncName = "door_control_sensor"
)

type Status string

const (
	StatusCreated Status = "created"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
)

type SensitiveJobs struct {
	ID        SensitiveJobsID `gorm:"type:char(36);primary_key;"`
	FuncName  FuncName        `gorm:"type:varchar(255);"`
	Status    Status          `gorm:"type:varchar(255);"`
	Payload   datatypes.JSON  `gorm:"type:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
