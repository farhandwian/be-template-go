package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type DaftarRisikoPrioritasApprovalUseCaseReq struct {
	ID     string             `json:"id"`
	Status sharedModel.Status `json:"status"`
}

type DaftarRisikoPrioritasApprovalUseCaseRes struct{}

type DaftarRisikoPrioritasApprovalUseCase = core.ActionHandler[DaftarRisikoPrioritasApprovalUseCaseReq, DaftarRisikoPrioritasApprovalUseCaseRes]

func ImplDaftarRisikoPrioritasApprovalUseCase(
	getDaftarRisikoPrioritasById gateway.DaftarRisikoPrioritasGetByID,
	ApprovalDaftarRisikoPrioritas gateway.DaftarRisikoPrioritasSave,
) DaftarRisikoPrioritasApprovalUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasApprovalUseCaseReq) (*DaftarRisikoPrioritasApprovalUseCaseRes, error) {

		res, err := getDaftarRisikoPrioritasById(ctx, gateway.DaftarRisikoPrioritasGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		daftarRisikoPrioritas := res.DaftarRisikoPrioritas
		daftarRisikoPrioritas.Status = req.Status

		if _, err := ApprovalDaftarRisikoPrioritas(ctx, gateway.DaftarRisikoPrioritasSaveReq{DaftarRisikoPrioritas: daftarRisikoPrioritas}); err != nil {
			return nil, err
		}

		return &DaftarRisikoPrioritasApprovalUseCaseRes{}, nil
	}
}
