package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type HasilAnalisisRisikoGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type HasilAnalisisRisikoGetByIDUseCaseRes struct {
	HasilAnalisisRisiko model.HasilAnalisisRisiko `json:"hasil_analisis_risiko"`
}

type HasilAnalisisRisikoGetByIDUseCase = core.ActionHandler[HasilAnalisisRisikoGetByIDUseCaseReq, HasilAnalisisRisikoGetByIDUseCaseRes]

func ImplHasilAnalisisRisikoGetByIDUseCase(getHasilAnalisisRisikoByID gateway.HasilAnalisisRisikoGetByID) HasilAnalisisRisikoGetByIDUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoGetByIDUseCaseReq) (*HasilAnalisisRisikoGetByIDUseCaseRes, error) {
		res, err := getHasilAnalisisRisikoByID(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &HasilAnalisisRisikoGetByIDUseCaseRes{HasilAnalisisRisiko: res.HasilAnalisisRisiko}, nil
	}
}
