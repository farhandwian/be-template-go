package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenilaianKegiatanPengendalianGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PenilaianKegiatanPengendalianGetByIDUseCaseRes struct {
	PenilaianKegiatanPengendalian model.PenilaianKegiatanPengendalian `json:"penilaian_kegiatan_pengendalians"`
}

type PenilaianKegiatanPengendalianGetByIDUseCase = core.ActionHandler[PenilaianKegiatanPengendalianGetByIDUseCaseReq, PenilaianKegiatanPengendalianGetByIDUseCaseRes]

func ImplPenilaianKegiatanPengendalianGetByIDUseCase(getPenilaianKegiatanPengendalianByID gateway.PenilaianKegiatanPengendalianGetByID) PenilaianKegiatanPengendalianGetByIDUseCase {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianGetByIDUseCaseReq) (*PenilaianKegiatanPengendalianGetByIDUseCaseRes, error) {
		res, err := getPenilaianKegiatanPengendalianByID(ctx, gateway.PenilaianKegiatanPengendalianGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &PenilaianKegiatanPengendalianGetByIDUseCaseRes{PenilaianKegiatanPengendalian: res.PenilaianKegiatanPengendalian}, nil
	}
}
