package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenilaianRisikoCreateUseCaseReq struct {
	DaftarRisikoPrioritasID   string `json:"daftar_risiko_prioritas_id"`
	CelahPengendalian         string `json:"celah_pengendalian"`
	RencanaTindakPengendalian string `json:"rencana_tindak_pengendalian"`
	PemilikPenanggungJawab    string `json:"pemilik_penanggung_jawab"`
	TargetWaktuPenyelesaian   string `json:"target_waktu_penyelesaian"`
}

type PenilaianRisikoCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PenilaianRisikoCreateUseCase = core.ActionHandler[PenilaianRisikoCreateUseCaseReq, PenilaianRisikoCreateUseCaseRes]

func ImplPenilaianRisikoCreateUseCase(
	generateId gateway.GenerateId,
	createPenilaianRisiko gateway.PenilaianRisikoSave,
	daftarRisikoPrioritasByID gateway.DaftarRisikoPrioritasGetByID,
	hasilAnalisisRisikoByID gateway.HasilAnalisisRisikoGetByID,
) PenilaianRisikoCreateUseCase {
	return func(ctx context.Context, req PenilaianRisikoCreateUseCaseReq) (*PenilaianRisikoCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		daftarRisikoPrioritasByIDRes, err := daftarRisikoPrioritasByID(ctx, gateway.DaftarRisikoPrioritasGetByIDReq{ID: req.DaftarRisikoPrioritasID})
		if err != nil {
			return nil, err
		}

		_, err = hasilAnalisisRisikoByID(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: *daftarRisikoPrioritasByIDRes.DaftarRisikoPrioritas.HasilAnalisisRisikoID})
		if err != nil {
			return nil, err
		}

		obj := model.PenilaianRisiko{
			ID:              &genObj.RandomId,
			RisikoPrioritas: daftarRisikoPrioritasByIDRes.DaftarRisikoPrioritas.RisikoPrioritas,
			KodeRisiko:      daftarRisikoPrioritasByIDRes.DaftarRisikoPrioritas.KodeRisiko,
			// UraianPengendalian:        hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.UraianControl,
			CelahPengendalian:         &req.CelahPengendalian,
			RencanaTindakPengendalian: &req.RencanaTindakPengendalian,
			PemilikPenanggungJawab:    &req.PemilikPenanggungJawab,
			TargetWaktuPenyelesaian:   &req.TargetWaktuPenyelesaian,
		}

		if _, err = createPenilaianRisiko(ctx, gateway.PenilaianRisikoSaveReq{PenilaianRisiko: obj}); err != nil {
			return nil, err
		}

		return &PenilaianRisikoCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
