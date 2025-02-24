package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type HasilAnalisisRisikoCreateUseCaseReq struct {
	IdentifikasiRisikoStrategisPemdaID string `json:"identifikasi_risiko_strategis_pemda_id"`
	KriteriaKemungkinanInherentRisk    string `json:"kriteria_kemungkinan_inherent_risk"`
	SkorKemungkinanInherentRisk        int    `json:"skor_kemungkinan_inherent_risk"`
	KriteriaDampakInherentRisk         string `json:"kriteria_dampak_inherent_risk"`
	SkorDampakInherentRisk             int    `json:"skor_dampak_inherent_risk"`
	StatusAda                          string `json:"status_ada"`
	UraianControl                      string `json:"uraian_control"`
	KlarifikasiSPIP                    string `json:"klarifikasi_spip"`
	MemadaiControl                     string `json:"memadai_control"` // enum memadai (can also be defined as a custom type)
	KriteriaKemungkinanResidualRisk    string `json:"kriteria_kemungkinan_residual_risk"`
	SkorKemungkinanResidualRisk        int    `json:"skor_kemungkinan_residual_risk"`
	KriteriaDampakResidualRisk         string `json:"kriteria_dampak_residual_risk"`
	SkorDampakResidualRisk             int    `json:"skor_dampak_residual_risk"`
	SkalaRisikoResidualRisk            int    `json:"skala_risiko_residual_risk"`
}

type HasilAnalisisRisikoCreateUseCaseRes struct {
	ID string `json:"id"`
}

type HasilAnalisisRisikoCreateUseCase = core.ActionHandler[HasilAnalisisRisikoCreateUseCaseReq, HasilAnalisisRisikoCreateUseCaseRes]

func ImplHasilAnalisisRisikoCreateUseCase(
	generateId gateway.GenerateId,
	createHasilAnalisisRisiko gateway.HasilAnalisisRisikoSave,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
) HasilAnalisisRisikoCreateUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoCreateUseCaseReq) (*HasilAnalisisRisikoCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		identifikasiRisikoStrategisPemdaRes, err := IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.IdentifikasiRisikoStrategisPemdaID})
		if err != nil {
			return nil, fmt.Errorf("error getting Identifikasi Risiko Strategis Pemda table: %v", err)
		}

		obj := model.HasilAnalisisRisiko{
			ID:                              &genObj.RandomId,
			RisikoTeridentifikasi:           identifikasiRisikoStrategisPemdaRes.IdentifikasiRisikoStrategisPemda.UraianRisiko,
			KodeRisiko:                      identifikasiRisikoStrategisPemdaRes.IdentifikasiRisikoStrategisPemda.KodeRisiko,
			KategoriRisiko:                  identifikasiRisikoStrategisPemdaRes.IdentifikasiRisikoStrategisPemda.KategoriRisikoName,
			KriteriaKemungkinanInherentRisk: &req.KriteriaKemungkinanInherentRisk,
			SkorKemungkinanInherentRisk:     &req.SkorKemungkinanInherentRisk,
			KriteriaDampakInherentRisk:      &req.KriteriaDampakInherentRisk,
			SkorDampakInherentRisk:          &req.SkorDampakInherentRisk,
			StatusAda:                       &req.StatusAda,
			UraianControl:                   &req.UraianControl,
			KlarifikasiSPIP:                 &req.KlarifikasiSPIP,
			MemadaiControl:                  &req.MemadaiControl,
			KriteriaKemungkinanResidualRisk: &req.KriteriaKemungkinanResidualRisk,
			SkorKemungkinanResidualRisk:     &req.SkorKemungkinanResidualRisk,
			KriteriaDampakResidualRisk:      &req.KriteriaDampakResidualRisk,
			SkorDampakResidualRisk:          &req.SkorDampakResidualRisk,
		}
		obj.SetSkalaRisiko()

		if _, err = createHasilAnalisisRisiko(ctx, gateway.HasilAnalisisRisikoSaveReq{HasilAnalisisRisiko: obj}); err != nil {
			return nil, err
		}

		return &HasilAnalisisRisikoCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
