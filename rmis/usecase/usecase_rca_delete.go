package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type RcaDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type RcaDeleteUseCaseRes struct{}

type RcaDeleteUseCase = core.ActionHandler[RcaDeleteUseCaseReq, RcaDeleteUseCaseRes]

func ImplRcaDeleteUseCase(deleteRca gateway.RcaDelete) RcaDeleteUseCase {
	return func(ctx context.Context, req RcaDeleteUseCaseReq) (*RcaDeleteUseCaseRes, error) {

		if _, err := deleteRca(ctx, gateway.RcaDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &RcaDeleteUseCaseRes{}, nil
	}
}
