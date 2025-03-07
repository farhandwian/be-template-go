package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type DaftarRisikoPrioritasUpdateUseCaseReq struct {
	ID                                     string `json:"id"`
	HasilAnalisisRisikoID                  string `json:"hasil_analisis_risiko_id"`
	PenetapanKonteksRisikoStrategisPemdaID string `json:"penetapan_konteks_risiko_strategis_pemda_id"`
}

type DaftarRisikoPrioritasUpdateUseCaseRes struct{}

type DaftarRisikoPrioritasUpdateUseCase = core.ActionHandler[DaftarRisikoPrioritasUpdateUseCaseReq, DaftarRisikoPrioritasUpdateUseCaseRes]

func ImplDaftarRisikoPrioritasUpdateUseCase(
	getDaftarRisikoPrioritasById gateway.DaftarRisikoPrioritasGetByID,
	updateDaftarRisikoPrioritas gateway.DaftarRisikoPrioritasSave,
	HasilAnalisisRisikoByID gateway.HasilAnalisisRisikoGetByID,
	IdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	PenetapanKonteksRisikoStrategisPemdaByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
) DaftarRisikoPrioritasUpdateUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasUpdateUseCaseReq) (*DaftarRisikoPrioritasUpdateUseCaseRes, error) {

		res, err := getDaftarRisikoPrioritasById(ctx, gateway.DaftarRisikoPrioritasGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		// hasilAnalisisRisikoByIDRes, err := HasilAnalisisRisikoByID(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.HasilAnalisisRisikoID})
		// if err != nil {
		// 	return nil, err
		// }

		// _, err = IdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{
		// 	ID: *hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.IdentifikasiRisikoStrategisPemdaID})

		// if err != nil {
		// 	return nil, err
		// }

		_, err = PenetapanKonteksRisikoStrategisPemdaByID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.PenetapanKonteksRisikoStrategisPemdaID})
		if err != nil {
			return nil, fmt.Errorf("error getting Penetapan Konteks Risiko Strategis Pemda table: %v", err)
		}

		// res.DaftarRisikoPrioritas.HasilAnalisisRisikoID = &req.HasilAnalisisRisikoID
		// res.DaftarRisikoPrioritas.RisikoPrioritas = hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.RisikoTeridentifikasi
		// res.DaftarRisikoPrioritas.KodeRisiko = hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.KodeRisiko
		// res.DaftarRisikoPrioritas.KategoriRisiko = hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.KategoriRisiko
		// res.DaftarRisikoPrioritas.PemilikRisiko = identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.PemilikRisiko
		// res.DaftarRisikoPrioritas.PenyebabRisiko = identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.UraianSebab
		// res.DaftarRisikoPrioritas.DampakRisiko = identifikasiRisikoStrategisPemdaByIDRes.IdentifikasiRisikoStrategisPemda.UraianDampak

		res.DaftarRisikoPrioritas.PenetapanKonteksRisikoStrategisPemdaID = &req.PenetapanKonteksRisikoStrategisPemdaID
		res.DaftarRisikoPrioritas.Status = sharedModel.StatusMenungguVerifikasi

		if _, err := updateDaftarRisikoPrioritas(ctx, gateway.DaftarRisikoPrioritasSaveReq{DaftarRisikoPrioritas: res.DaftarRisikoPrioritas}); err != nil {
			return nil, err
		}

		return &DaftarRisikoPrioritasUpdateUseCaseRes{}, nil
	}
}
