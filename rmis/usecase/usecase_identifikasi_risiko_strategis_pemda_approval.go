package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type IdentifikasiRisikoStrategisPemdaApprovalUseCaseReq struct {
	ID     string             `json:"id"`
	Status sharedModel.Status `json:"status"`
}

type IdentifikasiRisikoStrategisPemdaApprovalUseCaseRes struct{}

type IdentifikasiRisikoStrategisPemdaApprovalUseCase = core.ActionHandler[IdentifikasiRisikoStrategisPemdaApprovalUseCaseReq, IdentifikasiRisikoStrategisPemdaApprovalUseCaseRes]

func ImplIdentifikasiRisikoStrategisPemdaApprovalUseCase(
	getIdentifikasiRisikoStrategisPemdaById gateway.IdentifikasiRisikoStrategisPemdaGetByID,
	ApprovalIdentifikasiRisikoStrategisPemda gateway.IdentifikasiRisikoStrategisPemdaSave,
) IdentifikasiRisikoStrategisPemdaApprovalUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaApprovalUseCaseReq) (*IdentifikasiRisikoStrategisPemdaApprovalUseCaseRes, error) {

		res, err := getIdentifikasiRisikoStrategisPemdaById(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		identifikasiRisikoStrategisPemda := res.IdentifikasiRisikoStrategisPemda
		identifikasiRisikoStrategisPemda.Status = req.Status

		if _, err := ApprovalIdentifikasiRisikoStrategisPemda(ctx, gateway.IdentifikasiRisikoStrategisPemdaSaveReq{IdentifikasiRisikoStrategisPemda: identifikasiRisikoStrategisPemda}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisPemdaApprovalUseCaseRes{}, nil
	}
}
