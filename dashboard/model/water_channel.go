package model

import (
	"gorm.io/datatypes"
	"time"
)

type WaterChannel struct {
	ID               string         `gorm:"type:char(36);primary_key;"`
	ExternalID       int            `gorm:"uniqueIndex;column:external_id"`
	Name             string         `gorm:"type:varchar(255)"`
	Address          string         `gorm:"type:text"`
	Photos           datatypes.JSON `gorm:"type:json"`
	IrrigationAreaID int            `gorm:"column:irrigation_area_id"`
	ChannelID        *int           `gorm:"column:channel_id"`
	SMOPIChannelID   int            `gorm:"column:smopi_channel_id"`
	ApiCreatedAt     *time.Time
	ApiUpdatedAt     *time.Time
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}
