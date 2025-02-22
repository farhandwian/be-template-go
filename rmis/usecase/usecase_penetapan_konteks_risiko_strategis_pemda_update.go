package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes struct{}

type PenetapanKonteksRisikoStrategisPemdaUpdateUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq, PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisPemdaUpdateUseCase(
	getPenetapanKonteksRisikoStrategisPemdaById gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	updatePenetapanKonteksRisikoStrategisPemda gateway.PenetepanKonteksRisikoStrategisPemdaSave,
) PenetapanKonteksRisikoStrategisPemdaUpdateUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseReq) (*PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoStrategisPemdaById(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.PenetapanKonteksRisikoStrategisPemda.Nama = &req.Nama

		if _, err := updatePenetapanKonteksRisikoStrategisPemda(ctx, gateway.PenetepanKonteksRisikoStrategisPemdaSaveReq{PenetapanKonteksRisikoStrategisPemda: res.PenetapanKonteksRisikoStrategisPemda}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisPemdaUpdateUseCaseRes{}, nil
	}
}
