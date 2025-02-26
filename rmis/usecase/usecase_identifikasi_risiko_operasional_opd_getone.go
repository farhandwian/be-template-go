package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type IdentifikasiRisikoOperasionalOPDGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type IdentifikasiRisikoOperasionalOPDGetByIDUseCaseRes struct {
	IdentifikasiRisikoOperasionalOPD model.IdentifikasiRisikoOperasionalOPD `json:"identifikasi_risiko_operasional_opd"`
	OPD                              model.OPD                              `json:"opd"`
}

type IdentifikasiRisikoOperasionalOPDGetByIDUseCase = core.ActionHandler[IdentifikasiRisikoOperasionalOPDGetByIDUseCaseReq, IdentifikasiRisikoOperasionalOPDGetByIDUseCaseRes]

func ImplIdentifikasiRisikoOperasionalOPDGetByIDUseCase(getIdentifikasiRisikoOperasionalOPDByID gateway.IdentifikasiRisikoOperasionalOPDGetByID, getOneOPD gateway.OPDGetByID) IdentifikasiRisikoOperasionalOPDGetByIDUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDGetByIDUseCaseReq) (*IdentifikasiRisikoOperasionalOPDGetByIDUseCaseRes, error) {
		res, err := getIdentifikasiRisikoOperasionalOPDByID(ctx, gateway.IdentifikasiRisikoOperasionalOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *res.IdentifikasiRisikoOperasionalOPD.OPDID})
		if err != nil {
			return nil, err
		}
		return &IdentifikasiRisikoOperasionalOPDGetByIDUseCaseRes{IdentifikasiRisikoOperasionalOPD: res.IdentifikasiRisikoOperasionalOPD, OPD: opd.OPD}, nil
	}
}
