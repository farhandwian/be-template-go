package model

import "time"

type KategoriRisiko struct {
	ID        *string   `json:"id"`
	Nama      *string   `json:"nama"`
	Kode      *string   `json:"kode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
