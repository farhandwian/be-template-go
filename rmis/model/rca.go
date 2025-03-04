package model

import (
	"encoding/json"
	"fmt"
	sharedModel "shared/model"
	"time"

	"gorm.io/datatypes"
)

// Root Cause Analysis
type Rca struct {
	ID                                 *string            `json:"id" gorm:"type:VARCHAR(191)"`
	PemilikRisiko                      *string            `json:"pemilik_risiko" gorm:"type:VARCHAR(255)"`
	TahunPenilaian                     *time.Time         `json:"tahun_penilaian"`
	IdentifikasiRisikoStrategisPemdaID *string            `json:"test" gorm:"type:VARCHAR(191);not null"` // Foreign key field
	IdentifikasiRisikoStrategisPemda   *string            `json:"identifikasi_risiko_strategis_pemda"`
	Why                                *datatypes.JSON    `json:"why"` // JSON array field
	AkarPenyebab                       *string            `json:"akar_penyebab" gorm:"type:VARCHAR(255)"`
	PenyebabRisikoID                   *string            `json:"-" gorm:"type:VARCHAR(191);not null"` // Foreign key field
	PenyebabRisiko                     *PenyebabRisiko    `json:"penyebab_risiko"`
	KegiatanPengendalian               *string            `json:"kegiatan_pengendalian" gorm:"type:VARCHAR(255)"`
	Status                             sharedModel.Status `json:"status"`
	IdentifikasiRisikoUraianRisiko     *string            `json:"identifikasi_risiko_uraian_risiko"` // Add this field
	PenyebabRisikoNama                 *string            `json:"penyebab_risiko_nama"`
	CreatedAt                          time.Time          `json:"created_at"`
	UpdatedAt                          time.Time          `json:"updated_at"`
}

func (rca *Rca) SetAkarPenyebab() error {
	if rca.Why == nil {
		return fmt.Errorf("why field is empty")
	}

	var whyArray []string
	if err := json.Unmarshal(*rca.Why, &whyArray); err != nil {
		return fmt.Errorf("failed to unmarshal why field: %v", err)
	}

	// get the last element
	if len(whyArray) > 0 {
		rca.AkarPenyebab = &whyArray[len(whyArray)-1]
	} else {
		return fmt.Errorf("why field is empty")
	}

	return nil
}
