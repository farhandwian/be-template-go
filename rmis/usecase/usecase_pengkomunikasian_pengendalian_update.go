package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PengkomunikasianPengendalianUpdateUseCaseReq struct {
	ID                   string `json:"id"`
	PenilaianRisikoID    string `json:"penilaian_risiko_id"`
	MediaKomunikasi      string `json:"media_komunikasi"`
	PenyediaInformasi    string `json:"penyedia_informasi"`
	PenerimaInformasi    string `json:"penerima_informasi"`
	RencanaPelaksanaan   string `json:"rencana_pelaksanaan"`
	RealisasiPelaksanaan string `json:"realiasi_pelaksanaan"`
	Keterangan           string `json:"keterangan"`
}

type PengkomunikasianPengendalianUpdateUseCaseRes struct{}

type PengkomunikasianPengendalianUpdateUseCase = core.ActionHandler[PengkomunikasianPengendalianUpdateUseCaseReq, PengkomunikasianPengendalianUpdateUseCaseRes]

func ImplPengkomunikasianPengendalianUpdateUseCase(
	getPengkomunikasianPengendalianById gateway.PengkomunikasianPengendalianGetByID,
	updatePengkomunikasianPengendalian gateway.PengkomunikasianPengendalianSave,
	PenilaianRisikoByID gateway.PenilaianRisikoGetByID,
) PengkomunikasianPengendalianUpdateUseCase {
	return func(ctx context.Context, req PengkomunikasianPengendalianUpdateUseCaseReq) (*PengkomunikasianPengendalianUpdateUseCaseRes, error) {

		res, err := getPengkomunikasianPengendalianById(ctx, gateway.PengkomunikasianPengendalianGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		_, err = PenilaianRisikoByID(ctx, gateway.PenilaianRisikoGetByIDReq{ID: req.PenilaianRisikoID})
		if err != nil {
			return nil, err
		}
		pengkomunikasiPengendalian := res.PengkomunikasianPengendalian

		pengkomunikasiPengendalian.PenyediaInformasi = &req.PenyediaInformasi
		pengkomunikasiPengendalian.PenerimaInformasi = &req.PenerimaInformasi
		pengkomunikasiPengendalian.RencanaPelaksanaan = &req.RencanaPelaksanaan
		pengkomunikasiPengendalian.RealisasiPelaksanaan = &req.RealisasiPelaksanaan
		pengkomunikasiPengendalian.Keterangan = &req.Keterangan
		pengkomunikasiPengendalian.Status = sharedModel.StatusMenungguVerifikasi

		if _, err := updatePengkomunikasianPengendalian(ctx, gateway.PengkomunikasianPengendalianSaveReq{PengkomunikasianPengendalian: pengkomunikasiPengendalian}); err != nil {
			return nil, err
		}

		return &PengkomunikasianPengendalianUpdateUseCaseRes{}, nil
	}
}
