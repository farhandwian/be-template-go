package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type OPDGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type OPDGetByIDUseCaseRes struct {
	OPD model.OPD `json:"opd"`
}

type OPDGetByIDUseCase = core.ActionHandler[OPDGetByIDUseCaseReq, OPDGetByIDUseCaseRes]

func ImplOPDGetByIDUseCase(getOPDByID gateway.OPDGetByID) OPDGetByIDUseCase {
	return func(ctx context.Context, req OPDGetByIDUseCaseReq) (*OPDGetByIDUseCaseRes, error) {
		res, err := getOPDByID(ctx, gateway.OPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &OPDGetByIDUseCaseRes{OPD: res.OPD}, nil
	}
}
