package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type IKUDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type IKUDeleteUseCaseRes struct{}

type IKUDeleteUseCase = core.ActionHandler[IKUDeleteUseCaseReq, IKUDeleteUseCaseRes]

func ImplIKUDeleteUseCase(deleteIKU gateway.IKUDelete) IKUDeleteUseCase {
	return func(ctx context.Context, req IKUDeleteUseCaseReq) (*IKUDeleteUseCaseRes, error) {

		if _, err := deleteIKU(ctx, gateway.IKUDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &IKUDeleteUseCaseRes{}, nil
	}
}
