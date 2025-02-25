package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
)

type DaftarRisikoPrioritasGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type DaftarRisikoPrioritasGetByIDUseCaseRes struct {
	DaftarRisikoPrioritas model.DaftarRisikoPrioritas `json:"hasil_analisis_risiko"`
}

type DaftarRisikoPrioritasGetByIDUseCase = core.ActionHandler[DaftarRisikoPrioritasGetByIDUseCaseReq, DaftarRisikoPrioritasGetByIDUseCaseRes]

func ImplDaftarRisikoPrioritasGetByIDUseCase(getDaftarRisikoPrioritasByID gateway.DaftarRisikoPrioritasGetByID) DaftarRisikoPrioritasGetByIDUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasGetByIDUseCaseReq) (*DaftarRisikoPrioritasGetByIDUseCaseRes, error) {
		res, err := getDaftarRisikoPrioritasByID(ctx, gateway.DaftarRisikoPrioritasGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &DaftarRisikoPrioritasGetByIDUseCaseRes{DaftarRisikoPrioritas: res.DaftarRisikoPrioritas}, nil
	}
}
