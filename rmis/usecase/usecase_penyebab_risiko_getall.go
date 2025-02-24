package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PenyebabRisikoGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type PenyebabRisikoGetAllUseCaseRes struct {
	PenyebabRisiko []model.PenyebabRisiko `json:"PenyebabRisiko"`
	Metadata       *usecase.Metadata      `json:"metadata"`
}

type PenyebabRisikoGetAllUseCase = core.ActionHandler[PenyebabRisikoGetAllUseCaseReq, PenyebabRisikoGetAllUseCaseRes]

func ImplPenyebabRisikoGetAllUseCase(getAllPenyebabRisikos gateway.PenyebabRisikoGetAll) PenyebabRisikoGetAllUseCase {
	return func(ctx context.Context, req PenyebabRisikoGetAllUseCaseReq) (*PenyebabRisikoGetAllUseCaseRes, error) {

		res, err := getAllPenyebabRisikos(ctx, gateway.PenyebabRisikoGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &PenyebabRisikoGetAllUseCaseRes{
			PenyebabRisiko: res.PenyebabRisiko,
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
