package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PenetapanKonteksRisikoRenstraOPDGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PenetapanKonteksRisikoRenstraOPDGetByIDUseCaseRes struct {
	PenetapanKonteksRisikoRenstraOPD model.PenetapanKonteksRisikoStrategisRenstraOPDGet `json:"penetapan_konteks_risiko_strategis_renstra_opd"`
}

type PenetapanKonteksRisikoRenstraOPDGetByIDUseCase = core.ActionHandler[PenetapanKonteksRisikoRenstraOPDGetByIDUseCaseReq, PenetapanKonteksRisikoRenstraOPDGetByIDUseCaseRes]

func ImplPenetapanKonteksRisikoRenstraOPDGetByIDUseCase(getPenetapanKonteksRisikoRenstraOPDByID gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID, getAllIKUs gateway.IKUGetAll, getOneOPD gateway.OPDGetByID) PenetapanKonteksRisikoRenstraOPDGetByIDUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoRenstraOPDGetByIDUseCaseReq) (*PenetapanKonteksRisikoRenstraOPDGetByIDUseCaseRes, error) {
		res, err := getPenetapanKonteksRisikoRenstraOPDByID(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		ikus, err := getAllIKUs(ctx, gateway.IKUGetAllReq{
			ExternalID: *res.PenetapanKonteksRisikoStrategisRenstraOPD.ID,
		})

		if err != nil {
			return nil, err
		}

		opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *res.PenetapanKonteksRisikoStrategisRenstraOPD.OPDID})
		if err != nil {
			return nil, err
		}

		var penetapanKonteksRisikoRes model.PenetapanKonteksRisikoStrategisRenstraOPDGet
		penetapanKonteksRisikoRes.PenetapanKonteksRisikoStrategisRenstraOPD = res.PenetapanKonteksRisikoStrategisRenstraOPD
		penetapanKonteksRisikoRes.IKUs = ikus.IKU
		penetapanKonteksRisikoRes.OPD = opd.OPD

		return &PenetapanKonteksRisikoRenstraOPDGetByIDUseCaseRes{PenetapanKonteksRisikoRenstraOPD: penetapanKonteksRisikoRes}, nil
	}
}
