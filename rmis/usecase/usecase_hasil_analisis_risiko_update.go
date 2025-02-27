package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type HasilAnalisisRisikoUpdateUseCaseReq struct {
	ID                                 string                    `json:"-"`
	IdentifikasiRisikoStrategisPemdaID string                    `json:"identifikasi_risiko_strategis_pemda_id"`
	KriteriaKemungkinanInherentRisk    model.KriteriaKemungkinan `json:"kriteria_kemungkinan_inherent_risk"`
	SkorKemungkinanInherentRisk        int                       `json:"skor_kemungkinan_inherent_risk"`
	KriteriaDampakInherentRisk         string                    `json:"kriteria_dampak_inherent_risk"`
	SkorDampakInherentRisk             int                       `json:"skor_dampak_inherent_risk"`
	StatusAda                          string                    `json:"status_ada"`
	UraianControl                      string                    `json:"uraian_control"`
	KlarifikasiSPIP                    string                    `json:"klarifikasi_spip"`
	MemadaiControl                     string                    `json:"memadai_control"` // enum memadai (can also be defined as a custom type)
	KriteriaKemungkinanResidualRisk    model.KriteriaKemungkinan `json:"kriteria_kemungkinan_residual_risk"`
	SkorKemungkinanResidualRisk        int                       `json:"skor_kemungkinan_residual_risk"`
	KriteriaDampakResidualRisk         string                    `json:"kriteria_dampak_residual_risk"`
	SkorDampakResidualRisk             int                       `json:"skor_dampak_residual_risk"`
	SkalaRisikoResidualRisk            int                       `json:"skala_risiko_residual_risk"`
}

type HasilAnalisisRisikoUpdateUseCaseRes struct{}

type HasilAnalisisRisikoUpdateUseCase = core.ActionHandler[HasilAnalisisRisikoUpdateUseCaseReq, HasilAnalisisRisikoUpdateUseCaseRes]

func ImplHasilAnalisisRisikoUpdateUseCase(
	getHasilAnalisisRisikoById gateway.HasilAnalisisRisikoGetByID,
	updateHasilAnalisisRisiko gateway.HasilAnalisisRisikoSave,
	indeksPeringkatPrioritasByID gateway.IndeksPeringkatPrioritasGetByID,
	IndeksPeringkatPrioritasCreate gateway.IndeksPeringkatPrioritasSave,
) HasilAnalisisRisikoUpdateUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoUpdateUseCaseReq) (*HasilAnalisisRisikoUpdateUseCaseRes, error) {

		res, err := getHasilAnalisisRisikoById(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		hasilAnalisisRisiko := res.HasilAnalisisRisiko

		hasilAnalisisRisiko.IdentifikasiRisikoStrategisPemerintahDaerahID = &req.IdentifikasiRisikoStrategisPemdaID
		hasilAnalisisRisiko.SkorKemungkinanInherentRisk = &req.SkorKemungkinanInherentRisk
		hasilAnalisisRisiko.KriteriaDampakInherentRisk = &req.KriteriaDampakInherentRisk
		hasilAnalisisRisiko.SkorDampakInherentRisk = &req.SkorDampakInherentRisk
		hasilAnalisisRisiko.StatusAda = &req.StatusAda
		hasilAnalisisRisiko.UraianControl = &req.UraianControl
		hasilAnalisisRisiko.KlarifikasiSPIP = &req.KlarifikasiSPIP
		hasilAnalisisRisiko.MemadaiControl = &req.MemadaiControl
		hasilAnalisisRisiko.SkorKemungkinanResidualRisk = &req.SkorKemungkinanResidualRisk
		hasilAnalisisRisiko.KriteriaDampakResidualRisk = &req.KriteriaDampakResidualRisk
		hasilAnalisisRisiko.SkorDampakResidualRisk = &req.SkorDampakResidualRisk
		hasilAnalisisRisiko.SkalaRisikoResidualRisk = &req.SkalaRisikoResidualRisk

		hasilAnalisisRisiko.SetSkalaRisiko()
		if _, err := updateHasilAnalisisRisiko(ctx, gateway.HasilAnalisisRisikoSaveReq{HasilAnalisisRisiko: hasilAnalisisRisiko}); err != nil {
			return nil, err
		}

		indeksPeringkatPrioritasByIDRes, err := indeksPeringkatPrioritasByID(ctx, gateway.IndeksPeringkatPrioritasGetByIDReq{ID: *hasilAnalisisRisiko.ID})
		if err != nil {
			return nil, err
		}
		indeksPeringkatPrioritas := indeksPeringkatPrioritasByIDRes.IndeksPeringkatPrioritas

		indeksPeringkatPrioritas.SetToleransiRisiko(*hasilAnalisisRisiko.KategoriRisiko)
		indeksPeringkatPrioritas.SetMitigasi(*hasilAnalisisRisiko.SkalaRisikoResidualRisk)
		indeksPeringkatPrioritas.SetIntermediateRank(*hasilAnalisisRisiko.SkalaRisikoResidualRisk, *hasilAnalisisRisiko.NomorUraian)
		if _, err = IndeksPeringkatPrioritasCreate(ctx, gateway.IndeksPeringkatPrioritasSaveReq{IndeksPeringkatPrioritas: indeksPeringkatPrioritas}); err != nil {
			return nil, err
		}

		return &HasilAnalisisRisikoUpdateUseCaseRes{}, nil
	}
}
