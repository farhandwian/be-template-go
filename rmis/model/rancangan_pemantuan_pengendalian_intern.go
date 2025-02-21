package model

import "time"

// form 9

type RancanganPemantauanPengendalianIntern struct {
	ID                   *string    `json:"id"`
	NamaPemda            *string    `json:"nama_pemda"`
	TahunPenilaian       *time.Time `json:"tahun_penilaian"`
	TujuanStrategis      *string    `json:"tujuan_strategis"`
	UrusanPemerintahan   *string    `json:"urusan_pemerintahan"`
	KegiatanPengendalian *string    `json:"kegiatan_pengendalian"`
	BentukPemantauan     *string    `json:"bentuk_pemantuan"`
	PenanggungJawab      *string    `json:"penanggung_jawab"`
	RencanaPenyelesaian  *string    `json:"rencana_penyelesaian"`
	RealisasiPelaksanaan *string    `json:"realisasi_pelaksanaan"`
	Keterangan           *string    `json:"keterangan"`
}
