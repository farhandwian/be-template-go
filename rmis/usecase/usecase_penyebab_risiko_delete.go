package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenyebabRisikoDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type PenyebabRisikoDeleteUseCaseRes struct{}

type PenyebabRisikoDeleteUseCase = core.ActionHandler[PenyebabRisikoDeleteUseCaseReq, PenyebabRisikoDeleteUseCaseRes]

func ImplPenyebabRisikoDeleteUseCase(deletePenyebabRisiko gateway.PenyebabRisikoDelete) PenyebabRisikoDeleteUseCase {
	return func(ctx context.Context, req PenyebabRisikoDeleteUseCaseReq) (*PenyebabRisikoDeleteUseCaseRes, error) {

		if _, err := deletePenyebabRisiko(ctx, gateway.PenyebabRisikoDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &PenyebabRisikoDeleteUseCaseRes{}, nil
	}
}
