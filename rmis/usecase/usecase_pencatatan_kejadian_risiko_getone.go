package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type PencatatanKejadianRisikoGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type PencatatanKejadianRisikoGetByIDUseCaseRes struct {
	PencatatanKejadianRisiko model.PencatatanKejadianRisiko `json:"pencatatan_kejadian_risiko"`
}

type PencatatanKejadianRisikoGetByIDUseCase = core.ActionHandler[PencatatanKejadianRisikoGetByIDUseCaseReq, PencatatanKejadianRisikoGetByIDUseCaseRes]

func ImplPencatatanKejadianRisikoGetByIDUseCase(getPencatatanKejadianRisikoByID gateway.PencatatanKejadianRisikoGetByID) PencatatanKejadianRisikoGetByIDUseCase {
	return func(ctx context.Context, req PencatatanKejadianRisikoGetByIDUseCaseReq) (*PencatatanKejadianRisikoGetByIDUseCaseRes, error) {
		res, err := getPencatatanKejadianRisikoByID(ctx, gateway.PencatatanKejadianRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &PencatatanKejadianRisikoGetByIDUseCaseRes{PencatatanKejadianRisiko: res.PencatatanKejadianRisiko}, nil
	}
}
