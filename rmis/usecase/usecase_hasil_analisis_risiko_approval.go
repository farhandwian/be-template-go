package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type HasilAnalisisRisikoApprovalUseCaseReq struct {
	ID     string             `json:"-"`
	Status sharedModel.Status `json:"status"`
}

type HasilAnalisisRisikoApprovalUseCaseRes struct{}

type HasilAnalisisRisikoApprovalUseCase = core.ActionHandler[HasilAnalisisRisikoApprovalUseCaseReq, HasilAnalisisRisikoApprovalUseCaseRes]

func ImplHasilAnalisisRisikoApprovalUseCase(
	getHasilAnalisisRisikoById gateway.HasilAnalisisRisikoGetByID,
	ApprovalHasilAnalisisRisiko gateway.HasilAnalisisRisikoSave,
) HasilAnalisisRisikoApprovalUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoApprovalUseCaseReq) (*HasilAnalisisRisikoApprovalUseCaseRes, error) {

		res, err := getHasilAnalisisRisikoById(ctx, gateway.HasilAnalisisRisikoGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		hasilAnalisisRisiko := res.HasilAnalisisRisiko
		hasilAnalisisRisiko.Status = req.Status

		if _, err := ApprovalHasilAnalisisRisiko(ctx, gateway.HasilAnalisisRisikoSaveReq{HasilAnalisisRisiko: hasilAnalisisRisiko}); err != nil {
			return nil, err
		}

		return &HasilAnalisisRisikoApprovalUseCaseRes{}, nil
	}
}
