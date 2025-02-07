package model

import "time"

type Asset struct {
	ID        string      `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name"`
	PIC       string      `json:"pic"`
	Location  string      `json:"location"`
	Status    AssetStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type AssetStatus string

const (
	AssetStatusBroken AssetStatus = "broken"
	AssetStatusActive AssetStatus = "active"
)
