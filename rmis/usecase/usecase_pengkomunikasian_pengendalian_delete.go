package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PengkomunikasianPengendalianDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type PengkomunikasianPengendalianDeleteUseCaseRes struct{}

type PengkomunikasianPengendalianDeleteUseCase = core.ActionHandler[PengkomunikasianPengendalianDeleteUseCaseReq, PengkomunikasianPengendalianDeleteUseCaseRes]

func ImplPengkomunikasianPengendalianDeleteUseCase(deletePengkomunikasianPengendalian gateway.PengkomunikasianPengendalianDelete) PengkomunikasianPengendalianDeleteUseCase {
	return func(ctx context.Context, req PengkomunikasianPengendalianDeleteUseCaseReq) (*PengkomunikasianPengendalianDeleteUseCaseRes, error) {

		if _, err := deletePengkomunikasianPengendalian(ctx, gateway.PengkomunikasianPengendalianDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &PengkomunikasianPengendalianDeleteUseCaseRes{}, nil
	}
}
