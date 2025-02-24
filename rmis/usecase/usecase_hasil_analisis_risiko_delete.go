package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type HasilAnalisisRisikoDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type HasilAnalisisRisikoDeleteUseCaseRes struct{}

type HasilAnalisisRisikoDeleteUseCase = core.ActionHandler[HasilAnalisisRisikoDeleteUseCaseReq, HasilAnalisisRisikoDeleteUseCaseRes]

func ImplHasilAnalisisRisikoDeleteUseCase(deleteHasilAnalisisRisiko gateway.HasilAnalisisRisikoDelete) HasilAnalisisRisikoDeleteUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoDeleteUseCaseReq) (*HasilAnalisisRisikoDeleteUseCaseRes, error) {

		if _, err := deleteHasilAnalisisRisiko(ctx, gateway.HasilAnalisisRisikoDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &HasilAnalisisRisikoDeleteUseCaseRes{}, nil
	}
}
