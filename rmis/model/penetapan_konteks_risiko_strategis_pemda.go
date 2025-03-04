package model

import (
	sharedModel "shared/model"
	"time"
)

// Form 2A

type PenetapanKonteksRisikoStrategisPemda struct {
	ID                              *string            `json:"id" gorm:"primaryKey"`
	NamaPemda                       *string            `json:"nama_pemda"`
	TahunPenilaian                  *time.Time         `json:"tahun_penilaian"`
	Periode                         *string            `json:"periode"`
	SumberData                      *string            `json:"sumber_data"`
	UrusanPemerintahan              *string            `json:"urusan_pemerintahan"`
	TujuanStrategis                 *string            `json:"tujuan_strategis"`
	SasaranStrategis                *string            `json:"sasaran_strategis"`
	PrioritasPembangunan            *string            `json:"prioritas_pembangunan"`
	IkuSasaran                      *string            `json:"iku_sasaran"`
	PenetapanKonteksRisikoStrategis *string            `json:"penetapan_konteks_resiko_strategis"`
	NamaDinas                       *string            `json:"nama_dinas"`
	PenetapanTujuan                 *string            `json:"penetapan_tujuan"`
	PenetapanSasaran                *string            `json:"penetapan_sasaran"`
	PenetapanIku                    *string            `json:"penetapan_iku"`
	Status                          sharedModel.Status `json:"status"`
	CreatedAt                       time.Time
	UpdatedAt                       time.Time
}
