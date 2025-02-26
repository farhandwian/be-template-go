package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IdentifikasiRisikoStrategisOPDGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoStrategisOPDGetByIDUseCaseRes struct {
	IdentifikasiRisikoStrategisOPD model.IdentifikasiRisikoStrategisOPD `json:"identifikasi_risiko_strategis_opd"`
	OPD                            model.OPD                            `json:"opd"`
}

type IdentifikasiRisikoStrategisOPDGetByIDUseCase = core.ActionHandler[IdentifikasiRisikoStrategisOPDGetByIDUseCaseReq, IdentifikasiRisikoStrategisOPDGetByIDUseCaseRes]

func ImplIdentifikasiRisikoStrategisOPDGetByIDUseCase(getIdentifikasiRisikoStrategisOPDByID gateway.IdentifikasiRisikoStrategisOPDGetByID, getOneOPD gateway.OPDGetByID) IdentifikasiRisikoStrategisOPDGetByIDUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDGetByIDUseCaseReq) (*IdentifikasiRisikoStrategisOPDGetByIDUseCaseRes, error) {
		res, err := getIdentifikasiRisikoStrategisOPDByID(ctx, gateway.IdentifikasiRisikoStrategisOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *res.IdentifikasiRisikoStrategisOPD.OPDID})
		if err != nil {
			return nil, err
		}
		return &IdentifikasiRisikoStrategisOPDGetByIDUseCaseRes{IdentifikasiRisikoStrategisOPD: res.IdentifikasiRisikoStrategisOPD, OPD: opd.OPD}, nil
	}
}
