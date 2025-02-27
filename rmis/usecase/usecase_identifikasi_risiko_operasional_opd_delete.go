package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type IdentifikasiRisikoOperasionalOPDDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoOperasionalOPDDeleteUseCaseRes struct{}

type IdentifikasiRisikoOperasionalOPDDeleteUseCase = core.ActionHandler[IdentifikasiRisikoOperasionalOPDDeleteUseCaseReq, IdentifikasiRisikoOperasionalOPDDeleteUseCaseRes]

func ImplIdentifikasiRisikoOperasionalOPDDeleteUseCase(deleteIdentifikasiRisikoOperasionalOPD gateway.IdentifikasiRisikoOperasionalOPDDelete) IdentifikasiRisikoOperasionalOPDDeleteUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDDeleteUseCaseReq) (*IdentifikasiRisikoOperasionalOPDDeleteUseCaseRes, error) {

		if _, err := deleteIdentifikasiRisikoOperasionalOPD(ctx, gateway.IdentifikasiRisikoOperasionalOPDDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoOperasionalOPDDeleteUseCaseRes{}, nil
	}
}
