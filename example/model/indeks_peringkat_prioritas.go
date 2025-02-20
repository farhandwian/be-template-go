package model

import (
	"fmt"
	"sort"
)

// Form 4 bagian kanan

type IndeksPeringkatPrioritas struct {
	ID               *string  `json:"id"`
	PreRank          *int     `json:"pre_rank"`
	ToleransiRisiko  *int     `json:"toleransi_risiko"`
	Mitigasi         *string  `json:"mitigasi"`
	UnikNumber       *int     `json:"unik_number"`
	IntermediateRank *float64 `json:"intermediate_rank"`
	FinalRank        *int     `json:"final_rank"`
}

var toleransiMapping = map[string]int{
	"Risiko Pendapatan":  10,
	"Risiko Belanja":     10,
	"Risiko Pembiayaan":  10,
	"Risiko Strategis":   9,
	"Risiko Fraud":       9,
	"Risiko Kepatuhan":   9,
	"Risiko Operasional": 15,
	"Risiko Reputasi":    15,
}

type mitigasiStatus string

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
	if skalaRisikoResidualRisk <= *ipr.ToleransiRisiko {
		ipr.Mitigasi = stringPtr(string(MitigasiAcceptable))
		*ipr.UnikNumber = 0
	} else {
		ipr.Mitigasi = stringPtr(string(MitigasiNotAcceptable))
		*ipr.UnikNumber = 1
	}
}
func stringPtr(s string) *string {
	return &s
}

func (ipr *IndeksPeringkatPrioritas) SetIntermediateRank(skalaRisikoResidualRisk int, nomor int) {
	if ipr.IntermediateRank == nil {
		// If nil, initialize it so we can assign a value
		ipr.IntermediateRank = new(float64)
	}
	if ipr.UnikNumber == nil {
		// Handle nil ipr.UnikNumber if necessary
		// For example, default to 0 or return an error
		defaultVal := 0
		ipr.UnikNumber = &defaultVal
	}

	value := float64(skalaRisikoResidualRisk) + float64(*ipr.UnikNumber) + (0.01 * float64(nomor))
	*ipr.IntermediateRank = value
}

func (ipr *IndeksPeringkatPrioritas) SetFinalRank(list []*IndeksPeringkatPrioritas) error {
	if len(list) == 0 {
		return nil
	}

	sort.SliceStable(list, func(i, j int) bool {
		// Handle nil IntermediateRank by treating it as 0
		var iVal, jVal float64
		if list[i].IntermediateRank != nil {
			iVal = *list[i].IntermediateRank
		}
		if list[j].IntermediateRank != nil {
			jVal = *list[j].IntermediateRank
		}
		return iVal < jVal
	})

	// Assign FinalRank in sorted order (1-based)
	for idx, item := range list {
		if item == nil {
			// If there's a nil pointer in the slice, handle as needed
			return fmt.Errorf("found nil item in the slice at index %d", idx)
		}
		rank := idx + 1 // 1-based rank
		item.FinalRank = &rank
	}

	return nil
}
