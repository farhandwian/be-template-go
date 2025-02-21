package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type KategoriRisikoUpdateUseCaseReq struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
	Kode string `json:"kode"`
}

type KategoriRisikoUpdateUseCaseRes struct{}

type KategoriRisikoUpdateUseCase = core.ActionHandler[KategoriRisikoUpdateUseCaseReq, KategoriRisikoUpdateUseCaseRes]

func ImplKategoriRisikoUpdateUseCase(
	getKategoriRisikoById gateway.KategoriRisikoGetByID,
	updateKategoriRisiko gateway.KategoriRisikoSave,
) KategoriRisikoUpdateUseCase {
	return func(ctx context.Context, req KategoriRisikoUpdateUseCaseReq) (*KategoriRisikoUpdateUseCaseRes, error) {

		res, err := getKategoriRisikoById(ctx, gateway.KategoriRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.KategoriRisiko.Nama = &req.Nama
		res.KategoriRisiko.Kode = &req.Kode

		if _, err := updateKategoriRisiko(ctx, gateway.KategoriRisikoSaveReq{KategoriRisiko: res.KategoriRisiko}); err != nil {
			return nil, err
		}

		return &KategoriRisikoUpdateUseCaseRes{}, nil
	}
}
