package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type OPDUpdateUseCaseReq struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
	Kode string `json:"kode"`
}

type OPDUpdateUseCaseRes struct{}

type OPDUpdateUseCase = core.ActionHandler[OPDUpdateUseCaseReq, OPDUpdateUseCaseRes]

func ImplOPDUpdateUseCase(
	getOPDById gateway.OPDGetByID,
	updateOPD gateway.OPDSave,
) OPDUpdateUseCase {
	return func(ctx context.Context, req OPDUpdateUseCaseReq) (*OPDUpdateUseCaseRes, error) {

		res, err := getOPDById(ctx, gateway.OPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.OPD.Nama = &req.Nama

		if _, err := updateOPD(ctx, gateway.OPDSaveReq{OPD: res.OPD}); err != nil {
			return nil, err
		}

		return &OPDUpdateUseCaseRes{}, nil
	}
}
