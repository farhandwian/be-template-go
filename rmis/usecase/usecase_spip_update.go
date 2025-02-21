package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type SpipUpdateUseCaseReq struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

type SpipUpdateUseCaseRes struct{}

type SpipUpdateUseCase = core.ActionHandler[SpipUpdateUseCaseReq, SpipUpdateUseCaseRes]

func ImplSpipUpdateUseCase(
	getSpipById gateway.SpipGetByID,
	updateSpip gateway.SpipSave,
) SpipUpdateUseCase {
	return func(ctx context.Context, req SpipUpdateUseCaseReq) (*SpipUpdateUseCaseRes, error) {

		res, err := getSpipById(ctx, gateway.SpipGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.SPIP.Nama = &req.Nama

		if _, err := updateSpip(ctx, gateway.SpipSaveReq{Spip: res.SPIP}); err != nil {
			return nil, err
		}

		return &SpipUpdateUseCaseRes{}, nil
	}
}
