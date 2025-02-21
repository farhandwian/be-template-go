package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type KategoriRisikoGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type KategoriRisikoGetByIDUseCaseRes struct {
	KategoriRisiko model.KategoriRisiko `json:"KategoriRisiko"`
}

type KategoriRisikoGetByIDUseCase = core.ActionHandler[KategoriRisikoGetByIDUseCaseReq, KategoriRisikoGetByIDUseCaseRes]

func ImplKategoriRisikoGetByIDUseCase(getKategoriRisikoByID gateway.KategoriRisikoGetByID) KategoriRisikoGetByIDUseCase {
	return func(ctx context.Context, req KategoriRisikoGetByIDUseCaseReq) (*KategoriRisikoGetByIDUseCaseRes, error) {
		res, err := getKategoriRisikoByID(ctx, gateway.KategoriRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &KategoriRisikoGetByIDUseCaseRes{KategoriRisiko: res.KategoriRisiko}, nil
	}
}
