package usecase

import (
	"context"
	"fmt"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PengkomunikasianPengendalianCreateUseCaseReq struct {
	PenilaiKegiatanPengendalianID string `json:"penilai_kegiatan_pengendalian_id"`
	MediaSaranaPengkomunikasian   string `json:"media_sarana_pengkomunikasian"`
	PenyediaInformasi             string `json:"penyedia_informasi"`
	PenerimaInformasi             string `json:"penerima_informasi"`
	RencanaPelaksanaan            string `json:"rencana_pelaksanaan"`
	RealisasiPelaksanaan          string `json:"realiasi_pelaksanaan"`
	Keterangan                    string `json:"keterangan"`
}

type PengkomunikasianPengendalianCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PengkomunikasianPengendalianCreateUseCase = core.ActionHandler[PengkomunikasianPengendalianCreateUseCaseReq, PengkomunikasianPengendalianCreateUseCaseRes]

func ImplPengkomunikasianPengendalianCreateUseCase(
	generateId gateway.GenerateId,
	createPengkomunikasianPengendalian gateway.PengkomunikasianPengendalianSave,
	PenilaiKegiatanPengendalianByID gateway.PenilaianKegiatanPengendalianGetByID,
) PengkomunikasianPengendalianCreateUseCase {
	return func(ctx context.Context, req PengkomunikasianPengendalianCreateUseCaseReq) (*PengkomunikasianPengendalianCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		penilaiKegiatanPengendalianByIDRes, err := PenilaiKegiatanPengendalianByID(ctx, gateway.PenilaianKegiatanPengendalianGetByIDReq{ID: req.PenilaiKegiatanPengendalianID})
		if err != nil {
			return nil, err
		}
		fmt.Printf("Test %v", req.RealisasiPelaksanaan)

		obj := model.PengkomunikasianPengendalian{
			ID:                            &genObj.RandomId,
			PenilaiKegiatanPengendalianID: &req.PenilaiKegiatanPengendalianID,
			KegiatanPengendalian:          penilaiKegiatanPengendalianByIDRes.PenilaianKegiatanPengendalian.RencanaTindakPerbaikan,
			MediaSaranaPengkomunikasian:   &req.MediaSaranaPengkomunikasian,
			PenyediaInformasi:             &req.PenyediaInformasi,
			PenerimaInformasi:             &req.PenerimaInformasi,
			RencanaPelaksanaan:            &req.RencanaPelaksanaan,
			RealisasiPelaksanaan:          &req.RealisasiPelaksanaan,
			Keterangan:                    &req.Keterangan,
		}

		if _, err = createPengkomunikasianPengendalian(ctx, gateway.PengkomunikasianPengendalianSaveReq{PengkomunikasianPengendalian: obj}); err != nil {
			return nil, err
		}

		return &PengkomunikasianPengendalianCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
