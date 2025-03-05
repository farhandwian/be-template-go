package model

import (
	sharedModel "shared/model"
	"time"
)

// form 9

type RancanganPemantauan struct {
	ID                       *string           `json:"id"`
	PenilaianRisikoID        *string           `json:"penilaian_risiko_id"`
	MetodePemantauan         *MetodePemantauan `json:"metode_pemantuan"`
	PenanggungJawab          *string           `json:"penanggung_jawab"`
	RencanaWaktuPemantauan   *string           `json:"rencana_waktu_pemantauan"`
	RealisasiWaktuPemantauan *string           `json:"realisasi_waktu_pemantauan"`
	Keterangan               *string           `json:"keterangan"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type RancanganPemantauanResponse struct {
	ID                       *string           `json:"id"`
	PenilaianRisikoID        *string           `json:"penilaian_risiko_id"`
	MetodePemantauan         *MetodePemantauan `json:"metode_pemantuan"`
	PenanggungJawab          *string           `json:"penanggung_jawab"`
	RencanaWaktuPemantauan   *string           `json:"rencana_waktu_pemantauan"`
	RealisasiWaktuPemantauan *string           `json:"realisasi_waktu_pemantauan"`
	Keterangan               *string           `json:"keterangan"`

	RencanaTindakPengendalian *string `json:"rencana_tindak_pengendalian"`
	NamaPemda                 *string `json:"nama_pemda"`
	TahunPenilaian            *string `json:"tahun_penilaian"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type MetodePemantauan string

const (
	MetodePemantauanPersiapanDanLaporan MetodePemantauan = "Konfirmasi persiapan dan laporan pelaksanaan kegiatan"
	MetodePemantauanPelaksanakan        MetodePemantauan = "Konfirmasi pelaksanaan Laporan pelaksanaan kegiatan"
)
