package model

import (
	"fmt"
	"time"
)

// form 4

type HasilAnalisisRisiko struct {
	ID        *string `json:"id"`
	NamaPemda *string `json:"nama_pemda"`
	// data diambil dari form 3a, 3b, dan 3c
	// ===========
	IdentifikasiRisikoStrategisPemerintahDaerahID *string    `json:"-"`
	TahunPenilaian                                *time.Time `json:"tahun_penilaian"`
	TujuanStrategis                               *string    `json:"tujuan_strategis"`
	UrusanPemerintahan                            *string    `json:"urusan_pemerintahan"`
	RisikoTeridentifikasi                         *string    `json:"risiko_teridentifikasi"`
	KodeRisiko                                    *string    `json:"kode_risiko"`
	KategoriRisiko                                *string    `json:"kategori_risiko"`
	// ===========
	KriteriaKemungkinanInherentRisk *string   `json:"kriteria_kemungkinan_inherent_risk"`
	SkorKemungkinanInherentRisk     *int      `json:"skor_kemungkinan_inherent_risk"`
	KriteriaDampakInherentRisk      *string   `json:"kriteria_dampak_inherent_risk"`
	SkorDampakInherentRisk          *int      `json:"skor_dampak_inherent_risk"`
	SkalaRisikoInherentRisk         *int      `json:"skala_risiko_inherent_risk"`
	StatusAda                       *string   `json:"status_ada"`
	UraianControl                   *string   `json:"uraian_control"`
	KlarifikasiSPIP                 *string   `json:"klarifikasi_spip"`
	MemadaiControl                  *string   `json:"memadai_control"` // enum memadai (can also be defined as a custom type)
	KriteriaKemungkinanResidualRisk *string   `json:"kriteria_kemungkinan_residual_risk"`
	SkorKemungkinanResidualRisk     *int      `json:"skor_kemungkinan_residual_risk"`
	KriteriaDampakResidualRisk      *string   `json:"kriteria_dampak_residual_risk"`
	SkorDampakResidualRisk          *int      `json:"skor_dampak_residual_risk"`
	SkalaRisikoResidualRisk         *int      `json:"skala_risiko_residual_risk"`
	CreatedAt                       time.Time `json:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at"`
}

var RiskMatrix = [][]int{
	{1, 3, 5, 8, 20},    // Likelihood 1 (Hampir Tidak Terjadi)
	{2, 7, 11, 13, 21},  // Likelihood 2 (Jarang Terjadi)
	{4, 10, 14, 17, 22}, // Likelihood 3 (Kadang Terjadi)
	{6, 12, 16, 19, 24}, // Likelihood 4 (Sering Terjadi)
	{9, 15, 18, 23, 25}, // Likelihood 5 (Hampir Pasti Terjadi)
}

// GetRiskScore menghitung skor risiko berdasarkan likelihood dan impact
func GetRiskScore(likelihood, impact int) int {
	// Validasi input harus dalam rentang 1-5
	if likelihood < 1 || likelihood > 5 || impact < 1 || impact > 5 {
		fmt.Println("Error: Likelihood dan Impact harus bernilai antara 1 sampai 5")
		return -1
	}
	// Ambil nilai risiko dari matriks
	return RiskMatrix[likelihood-1][impact-1]
}

func (har *HasilAnalisisRisiko) SetSkalaRisiko() error {
	if har.SkorKemungkinanInherentRisk == nil {
		return fmt.Errorf("SkorKemungkinanInherentRisk is nil")
	}

	if har.SkorDampakInherentRisk == nil {
		return fmt.Errorf("SkorDampakInherentRisk is nil")
	}

	if har.SkorKemungkinanResidualRisk == nil {
		return fmt.Errorf("SkorKemungkinanResidualRisk is nil")
	}

	if har.SkorDampakResidualRisk == nil {
		return fmt.Errorf("SkorDampakResidualRisk is nil")
	}

	// Hitung skala risiko inherent risk
	skalaRisikoInherentRisk := GetRiskScore(*har.SkorKemungkinanInherentRisk, *har.SkorDampakInherentRisk)
	har.SkalaRisikoInherentRisk = &skalaRisikoInherentRisk
	// Hitung skala risiko residual risk
	skalaRisikoResidualRisk := GetRiskScore(*har.SkorKemungkinanResidualRisk, *har.SkorDampakResidualRisk)
	har.SkalaRisikoResidualRisk = &skalaRisikoResidualRisk
	return nil
}
