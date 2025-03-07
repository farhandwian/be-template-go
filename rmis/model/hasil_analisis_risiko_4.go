package model

import (
	"fmt"
	sharedModel "shared/model"
	"time"
)

// form 4

type HasilAnalisisRisiko struct {
	ID               *string `json:"id"`
	TipeIdentifikasi *string `json:"tipe_identifikasi" gorm:"type:VARCHAR(50)"` // "strategis_pemda", "operasional_opd", or "strategis_renstra_opd"
	IdentifikasiID   *string `json:"identifikasi_id" gorm:"type:VARCHAR(255)"`

	TipePenetapanKonteks *string `json:"tipe_penetapan_konteks" gorm:"type:VARCHAR(50)"` // "strategis_pemda", "operasional", or "strategis_renstra_opd"
	PenetapanKonteksID   *string `json:"penetapan_konteks_id" gorm:"type:VARCHAR(255)"`

	SkalaDampak      *int `json:"skala_dampak"`
	SkalaKemungkinan *int `json:"skala_kemungkinan"`
	SkalaRisiko      *int `json:"skala_risiko"`

	Status    sharedModel.Status `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type TipePenetapanKonteks string

const (
	TipePenetapanKonteksStrategisPemda      TipePenetapanKonteks = "strategis_pemda"
	TipePenetapanKonteksOperasional         TipePenetapanKonteks = "operasional"
	TipePenetapanKonteksStrategisRenstraOPD TipePenetapanKonteks = "strategis_renstra_opd"
)

type TipeIdentifikasi string

const (
	TipeIdentifikasiStrategisPemda TipeIdentifikasi = "strategis_pemda"
	TipeIdentifikasiOperasional    TipeIdentifikasi = "operasional_opd"
	TipeIdentifikasiStrategisOPD   TipeIdentifikasi = "strategis_opd"
)

type HasilAnalisisRisikoResponse struct {
	ID               *string `json:"id"`
	TipeIdentifikasi *string `json:"tipe_identifikasi" gorm:"type:VARCHAR(50)"` // "strategis_pemda", "operasional_opd", or "strategis_renstra_opd"
	IdentifikasiID   *string `json:"identifikasi_id" gorm:"type:VARCHAR(255)"`

	TipePenetapanKonteks *string `json:"tipe_penetapan_konteks" gorm:"type:VARCHAR(50)"` // "strategis_pemda", "operasional", or "strategis_renstra_opd"
	PenetapanKonteksID   *string `json:"penetapan_konteks_id" gorm:"type:VARCHAR(255)"`

	SkalaDampak      *int `json:"skala_dampak"`
	SkalaKemungkinan *int `json:"skala_kemungkinan"`
	SkalaRisiko      *int `json:"skala_risiko"`

	NamaPemda          *string            `json:"nama_pemda"`
	Tahun              *time.Time         `json:"tahun"`
	Periode            *string            `json:"periode"`
	PenetapanKonteks   *string            `json:"penetapan_konteks"`
	UrusanPemerintahan *string            `json:"urusan_pemerintahan"`
	Status             sharedModel.Status `json:"status"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

var RiskMatrix = [][]int{
	{1, 3, 5, 8, 20},    // Likelihood 1 (Hampir Tidak Terjadi)
	{2, 7, 11, 13, 21},  // Likelihood 2 (Jarang Terjadi)
	{4, 10, 14, 17, 22}, // Likelihood 3 (Kadang Terjadi)
	{6, 12, 16, 19, 24}, // Likelihood 4 (Sering Terjadi)
	{9, 15, 18, 23, 25}, // Likelihood 5 (Hampir Pasti Terjadi)
}

func (har *HasilAnalisisRisiko) SetSkalaRisiko() error {
	if har.SkalaDampak == nil || har.SkalaKemungkinan == nil {
		return fmt.Errorf("Failed to set Skala Risiko: SkalaDampak or SkalaKemungkinan is nil")
	}

	result := (*har.SkalaDampak) * (*har.SkalaKemungkinan)
	har.SkalaRisiko = &result

	return nil
}

// GetRiskScore menghitung skor risiko berdasarkan likelihood dan impact
// func GetRiskScore(likelihood, impact int) int {
// 	// Validasi input harus dalam rentang 1-5
// 	if likelihood < 1 || likelihood > 5 || impact < 1 || impact > 5 {
// 		fmt.Println("Error: Likelihood dan Impact harus bernilai antara 1 sampai 5")
// 		return -1
// 	}
// 	// Ambil nilai risiko dari matriks
// 	return RiskMatrix[likelihood-1][impact-1]
// }

// func (har *HasilAnalisisRisiko) SetSkalaRisiko() error {
// 	if har.SkorKemungkinanInherentRisk == nil {
// 		return fmt.Errorf("SkorKemungkinanInherentRisk is nil")
// 	}

// 	if har.SkorDampakInherentRisk == nil {
// 		return fmt.Errorf("SkorDampakInherentRisk is nil")
// 	}

// 	if har.SkorKemungkinanResidualRisk == nil {
// 		return fmt.Errorf("SkorKemungkinanResidualRisk is nil")
// 	}

// 	if har.SkorDampakResidualRisk == nil {
// 		return fmt.Errorf("SkorDampakResidualRisk is nil")
// 	}

// 	// Hitung skala risiko inherent risk
// 	skalaRisikoInherentRisk := GetRiskScore(*har.SkorKemungkinanInherentRisk, *har.SkorDampakInherentRisk)
// 	har.SkalaRisikoInherentRisk = &skalaRisikoInherentRisk
// 	// Hitung skala risiko residual risk
// 	skalaRisikoResidualRisk := GetRiskScore(*har.SkorKemungkinanResidualRisk, *har.SkorDampakResidualRisk)
// 	har.SkalaRisikoResidualRisk = &skalaRisikoResidualRisk
// 	return nil
// }

// func (har *HasilAnalisisRisiko) SetKriteriaKemungkinan(tipeRisk string, tipeKemungkinan string) error {

// 	if tipeKemungkinan == "persentase" && tipeRisk == "inherent" {
// 		har.KriteriaKemungkinanInherentRisk = &PersentasePerTahun
// 	} else if tipeKemungkinan == "frekuensi" && tipeRisk == "inherent" {
// 		har.KriteriaKemungkinanInherentRisk = &FrekuensiPerTahun
// 	} else if tipeKemungkinan == "persentase" && tipeRisk == "residual" {
// 		har.KriteriaKemungkinanResidualRisk = &PersentasePerTahun
// 	} else if tipeKemungkinan == "frekuensi" && tipeRisk == "residual" {
// 		har.KriteriaKemungkinanResidualRisk = &FrekuensiPerTahun
// 	}
// 	return nil
// }
