package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type KategoriRisikoDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type KategoriRisikoDeleteUseCaseRes struct{}

type KategoriRisikoDeleteUseCase = core.ActionHandler[KategoriRisikoDeleteUseCaseReq, KategoriRisikoDeleteUseCaseRes]

func ImplKategoriRisikoDeleteUseCase(deleteKategoriRisiko gateway.KategoriRisikoDelete) KategoriRisikoDeleteUseCase {
	return func(ctx context.Context, req KategoriRisikoDeleteUseCaseReq) (*KategoriRisikoDeleteUseCaseRes, error) {

		if _, err := deleteKategoriRisiko(ctx, gateway.KategoriRisikoDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &KategoriRisikoDeleteUseCaseRes{}, nil
	}
}
