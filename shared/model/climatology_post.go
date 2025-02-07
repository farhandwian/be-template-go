package model

import (
	"time"
)

type ClimatologyPost struct {
	ID         string `gorm:"primaryKey; length:36"`
	ExternalID int64  `gorm:"unique; not null"`
	Name       string
	Type       string
	PostVendor string
	Officer    string
	Elevation  float64
	Latitude   float64
	Longitude  float64
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
