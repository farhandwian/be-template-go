package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PenetapanKonteksRisikoOperasionalApprovalUseCaseReq struct {
	ID     string             `json:"id"`
	Status sharedModel.Status `json:"status"`
}

type PenetapanKonteksRisikoOperasionalApprovalUseCaseRes struct{}

type PenetapanKonteksRisikoOperasionalApprovalUseCase = core.ActionHandler[PenetapanKonteksRisikoOperasionalApprovalUseCaseReq, PenetapanKonteksRisikoOperasionalApprovalUseCaseRes]

func ImplPenetapanKonteksRisikoOperasionalApprovalUseCase(
	getPenetapanKonteksRisikoOperasionalById gateway.PenetapanKonteksRisikoOperasionalGetByID,
	ApprovalPenetapanKonteksRisikoOperasional gateway.PenetepanKonteksRisikoOperasionalSave,
) PenetapanKonteksRisikoOperasionalApprovalUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoOperasionalApprovalUseCaseReq) (*PenetapanKonteksRisikoOperasionalApprovalUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoOperasionalById(ctx, gateway.PenetapanKonteksRisikoOperasionalGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.PenetapanKonteksRisikoOperasional.Status = req.Status

		if _, err := ApprovalPenetapanKonteksRisikoOperasional(ctx, gateway.PenetapanKonteksRisikoOperasionalSaveReq{PenetepanKonteksRisikoOperasional: res.PenetapanKonteksRisikoOperasional}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoOperasionalApprovalUseCaseRes{}, nil
	}
}
