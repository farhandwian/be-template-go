package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type IdentifikasiRisikoStrategisOPDDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoStrategisOPDDeleteUseCaseRes struct{}

type IdentifikasiRisikoStrategisOPDDeleteUseCase = core.ActionHandler[IdentifikasiRisikoStrategisOPDDeleteUseCaseReq, IdentifikasiRisikoStrategisOPDDeleteUseCaseRes]

func ImplIdentifikasiRisikoStrategisOPDDeleteUseCase(deleteIdentifikasiRisikoStrategisOPD gateway.IdentifikasiRisikoStrategisOPDDelete) IdentifikasiRisikoStrategisOPDDeleteUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDDeleteUseCaseReq) (*IdentifikasiRisikoStrategisOPDDeleteUseCaseRes, error) {

		if _, err := deleteIdentifikasiRisikoStrategisOPD(ctx, gateway.IdentifikasiRisikoStrategisOPDDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisOPDDeleteUseCaseRes{}, nil
	}
}
