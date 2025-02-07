package model

import "time"

type RainfallPost struct {
	ID         string `gorm:"primaryKey; length:36"`
	ExternalID int64  `gorm:"unique; not null"`
	Name       string
	Type       string
	PostVendor string
	Officer    string
	Elevation  float64
	Latitude   float64
	Longitude  float64
	City       string
	Vendor     string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
