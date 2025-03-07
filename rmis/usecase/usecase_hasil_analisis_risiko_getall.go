package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type HasilAnalisisRisikoGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type HasilAnalisisRisikoGetAllUseCaseRes struct {
	HasilAnalisisRisiko []model.HasilAnalisisRisikoResponse `json:"hasil_analisis_risiko"`
	Metadata            *usecase.Metadata                   `json:"metadata"`
}

type HasilAnalisisRisikoGetAllUseCase = core.ActionHandler[HasilAnalisisRisikoGetAllUseCaseReq, HasilAnalisisRisikoGetAllUseCaseRes]

func ImplHasilAnalisisRisikoGetAllUseCase(getAllHasilAnalisisRisikos gateway.HasilAnalisisRisikoGetAll) HasilAnalisisRisikoGetAllUseCase {
	return func(ctx context.Context, req HasilAnalisisRisikoGetAllUseCaseReq) (*HasilAnalisisRisikoGetAllUseCaseRes, error) {

		res, err := getAllHasilAnalisisRisikos(ctx, gateway.HasilAnalisisRisikoGetAllReq{
			Page: req.Page, Size: req.Size, Keyword: req.Keyword, SortBy: req.SortBy, SortOrder: req.SortOrder, Status: req.Status,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &HasilAnalisisRisikoGetAllUseCaseRes{
			HasilAnalisisRisiko: res.HasilAnalisisRisiko,
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
