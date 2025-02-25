package model

import "time"

// form 8

type PengkomunikasianPengendalian struct {
	ID                            *string    `json:"id"`
	NamaPemda                     *string    `json:"nama_pemda"`
	TahunPenilaian                *time.Time `json:"tahun_penilaian"`
	TujuanStrategis               *string    `json:"tujuan_strategis"`
	UrusanPemerintahan            *string    `json:"urusan_pemerintahan"`
	KegiatanPengendalian          *string    `json:"kegiatan_pengendalian"`
	PenilaiKegiatanPengendalianID *string    `json:"-"`
	MediaSaranaPengkomunikasian   *string    `json:"media_sarana_pengkomunikasian"`
	PenyediaInformasi             *string    `json:"penyedia_informasi"`
	PenerimaInformasi             *string    `json:"penerima_informasi"`
	RencanaPelaksanaan            *string    `json:"rencana_pelaksanaan"`
	RealisasiPelaksanaan          *string    `json:"realiasi_pelaksanaan"`
	Keterangan                    *string    `json:"keterangan"`
}
