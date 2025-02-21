package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type KategoriRisikoCreateUseCaseReq struct {
	Nama string `json:"nama"`
	Kode string `json:"kode"`
}

type KategoriRisikoCreateUseCaseRes struct {
	ID string `json:"id"`
}

type KategoriRisikoCreateUseCase = core.ActionHandler[KategoriRisikoCreateUseCaseReq, KategoriRisikoCreateUseCaseRes]

func ImplKategoriRisikoCreateUseCase(
	generateId gateway.GenerateId,
	createKategoriRisiko gateway.KategoriRisikoSave,
) KategoriRisikoCreateUseCase {
	return func(ctx context.Context, req KategoriRisikoCreateUseCaseReq) (*KategoriRisikoCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.KategoriRisiko{
			ID:   &genObj.RandomId,
			Nama: &req.Nama,
			Kode: &req.Kode,
		}

		if _, err = createKategoriRisiko(ctx, gateway.KategoriRisikoSaveReq{KategoriRisiko: obj}); err != nil {
			return nil, err
		}

		return &KategoriRisikoCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
