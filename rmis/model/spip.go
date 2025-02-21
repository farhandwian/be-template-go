package model

import "time"

// Elemen Sub Unsur SPIP

type Spip struct {
	ID        *string   `json:"id"`
	Nama      *string   `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
