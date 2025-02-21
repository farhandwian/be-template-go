package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type SpipGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type SpipGetByIDUseCaseRes struct {
	Spip model.Spip `json:"spip"`
}

type SpipGetByIDUseCase = core.ActionHandler[SpipGetByIDUseCaseReq, SpipGetByIDUseCaseRes]

func ImplSpipGetByIDUseCase(getSpipByID gateway.SpipGetByID) SpipGetByIDUseCase {
	return func(ctx context.Context, req SpipGetByIDUseCaseReq) (*SpipGetByIDUseCaseRes, error) {
		res, err := getSpipByID(ctx, gateway.SpipGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &SpipGetByIDUseCaseRes{Spip: res.SPIP}, nil
	}
}
