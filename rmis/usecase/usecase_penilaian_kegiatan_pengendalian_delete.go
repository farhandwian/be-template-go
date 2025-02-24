package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenilaianKegiatanPengendalianDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type PenilaianKegiatanPengendalianDeleteUseCaseRes struct{}

type PenilaianKegiatanPengendalianDeleteUseCase = core.ActionHandler[PenilaianKegiatanPengendalianDeleteUseCaseReq, PenilaianKegiatanPengendalianDeleteUseCaseRes]

func ImplPenilaianKegiatanPengendalianDeleteUseCase(deletePenilaianKegiatanPengendalian gateway.PenilaianKegiatanPengendalianDelete) PenilaianKegiatanPengendalianDeleteUseCase {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianDeleteUseCaseReq) (*PenilaianKegiatanPengendalianDeleteUseCaseRes, error) {

		if _, err := deletePenilaianKegiatanPengendalian(ctx, gateway.PenilaianKegiatanPengendalianDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &PenilaianKegiatanPengendalianDeleteUseCaseRes{}, nil
	}
}
