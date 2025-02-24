package model

import (
	"time"

	"gorm.io/datatypes"
)

// Form1B

type SimpulanKondisiKelemahanLingkungan struct {
	ID                 *string         `json:"id"`
	NamaPemda          *string         `json:"nama_pemda"`
	TahunPenilaian     *string         `json:"tahun_penilaian"`
	UrusanPemerintahan *string         `json:"urusan_pemerintahan"`
	SumberData         *string         `json:"sumber_data"`
	UraianKelamahan    *datatypes.JSON `json:"uraian_kelemahan"`
	Klasifikasi        *string         `json:"klasifikasi"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}
