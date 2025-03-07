package model

import (
	sharedModel "shared/model"
	"time"
)

// Form5

type DaftarRisikoPrioritas struct {
	ID               *string `json:"id"`
	TipeIdentifikasi *string `json:"tipe_identifikasi" gorm:"type:VARCHAR(50)"` // "strategis_pemda", "operasional_opd", or "strategis_renstra_opd"
	IdentifikasiID   *string `json:"identifikasi_id" gorm:"type:VARCHAR(255)"`

	TipePenetapanKonteks *string `json:"tipe_penetapan_konteks" gorm:"type:VARCHAR(50)"` // "strategis_pemda", "operasional", or "strategis_renstra_opd"
	PenetapanKonteksID   *string `json:"penetapan_konteks_id" gorm:"type:VARCHAR(255)"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type DaftarRisikoPrioritasResponse struct {
	ID                                     *string `json:"id"`
	PenetapanKonteksRisikoStrategisPemdaID *string `json:"-" gorm:"type:VARCHAR(255)"`
	HasilAnalisisRisikoID                  *string `json:"hasil_analisis_risiko_id"`

	RisikoPrioritas *string `json:"risiko_prioritas"`
	KodeRisiko      *string `json:"kode_resiko"`
	KategoriRisiko  *string `json:"kategori_risiko"`
	PemilikRisiko   *string `json:"pemilik_risiko"`
	PenyebabRisiko  *string `json:"penyebab_risiko"`
	DampakRisiko    *string `json:"dampak_risiko"`

	// Fields from JOIN
	NamaPemda          *string    `json:"nama_pemda"`
	Tahun              *time.Time `json:"tahun"`
	Periode            *string    `json:"periode"`
	PenetapanKonteks   *string    `json:"penetapan_konteks"`
	UrusanPemerintahan *string    `json:"urusan_pemerintahan"`
	SkalaRisiko        *int       `json:"skala_risiko"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
