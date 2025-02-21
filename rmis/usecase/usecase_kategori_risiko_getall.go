package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type KategoriRisikoGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type KategoriRisikoGetAllUseCaseRes struct {
	KategoriRisiko []model.KategoriRisiko `json:"KategoriRisikos"`
	Metadata       *usecase.Metadata      `json:"metadata"`
}

type KategoriRisikoGetAllUseCase = core.ActionHandler[KategoriRisikoGetAllUseCaseReq, KategoriRisikoGetAllUseCaseRes]

func ImplKategoriRisikoGetAllUseCase(getAllKategoriRisikos gateway.KategoriRisikoGetAll) KategoriRisikoGetAllUseCase {
	return func(ctx context.Context, req KategoriRisikoGetAllUseCaseReq) (*KategoriRisikoGetAllUseCaseRes, error) {

		res, err := getAllKategoriRisikos(ctx, gateway.KategoriRisikoGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &KategoriRisikoGetAllUseCaseRes{
			KategoriRisiko: res.KategoriRisiko,
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
