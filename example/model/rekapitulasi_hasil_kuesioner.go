package model

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

// Form1A
// Rekapitulasi Hasil Kuesioner Penilaian Lingkungan Pengendalian Intern

type RekapitulasiHasilKuesioner struct {
	ID                *string           `json:"id"`
	NamaPemda         *string           `json:"nama_pemda"`
	SPIP              *string           `json:"spip"`
	JawabanResponden  *datatypes.JSON   `json:"jawaban_responden"`
	Modus             *int              `json:"modus"`
	SimpulanKuesioner SimpulanKuesioner `json:"simpulan_kuesioner"`
	SimpulanSPIP      SimpulanKuesioner `json:"simpulan_spip"`
}

type SimpulanKuesioner string

const (
	SimpulanKuesionerMemadai      SimpulanKuesioner = "Memadai"
	SimpulanKuesionerTidakMemadai SimpulanKuesioner = "Tidak Memadai"
)

func (rk *RekapitulasiHasilKuesioner) CalculateModus() error {
	// Unmarshal jawaban responden to []int
	var jawaban []int
	if err := json.Unmarshal(*rk.JawabanResponden, &jawaban); err != nil {
		return fmt.Errorf("failed to unmarshal jawaban responden: %v", err)
	}

	// calculate modus
	var sum int
	for _, value := range jawaban {
		sum += value
	}

	rk.Modus = &sum
	return nil
}

func (rk *RekapitulasiHasilKuesioner) SetSimpulanKuesioner() {
	if rk.Modus != nil && *rk.Modus > 2 {
		rk.SimpulanKuesioner = SimpulanKuesionerMemadai
	} else {
		rk.SimpulanKuesioner = SimpulanKuesionerTidakMemadai
	}
}

func (rk *RekapitulasiHasilKuesioner) SetSimpulanSPIP(allRecords []*RekapitulasiHasilKuesioner) {
	allMemadai := true
	for _, record := range allRecords {
		if record.SimpulanKuesioner != SimpulanKuesionerMemadai {
			allMemadai = false
			break
		}
	}

	if allMemadai {
		rk.SimpulanSPIP = SimpulanKuesionerMemadai
	} else {
		rk.SimpulanSPIP = SimpulanKuesionerTidakMemadai
	}
}
