package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type IdentifikasiRisikoStrategisOPDApprovalUseCaseReq struct {
	ID     string             `json:"-"`
	Status sharedModel.Status `json:"status"`
}

type IdentifikasiRisikoStrategisOPDApprovalUseCaseRes struct{}

type IdentifikasiRisikoStrategisOPDApprovalUseCase = core.ActionHandler[IdentifikasiRisikoStrategisOPDApprovalUseCaseReq, IdentifikasiRisikoStrategisOPDApprovalUseCaseRes]

func ImplIdentifikasiRisikoStrategisOPDApprovalUseCase(
	getIdentifikasiRisikoStrategisOPDById gateway.IdentifikasiRisikoStrategisOPDGetByID,
	ApprovalIdentifikasiRisikoStrategisOPD gateway.IdentifikasiRisikoStrategisOPDSave,
) IdentifikasiRisikoStrategisOPDApprovalUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDApprovalUseCaseReq) (*IdentifikasiRisikoStrategisOPDApprovalUseCaseRes, error) {

		res, err := getIdentifikasiRisikoStrategisOPDById(ctx, gateway.IdentifikasiRisikoStrategisOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.IdentifikasiRisikoStrategisOPD.Status = req.Status

		if _, err := ApprovalIdentifikasiRisikoStrategisOPD(ctx, gateway.IdentifikasiRisikoStrategisOPDSaveReq{IdentifikasiRisikoStrategisOPD: res.IdentifikasiRisikoStrategisOPD}); err != nil {
			return nil, err
		}

		return &IdentifikasiRisikoStrategisOPDApprovalUseCaseRes{}, nil
	}
}
