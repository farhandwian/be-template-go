package model

import "time"

type IKU struct {
	ID         *string `json:"id"`
	Nama       *string `json:"nama"`
	Periode    *string `json:"periode"`
	Target     *string `json:"target"`
	ExternalID *string `json:"external_id"` // nilai nya berupa id antara 3 tabel yang sesuai dengan typenya
	Type       *string `json:"type"`        // PenetapanKonteksRisikoOperasionalInspektoratDaerah | PenetapanKonteksRisikoStrategisInspektoratDaerah |PenetapanKonteksRisikoStrategisPemda
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
