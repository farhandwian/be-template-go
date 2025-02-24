package model

import (
	"time"

	"gorm.io/datatypes"
)

// Form 2C

type PenetapanKonteksRisikoOperasionalInspektoratDaerah struct {
	ID                 *string         `json:"id"`
	NamaPemda          *string         `json:"nama_pemda"`
	TahunPenilaian     *time.Time      `json:"tahun_penilaian"`
	Periode            *string         `json:"periode"`
	UrusanPemerintahan *string         `json:"urusan_pemerintahan"`
	OPD                *string         `json:"opd"` // references opd.nama
	TujuanStrategis    *string         `json:"tujuan_strategis"`
	ProgramInspektorat *datatypes.JSON `json:"program_inspektorat"`
	InformasiLain      *string         `json:"informasi_lain"`
	NamaYBS            *string         `json:"nama_ybs"`
}
