package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenetapanKonteksRisikoOperasionalGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoOperasionalGetByIDUseCaseRes struct {
	PenetapanKonteksRisikoOperasional model.PenetapanKonteksRisikoOperasional `json:"penetapan_konteks_risiko_operasional"`
	IKU                               []model.IKU                             `json:"ikus"`
	OPD                               model.OPD                               `json:"opd"`
}

type PenetapanKonteksRisikoOperasionalGetByIDUseCase = core.ActionHandler[PenetapanKonteksRisikoOperasionalGetByIDUseCaseReq, PenetapanKonteksRisikoOperasionalGetByIDUseCaseRes]

func ImplPenetapanKonteksRisikoOperasionalGetByIDUseCase(getPenetapanKonteksRisikoOperasionalByID gateway.PenetapanKonteksRisikoOperasionalGetByID, getAllIKUs gateway.IKUGetAll, getOneOPD gateway.OPDGetByID) PenetapanKonteksRisikoOperasionalGetByIDUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalGetByIDUseCaseReq) (*PenetapanKonteksRisikoOperasionalGetByIDUseCaseRes, error) {
		res, err := getPenetapanKonteksRisikoOperasionalByID(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		ikus, err := getAllIKUs(ctx, gateway.IKUGetAllReq{
			ExternalID: *res.PenetapanKonteksRisikoOperasional.ID,
		})

		opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *res.PenetapanKonteksRisikoOperasional.OpdID})
		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoOperasionalGetByIDUseCaseRes{PenetapanKonteksRisikoOperasional: res.PenetapanKonteksRisikoOperasional, IKU: ikus.IKU, OPD: opd.OPD}, nil
	}
}
