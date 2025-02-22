package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenetapanKonteksRisikoGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoGetByIDUseCaseRes struct {
	PenetapanKonteksRisiko model.PenetapanKonteksRisikoStrategisPemda `json:"rekapitulasi_hasil_kuesioner"`
}

type PenetapanKonteksRisikoGetByIDUseCase = core.ActionHandler[PenetapanKonteksRisikoGetByIDUseCaseReq, PenetapanKonteksRisikoGetByIDUseCaseRes]

func ImplPenetapanKonteksRisikoGetByIDUseCase(getPenetapanKonteksRisikoByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID) PenetapanKonteksRisikoGetByIDUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoGetByIDUseCaseReq) (*PenetapanKonteksRisikoGetByIDUseCaseRes, error) {
		res, err := getPenetapanKonteksRisikoByID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &PenetapanKonteksRisikoGetByIDUseCaseRes{PenetapanKonteksRisiko: res.PenetapanKonteksRisikoStrategisPemda}, nil
	}
}
