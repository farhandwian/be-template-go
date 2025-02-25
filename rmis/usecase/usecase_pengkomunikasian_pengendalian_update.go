package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PengkomunikasianPengendalianUpdateUseCaseReq struct {
	ID                            string `json:"id"`
	PenilaiKegiatanPengendalianID string `json:"penilai_kegiatan_pengendalian"`
	MediaSaranaPengkomunikasian   string `json:"media_sarana_pengkomunikasian"`
	PenyediaInformasi             string `json:"penyedia_informasi"`
	PenerimaInformasi             string `json:"penerima_informasi"`
	RencanaPelaksanaan            string `json:"rencana_pelaksanaan"`
	RealisasiPelaksanaan          string `json:"realiasi_pelaksanaan"`
	Keterangan                    string `json:"keterangan"`
}

type PengkomunikasianPengendalianUpdateUseCaseRes struct{}

type PengkomunikasianPengendalianUpdateUseCase = core.ActionHandler[PengkomunikasianPengendalianUpdateUseCaseReq, PengkomunikasianPengendalianUpdateUseCaseRes]

func ImplPengkomunikasianPengendalianUpdateUseCase(
	getPengkomunikasianPengendalianById gateway.PengkomunikasianPengendalianGetByID,
	updatePengkomunikasianPengendalian gateway.PengkomunikasianPengendalianSave,
	PenilaiKegiatanPengendalianByID gateway.PenilaianKegiatanPengendalianGetByID,

) PengkomunikasianPengendalianUpdateUseCase {
	return func(ctx context.Context, req PengkomunikasianPengendalianUpdateUseCaseReq) (*PengkomunikasianPengendalianUpdateUseCaseRes, error) {

		res, err := getPengkomunikasianPengendalianById(ctx, gateway.PengkomunikasianPengendalianGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		penilaiKegiatanPengendalianByIDRes, err := PenilaiKegiatanPengendalianByID(ctx, gateway.PenilaianKegiatanPengendalianGetByIDReq{ID: req.PenilaiKegiatanPengendalianID})
		if err != nil {
			return nil, err
		}
		pengkomunikasiPengendalian := res.PengkomunikasianPengendalian

		pengkomunikasiPengendalian.PenilaiKegiatanPengendalianID = &req.PenilaiKegiatanPengendalianID
		pengkomunikasiPengendalian.KegiatanPengendalian = penilaiKegiatanPengendalianByIDRes.PenilaianKegiatanPengendalian.RencanaTindakPerbaikan
		pengkomunikasiPengendalian.MediaSaranaPengkomunikasian = &req.MediaSaranaPengkomunikasian
		pengkomunikasiPengendalian.PenyediaInformasi = &req.PenyediaInformasi
		pengkomunikasiPengendalian.PenerimaInformasi = &req.PenerimaInformasi
		pengkomunikasiPengendalian.RencanaPelaksanaan = &req.RencanaPelaksanaan
		pengkomunikasiPengendalian.RealisasiPelaksanaan = &req.RealisasiPelaksanaan
		pengkomunikasiPengendalian.Keterangan = &req.Keterangan

		if _, err := updatePengkomunikasianPengendalian(ctx, gateway.PengkomunikasianPengendalianSaveReq{PengkomunikasianPengendalian: pengkomunikasiPengendalian}); err != nil {
			return nil, err
		}

		return &PengkomunikasianPengendalianUpdateUseCaseRes{}, nil
	}
}
