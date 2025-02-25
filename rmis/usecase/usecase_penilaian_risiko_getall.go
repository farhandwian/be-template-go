package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PenilaianRisikoGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type PenilaianRisikoGetAllUseCaseRes struct {
	PenilaianRisiko []model.PenilaianRisiko `json:"hasil_analisis_risiko"`
	Metadata        *usecase.Metadata       `json:"metadata"`
}

type PenilaianRisikoGetAllUseCase = core.ActionHandler[PenilaianRisikoGetAllUseCaseReq, PenilaianRisikoGetAllUseCaseRes]

func ImplPenilaianRisikoGetAllUseCase(getAllPenilaianRisikos gateway.PenilaianRisikoGetAll) PenilaianRisikoGetAllUseCase {
	return func(ctx context.Context, req PenilaianRisikoGetAllUseCaseReq) (*PenilaianRisikoGetAllUseCaseRes, error) {

		res, err := getAllPenilaianRisikos(ctx, gateway.PenilaianRisikoGetAllReq{
			Page:      req.Page,
			Size:      req.Size,
			Keyword:   req.Keyword,
			SortBy:    req.SortBy,
			SortOrder: req.SortOrder,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &PenilaianRisikoGetAllUseCaseRes{
			PenilaianRisiko: res.PenilaianRisiko,
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
