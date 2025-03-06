package model

import (
	sharedModel "shared/model"
	"time"
)

// Form 2B

type PenetapanKonteksRisikoStrategisRenstraOPD struct {
	ID                 *string    `json:"id"`
	NamaPemda          *string    `json:"nama_pemda"`
	TahunPenilaian     *time.Time `json:"tahun_penilaian"`
	Periode            *string    `json:"periode"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`
	OpdID              *string    `json:"opd_id"`
	SumberData         *string    `json:"sumber_data"`
	TujuanStrategis    *string    `json:"tujuan_strategis"`
	SasaranStrategis   *string    `json:"sasaran_strategis"`
	InformasiLain      *string    `json:"informasi_lain"`
	PenetapanTujuan    *string    `json:"penetapan_tujuan"`
	PenetapanSasaran   *string    `json:"penetapan_sasaran"`
	PenetapanIku       *string    `json:"penetapan_iku"`

	Status sharedModel.Status `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDResponse struct {
	ID                 *string `json:"id"`
	NamaPemda          *string `json:"nama_pemda"`
	TahunPenilaian     *string `json:"tahun_penilaian"`
	Periode            *string `json:"periode"`
	UrusanPemerintahan *string `json:"urusan_pemerintahan"`
	OpdID              *string `json:"opd_id"`
	OpdNama            *string `json:"opd_nama"`
	SumberData         *string `json:"sumber_data"`
	TujuanStrategis    *string `json:"tujuan_strategis"`
	SasaranStrategis   *string `json:"sasaran_strategis"`
	InformasiLain      *string `json:"informasi_lain"`
	PenetapanTujuan    *string `json:"penetapan_tujuan"`
	PenetapanSasaran   *string `json:"penetapan_sasaran"`
	PenetapanIku       *string `json:"penetapan_iku"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
type PenetapanKonteksRisikoStrategisRenstraOPDGet struct {
	PenetapanKonteksRisikoStrategisRenstraOPD PenetapanKonteksRisikoStrategisRenstraOPDResponse `json:"penetapan_konteks_risiko_strategis_renstra_opd"`
	IKUs                                      []IKU                                             `json:"ikus"`
	OPD                                       OPD                                               `json:"opd"`
}
