package model

import "time"

// Form 2B

type PenetapanKonteksRisikoStrategisInspektoratDaerah struct {
	ID                 *string    `json:"id"`
	NamaPemda          *string    `json:"nama_pemda"`
	TahunPenilaian     *time.Time `json:"tahun_penilaian"`
	Periode            *string    `json:"periode"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`
	OPD                *string    `json:"opd"` // references opd.nama
	TujuanStrategis    *string    `json:"tujuan_strategis"`
	SasaranStrategis   *string    `json:"sasaran_strategis"`
	InformasiLain      *string    `json:"informasi_lain"`
	Penilaian          *string    `json:"penilaian"`
	NamaYBS            *string    `json:"nama_ybs"`
}
