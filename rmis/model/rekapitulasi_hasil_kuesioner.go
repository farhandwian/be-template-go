package model

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/datatypes"
)

// Form1A
// Rekapitulasi Hasil Kuesioner Penilaian Lingkungan Pengendalian Intern

type RekapitulasiHasilKuesioner struct {
	ID                *string           `json:"id"`
	NamaPemda         *string           `json:"nama_pemda"`
	Pertanyaan        *string           `json:"pertanyaan"`
	SpipID            *string           `json:"-"`
	NamaSpip          *string           `json:"nama_spip"`
	JawabanResponden  *datatypes.JSON   `json:"jawaban_responden"`
	Modus             *int              `json:"modus"`
	SimpulanKuesioner SimpulanKuesioner `json:"simpulan_kuesioner"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
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
