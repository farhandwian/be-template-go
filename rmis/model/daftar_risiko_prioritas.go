package model

import (
	sharedModel "shared/model"
	"time"
)

// Form5

type DaftarRisikoPrioritas struct {
	ID                                     *string `json:"id"`
	PenetapanKonteksRisikoStrategisPemdaID *string `json:"-" gorm:"type:VARCHAR(255)"`
	// Taken from form 3a, 3b, 3c
	HasilAnalisisRisikoID *string `json:"-"`

	IndeksPeringkatPrioritasID *string `json:"-"`
	RisikoPrioritas            *string `json:"risiko_prioritas"`
	KodeRisiko                 *string `json:"kode_resiko"`
	KategoriRisiko             *string `json:"kategori_risiko"`
	PemilikRisiko              *string `json:"pemilik_risiko"`
	PenyebabRisiko             *string `json:"penyebab_risiko"`
	DampakRisiko               *string `json:"dampak_risiko"`

	NamaPemda        *string            `json:"nama_pemda"`
	Tahun            *time.Time         `json:"tahun"`
	Periode          *string            `json:"periode"`
	PenetapanKonteks *string            `json:"penetapan_konteks"`
	UrusanPemerintah *string            `json:"urusan_pemerintah"`
	Status           sharedModel.Status `json:"status"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}
