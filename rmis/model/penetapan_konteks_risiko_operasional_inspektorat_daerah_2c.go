package model

import (
	"gorm.io/datatypes"
)

// Form 2C

// Penetapan Konteks Risiko Operasional OPD. uraian / sumber data dari RKA, pj nya oleh tim MR

type PenetapanKonteksRisikoOperasional struct {
	ID                           *string         `json:"id"`
	NamaPemda                    *string         `json:"nama_pemda"`
	TahunPenilaian               *string         `json:"tahun_penilaian"`
	Periode                      *string         `json:"periode"`
	UrusanPemerintahan           *string         `json:"urusan_pemerintahan"`
	OPDID                        *string         `json:"opd_id"` // references opd
	TujuanStrategis              *string         `json:"tujuan_strategis"`
	ProgramInspektorat           *datatypes.JSON `json:"program_inspektorat"`
	InformasiLain                *string         `json:"informasi_lain"`
	KegiatanDanIndikatorKeluaran *string         `json:"kegiatan_dan_indikator_keluaran"` // Kegiatan,dan indikator keluaran yang akan dilakukan penilaian risiko
	NamaYBS                      *string         `json:"nama_ybs"`
}

type PenetapanKonteksRisikoOperasionalGet struct {
	PenetapanKonteksRisikoOperasional PenetapanKonteksRisikoOperasional `json:"penetapan_konteks_risiko_operasional"`
	IKUs                              []IKU                             `json:"ikus"`
	OPD                               OPD                               `json:"opd"`
}
