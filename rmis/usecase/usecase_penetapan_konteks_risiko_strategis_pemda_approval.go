package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseReq struct {
	ID     string             `json:"id"`
	Status sharedModel.Status `json:"status"`
}

type PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseRes struct{}

type PenetapanKonteksRisikoStrategisPemdaApprovalUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseReq, PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisPemdaApprovalUseCase(
	getPenetapanKonteksRisikoStrategisPemdaById gateway.PenetapanKonteksRisikoStrategisPemdaGetByID,
	ApprovalPenetapanKonteksRisikoStrategisPemda gateway.PenetapanKonteksRisikoStrategisPemdaSave,
) PenetapanKonteksRisikoStrategisPemdaApprovalUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseReq) (*PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoStrategisPemdaById(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		penetapanKonteksRisikoStrategisPemda := res.PenetapanKonteksRisikoStrategisPemda

		penetapanKonteksRisikoStrategisPemda.Status = req.Status

		if _, err := ApprovalPenetapanKonteksRisikoStrategisPemda(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaSaveReq{PenetepanKonteksRisikoStrategisPemda: penetapanKonteksRisikoStrategisPemda}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisPemdaApprovalUseCaseRes{}, nil
	}
}
