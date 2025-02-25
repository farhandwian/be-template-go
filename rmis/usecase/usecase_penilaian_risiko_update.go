package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenilaianRisikoUpdateUseCaseReq struct {
	ID                        string `json:"id"`
	DaftarRisikoPrioritasID   string `json:"daftar_risiko_prioritas_id"`
	CelahPengendalian         string `json:"celah_pengendalian"`
	RencanaTindakPengendalian string `json:"rencana_tindak_pengendalian"`
	PemilikPenanggungJawab    string `json:"pemilik_penanggung_jawab"`
	TargetWaktuPenyelesaian   string `json:"target_waktu_penyelesaian"`
}

type PenilaianRisikoUpdateUseCaseRes struct{}

type PenilaianRisikoUpdateUseCase = core.ActionHandler[PenilaianRisikoUpdateUseCaseReq, PenilaianRisikoUpdateUseCaseRes]

func ImplPenilaianRisikoUpdateUseCase(
	getPenilaianRisikoById gateway.PenilaianRisikoGetByID,
	updatePenilaianRisiko gateway.PenilaianRisikoSave,
	daftarRisikoPrioritasByID gateway.DaftarRisikoPrioritasGetByID,
	hasilAnalisisRisikoByID gateway.HasilAnalisisRisikoGetByID,
) PenilaianRisikoUpdateUseCase {
	return func(ctx context.Context, req PenilaianRisikoUpdateUseCaseReq) (*PenilaianRisikoUpdateUseCaseRes, error) {

		res, err := getPenilaianRisikoById(ctx, gateway.PenilaianRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		daftarRisikoPrioritasByIDRes, err := daftarRisikoPrioritasByID(ctx, gateway.DaftarRisikoPrioritasGetByIDReq{ID: req.DaftarRisikoPrioritasID})
		if err != nil {
			return nil, err
		}

		hasilAnalisisRisikoByIDRes, err := hasilAnalisisRisikoByID(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: *daftarRisikoPrioritasByIDRes.DaftarRisikoPrioritas.HasilAnalisisRisikoID})
		if err != nil {
			return nil, err
		}
		penilaianRisiko := res.PenilaianRisiko

		penilaianRisiko.RisikoPrioritas = daftarRisikoPrioritasByIDRes.DaftarRisikoPrioritas.RisikoPrioritas
		penilaianRisiko.KodeRisiko = daftarRisikoPrioritasByIDRes.DaftarRisikoPrioritas.KodeRisiko
		penilaianRisiko.UraianPengendalian = hasilAnalisisRisikoByIDRes.HasilAnalisisRisiko.UraianControl
		penilaianRisiko.DaftarRisikoPrioritasID = &req.DaftarRisikoPrioritasID
		penilaianRisiko.CelahPengendalian = &req.CelahPengendalian
		penilaianRisiko.RencanaTindakPengendalian = &req.RencanaTindakPengendalian
		penilaianRisiko.PemilikPenanggungJawab = &req.PemilikPenanggungJawab
		penilaianRisiko.TargetWaktuPenyelesaian = &req.TargetWaktuPenyelesaian

		if _, err := updatePenilaianRisiko(ctx, gateway.PenilaianRisikoSaveReq{PenilaianRisiko: penilaianRisiko}); err != nil {
			return nil, err
		}

		return &PenilaianRisikoUpdateUseCaseRes{}, nil
	}
}
