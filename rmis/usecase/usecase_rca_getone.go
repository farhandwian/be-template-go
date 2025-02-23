package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type RcaGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type RcaGetByIDUseCaseRes struct {
	Rca model.Rca `json:"rca"`
}

type RcaGetByIDUseCase = core.ActionHandler[RcaGetByIDUseCaseReq, RcaGetByIDUseCaseRes]

func ImplRcaGetByIDUseCase(getRcaByID gateway.RcaGetByID) RcaGetByIDUseCase {
	return func(ctx context.Context, req RcaGetByIDUseCaseReq) (*RcaGetByIDUseCaseRes, error) {
		res, err := getRcaByID(ctx, gateway.RcaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &RcaGetByIDUseCaseRes{Rca: res.Rca}, nil
	}
}
