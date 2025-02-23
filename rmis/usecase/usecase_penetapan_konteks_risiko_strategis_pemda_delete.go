package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenetapanKonteksRisikoDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoDeleteUseCaseRes struct{}

type PenetapanKonteksRisikoDeleteUseCase = core.ActionHandler[PenetapanKonteksRisikoDeleteUseCaseReq, PenetapanKonteksRisikoDeleteUseCaseRes]

func ImplPenetapanKonteksRisikoDeleteUseCase(deletePenetapanKonteksRisiko gateway.PenetepanKonteksRisikoStrategisPemdaDelete) PenetapanKonteksRisikoDeleteUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoDeleteUseCaseReq) (*PenetapanKonteksRisikoDeleteUseCaseRes, error) {

		if _, err := deletePenetapanKonteksRisiko(ctx, gateway.PenetepanKonteksRisikoStrategisPemdaDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoDeleteUseCaseRes{}, nil
	}
}
