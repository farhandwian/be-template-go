package model

import "time"

// Form 6

type PenilaianKegiatanPengendalian *struct {
	ID                            *string    `json:"id"`
	NamaPemda                     *string    `json:"nama_pemda"`
	TahunPenilaian                *time.Time `json:"tahun_penilaian"`
	SPIP                          *string    `json:"spip"` // references spip.name
	KondisiLingkunganPengendalian *string    `json:"kondisi_lingkungan_pengendalian"`
	RencanaTindakPerbaikan        *string    `json:"rencana_tindak_perbaikan"`
	PenanggungJawab               *string    `json:"penanggung_jawab"`
	TargetWaktuPenyelesaian       *string    `json:"target_waktu_penyelesaian"`
}
