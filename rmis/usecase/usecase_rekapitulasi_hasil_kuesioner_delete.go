package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type RekapitulasiHasilKuesionerDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type RekapitulasiHasilKuesionerDeleteUseCaseRes struct{}

type RekapitulasiHasilKuesionerDeleteUseCase = core.ActionHandler[RekapitulasiHasilKuesionerDeleteUseCaseReq, RekapitulasiHasilKuesionerDeleteUseCaseRes]

func ImplRekapitulasiHasilKuesionerDeleteUseCase(deleteRekapitulasiHasilKuesioner gateway.RekapitulasiHasilKuesionerDelete) RekapitulasiHasilKuesionerDeleteUseCase {
	return func(ctx context.Context, req RekapitulasiHasilKuesionerDeleteUseCaseReq) (*RekapitulasiHasilKuesionerDeleteUseCaseRes, error) {

		if _, err := deleteRekapitulasiHasilKuesioner(ctx, gateway.RekapitulasiHasilKuesionerDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &RekapitulasiHasilKuesionerDeleteUseCaseRes{}, nil
	}
}
