package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IdentifikasiRisikoStrategisPemdaGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoStrategisPemdaGetByIDUseCaseRes struct {
	IdentifikasiRisikoStrategisPemda model.IdentifikasiRisikoStrategisPemda `json:"identifikasi_risiko_strategis_pemda"`
}

type IdentifikasiRisikoStrategisPemdaGetByIDUseCase = core.ActionHandler[IdentifikasiRisikoStrategisPemdaGetByIDUseCaseReq, IdentifikasiRisikoStrategisPemdaGetByIDUseCaseRes]

func ImplIdentifikasiRisikoStrategisPemdaGetByIDUseCase(getIdentifikasiRisikoStrategisPemdaByID gateway.IdentifikasiRisikoStrategisPemdaGetByID) IdentifikasiRisikoStrategisPemdaGetByIDUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaGetByIDUseCaseReq) (*IdentifikasiRisikoStrategisPemdaGetByIDUseCaseRes, error) {
		res, err := getIdentifikasiRisikoStrategisPemdaByID(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.ID})

		if err != nil {
			return nil, err
		}
		return &IdentifikasiRisikoStrategisPemdaGetByIDUseCaseRes{IdentifikasiRisikoStrategisPemda: res.IdentifikasiRisikoStrategisPemda}, nil
	}
}
