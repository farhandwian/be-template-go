package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type IdentifikasiRisikoStrategisPemdaGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type IdentifikasiRisikoStrategisPemdaGetAllUseCaseRes struct {
	IdentifikasiRisikoStrategisPemda []model.IdentifikasiRisikoStrategisPemerintahDaerah `json:"identifkasi_risiko_strategis_pemda"`
	Metadata                         *usecase.Metadata                                   `json:"metadata"`
}

type IdentifikasiRisikoStrategisPemdaGetAllUseCase = core.ActionHandler[IdentifikasiRisikoStrategisPemdaGetAllUseCaseReq, IdentifikasiRisikoStrategisPemdaGetAllUseCaseRes]

func ImplIdentifikasiRisikoStrategisPemdaGetAllUseCase(getAllIdentifikasiRisikoStrategisPemdas gateway.IdentifikasiRisikoStrategisPemdaGetAll) IdentifikasiRisikoStrategisPemdaGetAllUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaGetAllUseCaseReq) (*IdentifikasiRisikoStrategisPemdaGetAllUseCaseRes, error) {

		res, err := getAllIdentifikasiRisikoStrategisPemdas(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &IdentifikasiRisikoStrategisPemdaGetAllUseCaseRes{
			IdentifikasiRisikoStrategisPemda: res.IdentifikasiRisikoStrategisPemda,
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
