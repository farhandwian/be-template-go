package model

import "time"

// form 8

type PengkomunikasikanPengendalian struct {
	ID                          *string    `json:"id"`
	NamaPemda                   *string    `json:"nama_pemda"`
	TahunPenilaian              *time.Time `json:"tahun_penilaian"`
	TujuanStrategis             *string    `json:"tujuan_strategis"`
	UrusanPemerintahan          *string    `json:"urusan_pemerintahan"`
	KegiatanPengendalian        *string    `json:"kegiatan_pengendalian"`
	MediaSaranaPengkomunikasian *string    `json:"media_sarana_pengkomunikasiian"`
	PenyediaInformasi           *string    `json:"penyedia_informasi"`
	PenerimaInformasi           *string    `json:"penerima_informasi"`
	RencanaPenyelesaian         *string    `json:"rencana_penyelesaian"`
	RealisasiPelaksanaan        *string    `json:"realiassi_pelaksanaan"`
	Keterangan                  *string    `json:"keterangan"`
}
