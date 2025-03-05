package model

import (
	sharedModel "shared/model"
	"time"
)

// form 8

type PengkomunikasianPengendalian struct {
	ID                   *string            `json:"id"`
	PenilaianRisikoID    *string            `json:"penilaian_risiko_id"`
	MediaKomunikasi      *string            `json:"media_komunikasi"`
	PenyediaInformasi    *string            `json:"penyedia_informasi"`
	PenerimaInformasi    *string            `json:"penerima_informasi"`
	RencanaPelaksanaan   *string            `json:"rencana_pelaksanaan"`
	RealisasiPelaksanaan *string            `json:"realiasi_pelaksanaan"`
	Keterangan           *string            `json:"keterangan"`
	Status               sharedModel.Status `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PengkomunikasianPengendalianResponse struct {
	ID                        *string            `json:"id"`
	PenilaianRisikoID         *string            `json:"penilaian_risiko_id"`
	RencanaTindakPengendalian *string            `json:"rencana_tindak_pengendalian"`
	NamaPemda                 *string            `json:"nama_pemda"`
	TahunPenilaian            *string            `json:"tahun_penilaian"`
	MediaKomunikasi           *string            `json:"media_komunikasi"`
	PenyediaInformasi         *string            `json:"penyedia_informasi"`
	PenerimaInformasi         *string            `json:"penerima_informasi"`
	RencanaPelaksanaan        *string            `json:"rencana_pelaksanaan"`
	RealisasiPelaksanaan      *string            `json:"realiasi_pelaksanaan"`
	Keterangan                *string            `json:"keterangan"`
	Status                    sharedModel.Status `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
