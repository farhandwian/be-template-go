package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenyebabRisikoGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PenyebabRisikoGetByIDUseCaseRes struct {
	PenyebabRisiko model.PenyebabRisiko `json:"PenyebabRisiko"`
}

type PenyebabRisikoGetByIDUseCase = core.ActionHandler[PenyebabRisikoGetByIDUseCaseReq, PenyebabRisikoGetByIDUseCaseRes]

func ImplPenyebabRisikoGetByIDUseCase(getPenyebabRisikoByID gateway.PenyebabRisikoGetByID) PenyebabRisikoGetByIDUseCase {
	return func(ctx context.Context, req PenyebabRisikoGetByIDUseCaseReq) (*PenyebabRisikoGetByIDUseCaseRes, error) {
		res, err := getPenyebabRisikoByID(ctx, gateway.PenyebabRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &PenyebabRisikoGetByIDUseCaseRes{PenyebabRisiko: res.PenyebabRisiko}, nil
	}
}
