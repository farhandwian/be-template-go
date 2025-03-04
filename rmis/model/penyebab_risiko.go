package model

import "time"

type PenyebabRisiko struct {
	ID        *string   `json:"id" gorm:"primaryKey"`
	Nama      *string   `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
