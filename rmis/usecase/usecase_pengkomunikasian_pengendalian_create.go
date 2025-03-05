package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	sharedModel "shared/model"
)

type PengkomunikasianPengendalianCreateUseCaseReq struct {
	PenilaianRisikoID    string `json:"penilaian_risiko_id"`
	MediaKomunikasi      string `json:"media_komunikasi"`
	PenyediaInformasi    string `json:"penyedia_informasi"`
	PenerimaInformasi    string `json:"penerima_informasi"`
	RencanaPelaksanaan   string `json:"rencana_pelaksanaan"`
	RealisasiPelaksanaan string `json:"realiasi_pelaksanaan"`
	Keterangan           string `json:"keterangan"`
}

type PengkomunikasianPengendalianCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PengkomunikasianPengendalianCreateUseCase = core.ActionHandler[PengkomunikasianPengendalianCreateUseCaseReq, PengkomunikasianPengendalianCreateUseCaseRes]

func ImplPengkomunikasianPengendalianCreateUseCase(
	generateId gateway.GenerateId,
	createPengkomunikasianPengendalian gateway.PengkomunikasianPengendalianSave,
	PenilaianRisikoByID gateway.PenilaianRisikoGetByID,
) PengkomunikasianPengendalianCreateUseCase {
	return func(ctx context.Context, req PengkomunikasianPengendalianCreateUseCaseReq) (*PengkomunikasianPengendalianCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		_, err = PenilaianRisikoByID(ctx, gateway.PenilaianRisikoGetByIDReq{ID: req.PenilaianRisikoID})
		if err != nil {
			return nil, err
		}
		obj := model.PengkomunikasianPengendalian{
			ID:                   &genObj.RandomId,
			PenilaianRisikoID:    &req.PenilaianRisikoID,
			MediaKomunikasi:      &req.MediaKomunikasi,
			PenyediaInformasi:    &req.PenyediaInformasi,
			PenerimaInformasi:    &req.PenerimaInformasi,
			RencanaPelaksanaan:   &req.RencanaPelaksanaan,
			RealisasiPelaksanaan: &req.RealisasiPelaksanaan,
			Keterangan:           &req.Keterangan,
			Status:               sharedModel.StatusMenungguVerifikasi,
		}

		if _, err = createPengkomunikasianPengendalian(ctx, gateway.PengkomunikasianPengendalianSaveReq{PengkomunikasianPengendalian: obj}); err != nil {
			return nil, err
		}

		return &PengkomunikasianPengendalianCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
