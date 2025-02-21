package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type RekapitulasiHasilKuesionerGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type RekapitulasiHasilKuesionerGetByIDUseCaseRes struct {
	RekapitulasiHasilKuesioner model.RekapitulasiHasilKuesioner `json:"rekapitulasi_hasil_kuesioner"`
}

type RekapitulasiHasilKuesionerGetByIDUseCase = core.ActionHandler[RekapitulasiHasilKuesionerGetByIDUseCaseReq, RekapitulasiHasilKuesionerGetByIDUseCaseRes]

func ImplRekapitulasiHasilKuesionerGetByIDUseCase(getRekapitulasiHasilKuesionerByID gateway.RekapitulasiHasilKuesionerGetByID) RekapitulasiHasilKuesionerGetByIDUseCase {
	return func(ctx context.Context, req RekapitulasiHasilKuesionerGetByIDUseCaseReq) (*RekapitulasiHasilKuesionerGetByIDUseCaseRes, error) {
		res, err := getRekapitulasiHasilKuesionerByID(ctx, gateway.RekapitulasiHasilKuesionerGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &RekapitulasiHasilKuesionerGetByIDUseCaseRes{RekapitulasiHasilKuesioner: res.RekapitulasiHasilKuesioner}, nil
	}
}
