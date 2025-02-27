package model

import (
	"fmt"
)

// Form 4 bagian kanan

type mitigasiStatus string
type IndeksPeringkatPrioritas struct {
	ID                    *string         `json:"id"`
	HasilAnalisisRisikoID *string         `json:"-"`
	PreRank               *int            `json:"pre_rank"`
	ToleransiRisiko       *int            `json:"toleransi_risiko"`
	Mitigasi              *mitigasiStatus `json:"mitigasi"`
	UnikNumber            *int            `json:"unik_number"`
	IntermediateRank      *float64        `json:"intermediate_rank"`
}

var toleransiMapping = map[string]int{
	"Risiko Pendapatan":  10,
	"Risiko Belanja":     10,
	"Risiko Pembiayaan":  10,
	"Risiko Strategis":   9,
	"Risiko Fraud":       4,
	"Risiko Kepatuhan":   9,
	"Risiko Operasional": 15,
	"Risiko Reputasi":    15,
}

const (
	MitigasiAcceptable    mitigasiStatus = "Acceptable"
	MitigasiNotAcceptable mitigasiStatus = "Unacceptable"
)

func (ipr *IndeksPeringkatPrioritas) SetToleransiRisiko(kategoriRisiko string) error {
	val, ok := toleransiMapping[kategoriRisiko]
	if !ok {
		return fmt.Errorf("kategori risiko %s tidak ditemukan", kategoriRisiko)
	}

	ipr.ToleransiRisiko = &val
	return nil
}

func (ipr *IndeksPeringkatPrioritas) SetMitigasi(skalaRisikoResidualRisk int) {
	acceptable := MitigasiAcceptable
	notAcceptable := MitigasiNotAcceptable

	if skalaRisikoResidualRisk <= *ipr.ToleransiRisiko {
		ipr.Mitigasi = &acceptable
		ipr.UnikNumber = new(int) // Ensure UnikNumber is not nil
		*ipr.UnikNumber = 0
	} else {
		ipr.Mitigasi = &notAcceptable
		ipr.UnikNumber = new(int) // Ensure UnikNumber is not nil
		*ipr.UnikNumber = 1
	}
}
func (ipr *IndeksPeringkatPrioritas) SetIntermediateRank(skalaRisikoResidualRisk int, nomor int) {
	if ipr.IntermediateRank == nil {
		ipr.IntermediateRank = new(float64)
	}
	if ipr.UnikNumber == nil {

		defaultVal := 0
		ipr.UnikNumber = &defaultVal
	}

	value := float64(skalaRisikoResidualRisk) + float64(*ipr.UnikNumber) + (0.01 * float64(nomor))
	*ipr.IntermediateRank = value
}
