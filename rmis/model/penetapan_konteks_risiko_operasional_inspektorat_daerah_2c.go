package model

import (
	sharedModel "shared/model"
	"time"

	"gorm.io/datatypes"
)

// Form 2C

// Penetapan Konteks Risiko Operasional OPD. uraian / sumber data dari RKA, pj nya oleh tim MR

type PenetapanKonteksRisikoOperasional struct {
	ID                        *string         `json:"id"`
	NamaPemda                 *string         `json:"nama_pemda"`
	TahunPenilaian            *time.Time      `json:"tahun_penilaian"`
	Periode                   *string         `json:"periode"`
	SumberData                *string         `json:"sumber_data"`
	UrusanPemerintahan        *string         `json:"urusan_pemerintahan"`
	OpdID                     *string         `json:"opd_id"` // references opd
	TujuanStrategis           *string         `json:"tujuan_strategis"`
	KegiatanUtama             *string         `json:"kegiatan_utama"`
	InformasiLain             *string         `json:"informasi_lain"`
	KeluaranAtauHasilKegiatan *datatypes.JSON `json:"keluaran_atau_hasil_kegiatan"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type PenetapanKonteksRisikoOperasionalResponse struct {
	ID                        *string         `json:"id"`
	NamaPemda                 *string         `json:"nama_pemda"`
	TahunPenilaian            *time.Time      `json:"tahun_penilaian"`
	Periode                   *string         `json:"periode"`
	SumberData                *string         `json:"sumber_data"`
	UrusanPemerintahan        *string         `json:"urusan_pemerintahan"`
	OpdID                     *string         `json:"opd_id"` // references opd
	TujuanStrategis           *string         `json:"tujuan_strategis"`
	KegiatanUtama             *string         `json:"kegiatan_utama"`
	InformasiLain             *string         `json:"informasi_lain"`
	KeluaranAtauHasilKegiatan *datatypes.JSON `json:"keluaran_atau_hasil_kegiatan"`

	OpdNama   *string            `json:"opd_nama"`
	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type KeluaranAtauHasilKegiatan struct {
	NamaKegiatan  *string `json:"nama_kegiatan"`
	JumlahPeserta *int    `json:"jumlah_peserta"`
}
type PenetapanKonteksRisikoOperasionalGet struct {
	PenetapanKonteksRisikoOperasional PenetapanKonteksRisikoOperasional `json:"penetapan_konteks_risiko_operasional"`
	IKUs                              []IKU                             `json:"ikus"`
	OPD                               OPD                               `json:"opd"`
}
