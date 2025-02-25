package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenilaianRisikoGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PenilaianRisikoGetByIDUseCaseRes struct {
	PenilaianRisiko model.PenilaianRisiko `json:"PenilaianRisiko"`
}

type PenilaianRisikoGetByIDUseCase = core.ActionHandler[PenilaianRisikoGetByIDUseCaseReq, PenilaianRisikoGetByIDUseCaseRes]

func ImplPenilaianRisikoGetByIDUseCase(getPenilaianRisikoByID gateway.PenilaianRisikoGetByID) PenilaianRisikoGetByIDUseCase {
	return func(ctx context.Context, req PenilaianRisikoGetByIDUseCaseReq) (*PenilaianRisikoGetByIDUseCaseRes, error) {
		res, err := getPenilaianRisikoByID(ctx, gateway.PenilaianRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &PenilaianRisikoGetByIDUseCaseRes{PenilaianRisiko: res.PenilaianRisiko}, nil
	}
}
