package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PengkomunikasianPengendalianGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PengkomunikasianPengendalianGetByIDUseCaseRes struct {
	PengkomunikasianPengendalian model.PengkomunikasianPengendalian `json:"pengkomunikasian_pengendalian"`
}

type PengkomunikasianPengendalianGetByIDUseCase = core.ActionHandler[PengkomunikasianPengendalianGetByIDUseCaseReq, PengkomunikasianPengendalianGetByIDUseCaseRes]

func ImplPengkomunikasianPengendalianGetByIDUseCase(getPengkomunikasianPengendalianByID gateway.PengkomunikasianPengendalianGetByID) PengkomunikasianPengendalianGetByIDUseCase {
	return func(ctx context.Context, req PengkomunikasianPengendalianGetByIDUseCaseReq) (*PengkomunikasianPengendalianGetByIDUseCaseRes, error) {
		res, err := getPengkomunikasianPengendalianByID(ctx, gateway.PengkomunikasianPengendalianGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &PengkomunikasianPengendalianGetByIDUseCaseRes{PengkomunikasianPengendalian: res.PengkomunikasianPengendalian}, nil
	}
}
