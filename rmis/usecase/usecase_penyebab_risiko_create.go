package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenyebabRisikoCreateUseCaseReq struct {
	Nama string `json:"nama"`
}

type PenyebabRisikoCreateUseCaseRes struct {
	ID string `json:"id"`
}

type PenyebabRisikoCreateUseCase = core.ActionHandler[PenyebabRisikoCreateUseCaseReq, PenyebabRisikoCreateUseCaseRes]

func ImplPenyebabRisikoCreateUseCase(
	generateId gateway.GenerateId,
	createPenyebabRisiko gateway.PenyebabRisikoSave,
) PenyebabRisikoCreateUseCase {
	return func(ctx context.Context, req PenyebabRisikoCreateUseCaseReq) (*PenyebabRisikoCreateUseCaseRes, error) {

		// Generate a unique ID
		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.PenyebabRisiko{
			ID:   &genObj.RandomId,
			Nama: &req.Nama,
		}

		// Save the PenyebabRisiko entry
		if _, err = createPenyebabRisiko(ctx, gateway.PenyebabRisikoSaveReq{PenyebabRisiko: obj}); err != nil {
			return nil, err
		}

		return &PenyebabRisikoCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
