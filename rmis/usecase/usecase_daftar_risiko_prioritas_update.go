package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type DaftarRisikoPrioritasUpdateUseCaseReq struct {
	ID                    string `json:"id"`
	HasilAnalisisRisikoID string `json:"hasil_analisis_risiko_id"`
}

type DaftarRisikoPrioritasUpdateUseCaseRes struct{}

type DaftarRisikoPrioritasUpdateUseCase = core.ActionHandler[DaftarRisikoPrioritasUpdateUseCaseReq, DaftarRisikoPrioritasUpdateUseCaseRes]

func ImplDaftarRisikoPrioritasUpdateUseCase(
	getDaftarRisikoPrioritasById gateway.DaftarRisikoPrioritasGetByID,
	updateDaftarRisikoPrioritas gateway.DaftarRisikoPrioritasSave,
	HasilAnalisisRisikoByID gateway.HasilAnalisisRisikoGetByID,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
) DaftarRisikoPrioritasUpdateUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasUpdateUseCaseReq) (*DaftarRisikoPrioritasUpdateUseCaseRes, error) {

		res, err := getDaftarRisikoPrioritasById(ctx, gateway.DaftarRisikoPrioritasGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		hasilAnalisisRisikoByIDRes, err := HasilAnalisisRisikoByID(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.HasilAnalisisRisikoID})
		if err != nil {
			return nil, err
		}

		identifikasiRisikoStrategisPemdaByIDRes, err := IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{
			ID: *hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.IdentifikasiRisikoStrategisPemdaID})

		if err != nil {
			return nil, err
		}

		// res.DaftarRisikoPrioritas.HasilAnalisisRisikoID = &req.HasilAnalisisRisikoID
		// res.DaftarRisikoPrioritas.RisikoPrioritas = hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.RisikoTeridentifikasi
		// res.DaftarRisikoPrioritas.KodeRisiko = hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.KodeRisiko
		// res.DaftarRisikoPrioritas.KategoriRisiko = hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.KategoriRisiko
		res.DaftarRisikoPrioritas.PemilikRisiko = identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.PemilikRisiko
		res.DaftarRisikoPrioritas.PenyebabRisiko = identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.UraianSebab
		res.DaftarRisikoPrioritas.DampakRisiko = identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.UraianDampak

		if _, err := updateDaftarRisikoPrioritas(ctx, gateway.DaftarRisikoPrioritasSaveReq{DaftarRisikoPrioritas: res.DaftarRisikoPrioritas}); err != nil {
			return nil, err
		}

		return &DaftarRisikoPrioritasUpdateUseCaseRes{}, nil
	}
}
