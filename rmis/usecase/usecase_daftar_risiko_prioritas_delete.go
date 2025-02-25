package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
)

type DaftarRisikoPrioritasDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type DaftarRisikoPrioritasDeleteUseCaseRes struct{}

type DaftarRisikoPrioritasDeleteUseCase = core.ActionHandler[DaftarRisikoPrioritasDeleteUseCaseReq, DaftarRisikoPrioritasDeleteUseCaseRes]

func ImplDaftarRisikoPrioritasDeleteUseCase(deleteDaftarRisikoPrioritas gateway.DaftarRisikoPrioritasDelete) DaftarRisikoPrioritasDeleteUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasDeleteUseCaseReq) (*DaftarRisikoPrioritasDeleteUseCaseRes, error) {

		if _, err := deleteDaftarRisikoPrioritas(ctx, gateway.DaftarRisikoPrioritasDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &DaftarRisikoPrioritasDeleteUseCaseRes{}, nil
	}
}
