package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenilaianRisikoDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type PenilaianRisikoDeleteUseCaseRes struct{}

type PenilaianRisikoDeleteUseCase = core.ActionHandler[PenilaianRisikoDeleteUseCaseReq, PenilaianRisikoDeleteUseCaseRes]

func ImplPenilaianRisikoDeleteUseCase(deletePenilaianRisiko gateway.PenilaianRisikoDelete) PenilaianRisikoDeleteUseCase {
	return func(ctx context.Context, req PenilaianRisikoDeleteUseCaseReq) (*PenilaianRisikoDeleteUseCaseRes, error) {

		if _, err := deletePenilaianRisiko(ctx, gateway.PenilaianRisikoDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &PenilaianRisikoDeleteUseCaseRes{}, nil
	}
}
