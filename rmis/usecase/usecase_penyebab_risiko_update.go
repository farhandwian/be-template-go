package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenyebabRisikoUpdateUseCaseReq struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

type PenyebabRisikoUpdateUseCaseRes struct{}

type PenyebabRisikoUpdateUseCase = core.ActionHandler[PenyebabRisikoUpdateUseCaseReq, PenyebabRisikoUpdateUseCaseRes]

func ImplPenyebabRisikoUpdateUseCase(
	getPenyebabRisikoById gateway.PenyebabRisikoGetByID,
	updatePenyebabRisiko gateway.PenyebabRisikoSave,
) PenyebabRisikoUpdateUseCase {
	return func(ctx context.Context, req PenyebabRisikoUpdateUseCaseReq) (*PenyebabRisikoUpdateUseCaseRes, error) {

		res, err := getPenyebabRisikoById(ctx, gateway.PenyebabRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		PenyebabRisiko := res.PenyebabRisiko
		PenyebabRisiko.Nama = &req.Nama

		if _, err := updatePenyebabRisiko(ctx, gateway.PenyebabRisikoSaveReq{PenyebabRisiko: PenyebabRisiko}); err != nil {
			return nil, err
		}

		return &PenyebabRisikoUpdateUseCaseRes{}, nil
	}
}
