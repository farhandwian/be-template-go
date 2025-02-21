package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type SpipDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type SpipDeleteUseCaseRes struct{}

type SpipDeleteUseCase = core.ActionHandler[SpipDeleteUseCaseReq, SpipDeleteUseCaseRes]

func ImplSpipDeleteUseCase(deleteSpip gateway.SpipDelete) SpipDeleteUseCase {
	return func(ctx context.Context, req SpipDeleteUseCaseReq) (*SpipDeleteUseCaseRes, error) {

		if _, err := deleteSpip(ctx, gateway.SpipDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &SpipDeleteUseCaseRes{}, nil
	}
}
