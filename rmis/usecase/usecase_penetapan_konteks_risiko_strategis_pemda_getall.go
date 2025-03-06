package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PenetapanKonteksRisikoGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type PenetapanKonteksRisikoGetAllUseCaseRes struct {
	PenetapanKonteksRisikos []model.PenetapanKonteksRisikoStrategisPemda `json:"penetapan_konteks_risikos"`
	Metadata                *usecase.Metadata                            `json:"metadata"`
}

type PenetapanKonteksRisikoGetAllUseCase = core.ActionHandler[PenetapanKonteksRisikoGetAllUseCaseReq, PenetapanKonteksRisikoGetAllUseCaseRes]

func ImplPenetapanKonteksRisikoGetAllUseCase(getAllPenetapanKonteksRisikos gateway.PenetapanKonteksRisikoStrategisPemdaGetAll, getAllIKUs gateway.IKUGetAll) PenetapanKonteksRisikoGetAllUseCase {
	return func(ctx context.Context, req PenetapanKonteksRisikoGetAllUseCaseReq) (*PenetapanKonteksRisikoGetAllUseCaseRes, error) {

		res, err := getAllPenetapanKonteksRisikos(ctx, gateway.PenetapanKonteksRisikoStrategisPemdaGetAllReq{
			Page: req.Page, Size: req.Size, Keyword: req.Keyword, SortBy: req.SortBy, SortOrder: req.SortOrder, Status: req.Status,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &PenetapanKonteksRisikoGetAllUseCaseRes{
			PenetapanKonteksRisikos: res.PenetapanKonteksRisikoStrategisPemda,
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
