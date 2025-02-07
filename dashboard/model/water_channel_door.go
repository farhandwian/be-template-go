package model

import (
	"time"

	"gorm.io/datatypes"
)

type WaterChannelDoor struct {
	ID                  string `gorm:"type:char(36);primary_key;"`
	ExternalID          int    `gorm:"uniqueIndex;column:external_id"`
	Name                string `gorm:"type:varchar(255);"`
	Latitude            string
	Longitude           string
	Address             string
	IPGateway           string
	Photos              datatypes.JSON `gorm:"type:json"`
	Width               float64
	Cc                  string
	WaterChannelID      int
	SmopiChannelID      int
	ForecastBuildingID  *string
	AreaSize            string
	DebitRequirement    string
	DebitActual         string
	DebitRecommendation string
	IsContainTrash      bool
	ApiCreatedAt        *time.Time
	ApiUpdatedAt        *time.Time
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
}
