package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type HasilAnalisisRisikoUpdateUseCaseReq struct {
	ID                                 string                    `json:"id"`
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
) HasilAnalisisRisikoUpdateUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoUpdateUseCaseReq) (*HasilAnalisisRisikoUpdateUseCaseRes, error) {

		res, err := getHasilAnalisisRisikoById(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.HasilAnalisisRisiko.IdentifikasiRisikoStrategisPemerintahDaerahID = &req.IdentifikasiRisikoStrategisPemdaID
		res.HasilAnalisisRisiko.SkorKemungkinanInherentRisk = &req.SkorKemungkinanInherentRisk
		res.HasilAnalisisRisiko.KriteriaDampakInherentRisk = &req.KriteriaDampakInherentRisk
		res.HasilAnalisisRisiko.SkorDampakInherentRisk = &req.SkorDampakInherentRisk
		res.HasilAnalisisRisiko.StatusAda = &req.StatusAda
		res.HasilAnalisisRisiko.UraianControl = &req.UraianControl
		res.HasilAnalisisRisiko.KlarifikasiSPIP = &req.KlarifikasiSPIP
		res.HasilAnalisisRisiko.MemadaiControl = &req.MemadaiControl
		res.HasilAnalisisRisiko.SkorKemungkinanResidualRisk = &req.SkorKemungkinanResidualRisk
		res.HasilAnalisisRisiko.KriteriaDampakResidualRisk = &req.KriteriaDampakResidualRisk
		res.HasilAnalisisRisiko.SkorDampakResidualRisk = &req.SkorDampakResidualRisk
		res.HasilAnalisisRisiko.SkalaRisikoResidualRisk = &req.SkalaRisikoResidualRisk

		res.HasilAnalisisRisiko.SetKriteriaKemungkinan("inherent", req.KriteriaKemungkinanInherentRisk)
		res.HasilAnalisisRisiko.SetKriteriaKemungkinan("residual", req.KriteriaKemungkinanResidualRisk)

		res.HasilAnalisisRisiko.SetSkalaRisiko()
		if _, err := updateHasilAnalisisRisiko(ctx, gateway.HasilAnalisisRisikoSaveReq{HasilAnalisisRisiko: res.HasilAnalisisRisiko}); err != nil {
			return nil, err
		}

		return &HasilAnalisisRisikoUpdateUseCaseRes{}, nil
	}
}
