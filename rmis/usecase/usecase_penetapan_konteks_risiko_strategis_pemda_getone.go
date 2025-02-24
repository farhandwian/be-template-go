package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenetapanKonteksRisikoGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoGetByIDUseCaseRes struct {
	PenetapanKonteksRisiko model.PenetapanKonteksRisikoStrategisPemda `json:"penetapan_konteks_risiko"`
	IKU                    []model.IKU                                `json:"ikus"`
}

type PenetapanKonteksRisikoGetByIDUseCase = core.ActionHandler[PenetapanKonteksRisikoGetByIDUseCaseReq, PenetapanKonteksRisikoGetByIDUseCaseRes]

func ImplPenetapanKonteksRisikoGetByIDUseCase(getPenetapanKonteksRisikoByID gateway.PenetapanKonteksRisikoStrategisPemdaGetByID, getAllIKUs gateway.IKUGetAll) PenetapanKonteksRisikoGetByIDUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoGetByIDUseCaseReq) (*PenetapanKonteksRisikoGetByIDUseCaseRes, error) {
		res, err := getPenetapanKonteksRisikoByID(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		ikus, err := getAllIKUs(ctx, gateway.IKUGetAllReq{
			ExternalID: *res.PenetapanKonteksRisikoStrategisPemda.ID,
		})

		if err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoGetByIDUseCaseRes{PenetapanKonteksRisiko: res.PenetapanKonteksRisikoStrategisPemda, IKU: ikus.IKU}, nil
	}
}
