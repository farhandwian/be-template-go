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
	ID                                 *string                           `json:"id"`
	PemilikRisiko                      *string                           `json:"pemilik_risiko"`
	TahunPenilaian                     *time.Time                        `json:"tahun_penilaian"`
	IdentifikasiRisikoStrategisPemdaID *string                           `json:"-"` // Foreign key field
	IdentifikasiRisikoStrategisPemda   *IdentifikasiRisikoStrategisPemda `json:"identifikasi_risiko_strategis_pemda" gorm:"foreignKey:IdentifikasiRisikoStrategisPemdaID"`
	Why                                *datatypes.JSON                   `json:"why"` // JSON array field
	AkarPenyebab                       *string                           `json:"akar_penyebab"`
	PenyebabRisikoID                   *string                           `json:"-"` // Foreign key field
	PenyebabRisiko                     *PenyebabRisiko                   `json:"penyebab_risiko" gorm:"foreignKey:PenyebabRisikoID"`
	KegiatanPengendalian               *string                           `json:"kegiatan_pengendalian"`
	Status                             sharedModel.Status                `json:"status"`
	CreatedAt                          time.Time                         `json:"created_at"`
	UpdatedAt                          time.Time                         `json:"updated_at"`
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
