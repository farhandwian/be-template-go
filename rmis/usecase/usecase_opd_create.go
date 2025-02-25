package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type OPDCreateUseCaseReq struct {
	Nama string `json:"nama"`
	Kode string `json:"kode"`
}

type OPDCreateUseCaseRes struct {
	ID string `json:"id"`
}

type OPDCreateUseCase = core.ActionHandler[OPDCreateUseCaseReq, OPDCreateUseCaseRes]

func ImplOPDCreateUseCase(
	generateId gateway.GenerateId,
	createOPD gateway.OPDSave,
) OPDCreateUseCase {
	return func(ctx context.Context, req OPDCreateUseCaseReq) (*OPDCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.OPD{
			ID:   &genObj.RandomId,
			Nama: &req.Nama,
		}

		if _, err = createOPD(ctx, gateway.OPDSaveReq{OPD: obj}); err != nil {
			return nil, err
		}

		return &OPDCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
