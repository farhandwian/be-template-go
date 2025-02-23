package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type RcaCreateUseCaseReq struct {
	Nama string `json:"nama"`
}

type RcaCreateUseCaseRes struct {
	ID string `json:"id"`
}

type RcaCreateUseCase = core.ActionHandler[RcaCreateUseCaseReq, RcaCreateUseCaseRes]

func ImplRcaCreateUseCase(
	generateId gateway.GenerateId,
	createRca gateway.RcaSave,
) RcaCreateUseCase {
	return func(ctx context.Context, req RcaCreateUseCaseReq) (*RcaCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.Rca{
			ID: &genObj.RandomId,
		}

		if _, err = createRca(ctx, gateway.RcaSaveReq{Rca: obj}); err != nil {
			return nil, err
		}

		return &RcaCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
