package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PengkomunikasianPengendalianGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
	Media     string
}

type PengkomunikasianPengendalianGetAllUseCaseRes struct {
	PengkomunikasianPengendalian []model.PengkomunikasianPengendalianResponse `json:"pengkomunikasian_pengendalian"`
	Metadata                     *usecase.Metadata                            `json:"metadata"`
}

type PengkomunikasianPengendalianGetAllUseCase = core.ActionHandler[PengkomunikasianPengendalianGetAllUseCaseReq, PengkomunikasianPengendalianGetAllUseCaseRes]

func ImplPengkomunikasianPengendalianGetAllUseCase(getAllPengkomunikasianPengendalians gateway.PengkomunikasianPengendalianGetAll) PengkomunikasianPengendalianGetAllUseCase {
	return func(ctx context.Context, req PengkomunikasianPengendalianGetAllUseCaseReq) (*PengkomunikasianPengendalianGetAllUseCaseRes, error) {

		res, err := getAllPengkomunikasianPengendalians(ctx, gateway.PengkomunikasianPengendalianGetAllReq{
			Page:      req.Page,
			Size:      req.Size,
			Keyword:   req.Keyword,
			SortBy:    req.SortBy,
			SortOrder: req.SortOrder,
			Status:    req.Status,
			Media:     req.Media,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &PengkomunikasianPengendalianGetAllUseCaseRes{
			PengkomunikasianPengendalian: res.PengkomunikasianPengendalian,
			Metadata: &usecase.Metadata{
				Pagination: usecase.Pagination{
					Page:       req.Page,
					Limit:      req.Size,
					TotalPages: totalPages,
					TotalItems: totalItems,
				},
			},
		}, nil
	}
}
