package model

import (
	"time"
)

type JDIH struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Title       string
	PublishedAt time.Time
	Status      JDIHStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type JDIHStatus string

const (
	JDIHStatusRevoked    JDIHStatus = "revoked"    // dicabut
	JDIHStatusApplicable JDIHStatus = "applicable" // berlaku
)
