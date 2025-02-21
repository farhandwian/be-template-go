package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type SpipCreateUseCaseReq struct {
	Nama string `json:"nama"`
}

type SpipCreateUseCaseRes struct {
	ID string `json:"id"`
}

type SpipCreateUseCase = core.ActionHandler[SpipCreateUseCaseReq, SpipCreateUseCaseRes]

func ImplSpipCreateUseCase(
	generateId gateway.GenerateId,
	createSpip gateway.SpipSave,
) SpipCreateUseCase {
	return func(ctx context.Context, req SpipCreateUseCaseReq) (*SpipCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.Spip{
			ID:   &genObj.RandomId,
			Nama: &req.Nama,
		}

		if _, err = createSpip(ctx, gateway.SpipSaveReq{Spip: obj}); err != nil {
			return nil, err
		}

		return &SpipCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
