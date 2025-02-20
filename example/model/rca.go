package model

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/datatypes"
)

// Root Cause Analysis

type RCA struct {
	ID                    *string         `json:"id"`
	NamaUnitPemilikRisiko *string         `json:"nama_unit_pemilik_risiko"`
	TahunPenilaian        *time.Time      `json:"tahun_penilaian"`
	PenyebabRisiko        *string         `json:"penyebab_risiko"`
	Why                   *datatypes.JSON `json:"why"` // it's json because it will contain array of strings
	AkarPenyebab          *string         `json:"akar_penyebab"`
	JenisPenyebab         *string         `json:"jenis_penyebab"`
	KegiatanPengendalian  *string         `json:"kegiatan_pengendalian"`
}

func (rca *RCA) SetAkarPenyebab() error {
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
