package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenetapanKonteksRisikoOperasionalDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoOperasionalDeleteUseCaseRes struct{}

type PenetapanKonteksRisikoOperasionalDeleteUseCase = core.ActionHandler[PenetapanKonteksRisikoOperasionalDeleteUseCaseReq, PenetapanKonteksRisikoOperasionalDeleteUseCaseRes]

func ImplPenetapanKonteksRisikoOperasionalDeleteUseCase(deletePenetapanKonteksRisikoOperasional gateway.PenetepanKonteksRisikoOperasionalDelete) PenetapanKonteksRisikoOperasionalDeleteUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalDeleteUseCaseReq) (*PenetapanKonteksRisikoOperasionalDeleteUseCaseRes, error) {

		if _, err := deletePenetapanKonteksRisikoOperasional(ctx, gateway.PenetepanKonteksRisikoOperasionalDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoOperasionalDeleteUseCaseRes{}, nil
	}
}
