package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type RekapitulasiHasilKuesionerGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type RekapitulasiHasilKuesionerGetAllUseCaseRes struct {
	RekapitulasiHasilKuesioner []model.RekapitulasiHasilKuesioner `json:"RekapitulasiHasilKuesioners"`
	Metadata                   *usecase.Metadata                  `json:"metadata"`
}

type RekapitulasiHasilKuesionerGetAllUseCase = core.ActionHandler[RekapitulasiHasilKuesionerGetAllUseCaseReq, RekapitulasiHasilKuesionerGetAllUseCaseRes]

func ImplRekapitulasiHasilKuesionerGetAllUseCase(getAllRekapitulasiHasilKuesioners gateway.RekapitulasiHasilKuesionerGetAll) RekapitulasiHasilKuesionerGetAllUseCase {
	return func(ctx context.Context, req RekapitulasiHasilKuesionerGetAllUseCaseReq) (*RekapitulasiHasilKuesionerGetAllUseCaseRes, error) {

		res, err := getAllRekapitulasiHasilKuesioners(ctx, gateway.RekapitulasiHasilKuesionerGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &RekapitulasiHasilKuesionerGetAllUseCaseRes{
			RekapitulasiHasilKuesioner: res.RekapitulasiHasilKuesioner,
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
