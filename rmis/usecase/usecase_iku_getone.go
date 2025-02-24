package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IKUGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type IKUGetByIDUseCaseRes struct {
	IKU model.IKU `json:"iku"`
}

type IKUGetByIDUseCase = core.ActionHandler[IKUGetByIDUseCaseReq, IKUGetByIDUseCaseRes]

func ImplIKUGetByIDUseCase(getIKUByID gateway.IKUGetByID) IKUGetByIDUseCase {
	return func(ctx context.Context, req IKUGetByIDUseCaseReq) (*IKUGetByIDUseCaseRes, error) {
		res, err := getIKUByID(ctx, gateway.IKUGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &IKUGetByIDUseCaseRes{IKU: res.IKU}, nil
	}
}
