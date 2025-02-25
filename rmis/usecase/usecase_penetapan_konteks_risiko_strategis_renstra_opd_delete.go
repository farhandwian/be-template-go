package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenetapanKonteksRisikoRenstraOPDDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoRenstraOPDDeleteUseCaseRes struct{}

type PenetapanKonteksRisikoRenstraOPDDeleteUseCase = core.ActionHandler[PenetapanKonteksRisikoRenstraOPDDeleteUseCaseReq, PenetapanKonteksRisikoRenstraOPDDeleteUseCaseRes]

func ImplPenetapanKonteksRisikoRenstraOPDDeleteUseCase(deletePenetapanKonteksRisikoRenstraOPD gateway.PenetepanKonteksRisikoStrategisRenstraOPDDelete) PenetapanKonteksRisikoRenstraOPDDeleteUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoRenstraOPDDeleteUseCaseReq) (*PenetapanKonteksRisikoRenstraOPDDeleteUseCaseRes, error) {

		if _, err := deletePenetapanKonteksRisikoRenstraOPD(ctx, gateway.PenetepanKonteksRisikoStrategisRenstraOPDDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoRenstraOPDDeleteUseCaseRes{}, nil
	}
}
