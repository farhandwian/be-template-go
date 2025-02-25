package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type OPDDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type OPDDeleteUseCaseRes struct{}

type OPDDeleteUseCase = core.ActionHandler[OPDDeleteUseCaseReq, OPDDeleteUseCaseRes]

func ImplOPDDeleteUseCase(deleteOPD gateway.OPDDelete) OPDDeleteUseCase {
	return func(ctx context.Context, req OPDDeleteUseCaseReq) (*OPDDeleteUseCaseRes, error) {

		if _, err := deleteOPD(ctx, gateway.OPDDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &OPDDeleteUseCaseRes{}, nil
	}
}
