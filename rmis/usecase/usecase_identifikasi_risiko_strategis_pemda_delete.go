package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type IdentifikasiRisikoStrategisPemdaDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoStrategisPemdaDeleteUseCaseRes struct{}

type IdentifikasiRisikoStrategisPemdaDeleteUseCase = core.ActionHandler[IdentifikasiRisikoStrategisPemdaDeleteUseCaseReq, IdentifikasiRisikoStrategisPemdaDeleteUseCaseRes]

func ImplIdentifikasiRisikoStrategisPemdaDeleteUseCase(deleteIdentifikasiRisikoStrategisPemda gateway.IdentifikasiRisikoStrategisPemdaDelete) IdentifikasiRisikoStrategisPemdaDeleteUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaDeleteUseCaseReq) (*IdentifikasiRisikoStrategisPemdaDeleteUseCaseRes, error) {

		if _, err := deleteIdentifikasiRisikoStrategisPemda(ctx, gateway.IdentifikasiRisikoStrategisPemdaDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisPemdaDeleteUseCaseRes{}, nil
	}
}
