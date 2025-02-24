package model

import "time"

// Form 6

type PenilaianKegiatanPengendalian struct {
	ID                            *string   `json:"id"`
	NamaPemda                     *string   `json:"nama_pemda"`
	TahunPenilaian                *string   `json:"tahun_penilaian"`
	SPIPId                        *string   `json:"-"`
	SPIPName                      *string   `json:"spip_name"` // references spip.name
	KondisiLingkunganPengendalian *string   `json:"kondisi_lingkungan_pengendalian"`
	RencanaTindakPerbaikan        *string   `json:"rencana_tindak_perbaikan"`
	PenanggungJawab               *string   `json:"penanggung_jawab"`
	TargetWaktuPenyelesaian       *string   `json:"target_waktu_penyelesaian"`
	CreatedAt                     time.Time `json:"created_at"`
	UpdatedAt                     time.Time `json:"updated_at"`
}
