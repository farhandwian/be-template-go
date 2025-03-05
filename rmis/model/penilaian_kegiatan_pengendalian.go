package model

import (
	sharedModel "shared/model"
	"time"
)

// Form 6

type PenilaianKegiatanPengendalian struct {
	ID                            *string            `json:"id"`
	NamaPemda                     *string            `json:"nama_pemda"`
	TahunPenilaian                *string            `json:"tahun_penilaian"`
	SpipId                        *string            `json:"-"`
	KondisiLingkunganPengendalian *string            `json:"kondisi_lingkungan_pengendalian"`
	RencanaTindakPerbaikan        *string            `json:"rencana_tindak_perbaikan"`
	PenanggungJawab               *string            `json:"penanggung_jawab"`
	TargetWaktuPenyelesaian       *string            `json:"target_waktu_penyelesaian"`
	Status                        sharedModel.Status `json:"status"`
	CreatedAt                     time.Time          `json:"created_at"`
	UpdatedAt                     time.Time          `json:"updated_at"`
}
