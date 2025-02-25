package model

import "time"

// Form5

type DaftarRisikoPrioritas struct {
	ID                 *string    `json:"id"`
	NamaPemda          *string    `json:"nama_pemda"`
	TahunPenilaian     *time.Time `json:"tahun_penilaian"`
	TujuanStrategis    *string    `json:"tujuan_strategis"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`
	// Taken from form 3a, 3b, 3c
	HasilAnalisisRisikoID *string `json:"-"`
	RisikoPrioritas       *string `json:"risiko_prioritas"`
	KodeRisiko            *string `json:"kode_resiko"`
	KategoriRisiko        *string `json:"kategori_risiko"`
	PemilikRisiko         *string `json:"pemilik_risiko"`
	PenyebabRisiko        *string `json:"penyebab_risiko"`
	DampakRisiko          *string `json:"dampak_risiko"`
}
