package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type IdentifikasiRisikoOperasionalOPDApprovalUseCaseReq struct {
	ID     string             `json:"id"`
	Status sharedModel.Status `json:"status"`
}

type IdentifikasiRisikoOperasionalOPDApprovalUseCaseRes struct{}

type IdentifikasiRisikoOperasionalOPDApprovalUseCase = core.ActionHandler[IdentifikasiRisikoOperasionalOPDApprovalUseCaseReq, IdentifikasiRisikoOperasionalOPDApprovalUseCaseRes]

func ImplIdentifikasiRisikoOperasionalOPDApprovalUseCase(
	getIdentifikasiRisikoOperasionalOPDById gateway.IdentifikasiRisikoOperasionalOPDGetByID,
	ApprovalIdentifikasiRisikoOperasionalOPD gateway.IdentifikasiRisikoOperasionalOPDSave,
) IdentifikasiRisikoOperasionalOPDApprovalUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDApprovalUseCaseReq) (*IdentifikasiRisikoOperasionalOPDApprovalUseCaseRes, error) {

		res, err := getIdentifikasiRisikoOperasionalOPDById(ctx, gateway.IdentifikasiRisikoOperasionalOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		identifikasiOperasionalOpd := res.IdentifikasiRisikoOperasionalOPD
		identifikasiOperasionalOpd.Status = req.Status

		if _, err := ApprovalIdentifikasiRisikoOperasionalOPD(ctx, gateway.IdentifikasiRisikoOperasionalOPDSaveReq{IdentifikasiRisikoOperasionalOPD: identifikasiOperasionalOpd}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoOperasionalOPDApprovalUseCaseRes{}, nil
	}
}
