package model

import "time"

type GeneralInfo struct {
	ID              string `gorm:"type:char(36);primary_key;"`
	ExternalID      int    `gorm:"uniqueIndex;column:external_id"`
	PlantingSeason  string
	PlantingPattern string
	PlantingArea    string
	RequiredWater   string
	ApiCreatedAt    time.Time
	ApiUpdatedAt    time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
