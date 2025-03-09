package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseReq struct {
	ID     string             `json:"id"`
	Status sharedModel.Status `json:"status"`
}

type PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseRes struct{}

type PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCase = core.ActionHandler[PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseReq, PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseRes]

func ImplPenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCase(
	getPenetapanKonteksRisikoStrategisRenstraOPDById gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByID,
	ApprovalPenetapanKonteksRisikoStrategisRenstraOPD gateway.PenetepanKonteksRisikoStrategisRenstraOPDSave,
	OpdByID gateway.OPDGetByID,
) PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseReq) (*PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseRes, error) {

		res, err := getPenetapanKonteksRisikoStrategisRenstraOPDById(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		penetapanKonteksRisikoStrategisRenstraOPD := res.PenetapanKonteksRisikoStrategisRenstraOPD

		penetapanKonteksRisikoStrategisRenstraOPD.Status = req.Status

		if _, err := ApprovalPenetapanKonteksRisikoStrategisRenstraOPD(ctx, gateway.PenetapanKonteksRisikoStrategisRenstraOPDSaveReq{PenetepanKonteksRisikoStrategisRenstraOPD: penetapanKonteksRisikoStrategisRenstraOPD}); err != nil {
			return nil, err
		}

		return &PenetapanKonteksRisikoStrategisRenstraOPDApprovalUseCaseRes{}, nil
	}
}
