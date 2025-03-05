package model

import "time"

type PenilaianRisiko struct {
	ID                        *string   `json:"id"`
	DaftarRisikoPrioritasID   *string   `json:"-"`
	RisikoPrioritas           *string   `json:"risiko_prioritas"`
	KodeRisiko                *string   `json:"kode_resiko"`
	UraianPengendalian        *string   `json:"uraian_pengendalian"`
	CelahPengendalian         *string   `json:"celah_pengendalian"`
	RencanaTindakPengendalian *string   `json:"rencana_tindak_pengendalian"`
	PemilikPenanggungJawab    *string   `json:"pemilik_penanggung_jawab"`
	TargetWaktuPenyelesaian   *string   `json:"target_waktu_penyelesaian"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
}

type PenilaianRisikoResponse struct {
	ID                        *string    `json:"id"`
	NamaPemda                 *string    `json:"nama_pemda"`
	TahunPenilaian            *time.Time `json:"tahun_penilaian"`
	DaftarRisikoPrioritasID   *string    `json:"-"`
	RisikoPrioritas           *string    `json:"risiko_prioritas"`
	KodeRisiko                *string    `json:"kode_resiko"`
	UraianPengendalian        *string    `json:"uraian_pengendalian"`
	CelahPengendalian         *string    `json:"celah_pengendalian"`
	RencanaTindakPengendalian *string    `json:"rencana_tindak_pengendalian"`
	PemilikPenanggungJawab    *string    `json:"pemilik_penanggung_jawab"`
	TargetWaktuPenyelesaian   *string    `json:"target_waktu_penyelesaian"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at"`
}
