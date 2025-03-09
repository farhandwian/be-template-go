package usecase

import (
	"context"
	"rmis/gateway"
	"shared/core"
	sharedModel "shared/model"
)

type PenilaianKegiatanPengendalianApprovalUseCaseReq struct {
	ID     string             `json:"id"`
	Status sharedModel.Status `json:"status"`
}

type PenilaianKegiatanPengendalianApprovalUseCaseRes struct{}

type PenilaianKegiatanPengendalianApprovalUseCase = core.ActionHandler[PenilaianKegiatanPengendalianApprovalUseCaseReq, PenilaianKegiatanPengendalianApprovalUseCaseRes]

func ImplPenilaianKegiatanPengendalianApprovalUseCase(
	getPenilaianKegiatanPengendalianById gateway.PenilaianKegiatanPengendalianGetByID,
	ApprovalPenilaianKegiatanPengendalian gateway.PenilaianKegiatanPengendalianSave,
) PenilaianKegiatanPengendalianApprovalUseCase {
	return func(ctx context.Context, req PenilaianKegiatanPengendalianApprovalUseCaseReq) (*PenilaianKegiatanPengendalianApprovalUseCaseRes, error) {

		res, err := getPenilaianKegiatanPengendalianById(ctx, gateway.PenilaianKegiatanPengendalianGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		penilaian := res.PenilaianKegiatanPengendalian

		penilaian.Status = req.Status
		if _, err := ApprovalPenilaianKegiatanPengendalian(ctx, gateway.PenilaianKegiatanPengendalianSaveReq{PenilaianKegiatanPengendalian: penilaian}); err != nil {
			return nil, err
		}

		return &PenilaianKegiatanPengendalianApprovalUseCaseRes{}, nil
	}
}
