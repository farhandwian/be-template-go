package model

import "time"

type WaterChannelOfficer struct {
	ID                 string `gorm:"type:char(36);primary_key;"`
	ExternalID         int    `gorm:"type:int;uniqueIndex"`
	Task               string
	Photo              string
	Name               string
	PhoneNumber        string
	Address            string
	WorkRegion         string
	WaterChannelDoorID int
	APICreatedAt       *time.Time
	APIUpdatedAt       *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}
