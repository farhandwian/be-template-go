package model

import "time"

type SimpulanHasilSpip struct {
	ID        *string       `json:"id"`
	Spip      *string       `json:"-"`
	Simpulan  simpulanHasil `json:"simpulan"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type simpulanHasil string

const (
	simpulanHasilMemadai      = "Memadai"
	simpulanHasilTidakMemadai = "Tidak Memadai"
)

func (shs *SimpulanHasilSpip) SetSimpulanSPIP(allRecords []*RekapitulasiHasilKuesioner) {
	allMemadai := true
	for _, record := range allRecords {
		if record.SimpulanKuesioner != SimpulanKuesionerMemadai {
			allMemadai = false
			break
		}
	}

	if allMemadai {
		shs.Simpulan = simpulanHasilMemadai
	} else {
		shs.Simpulan = simpulanHasilTidakMemadai
	}
}
