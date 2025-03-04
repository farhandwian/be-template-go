package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type DaftarRisikoPrioritasCreateUseCaseReq struct {
	HasilAnalisisRisikoID string `json:"hasil_analisis_risiko_id"`
}

type DaftarRisikoPrioritasCreateUseCaseRes struct {
	ID string `json:"id"`
}

type DaftarRisikoPrioritasCreateUseCase = core.ActionHandler[DaftarRisikoPrioritasCreateUseCaseReq, DaftarRisikoPrioritasCreateUseCaseRes]

func ImplDaftarRisikoPrioritasCreateUseCase(
	generateId gateway.GenerateId,
	createDaftarRisikoPrioritas gateway.DaftarRisikoPrioritasSave,
	HasilAnalisisRisikoByID gateway.HasilAnalisisRisikoGetByID,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
) DaftarRisikoPrioritasCreateUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasCreateUseCaseReq) (*DaftarRisikoPrioritasCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		hasilAnalisisRisikoByIDRes, err := HasilAnalisisRisikoByID(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.HasilAnalisisRisikoID})
		if err != nil {
			return nil, err
		}
		_, err = IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{
			ID: *hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.IdentifikasiRisikoStrategisPemdaID})

		if err != nil {
			return nil, err
		}

		obj := model.DaftarRisikoPrioritas{
			ID: &genObj.RandomId,
			// HasilAnalisisRisikoID: &req.HasilAnalisisRisikoID,
			// RisikoPrioritas:       hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.RisikoTeridentifikasi,
			// KodeRisiko:            hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.KodeRisiko,
			// KategoriRisiko:        hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.KategoriRisiko,
			// PemilikRisiko:         identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.PemilikRisiko,
			// PenyebabRisiko:        identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.UraianSebab,
			// DampakRisiko:          identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.UraianDampak,
		}

		if _, err = createDaftarRisikoPrioritas(ctx, gateway.DaftarRisikoPrioritasSaveReq{DaftarRisikoPrioritas: obj}); err != nil {
			return nil, err
		}

		return &DaftarRisikoPrioritasCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
