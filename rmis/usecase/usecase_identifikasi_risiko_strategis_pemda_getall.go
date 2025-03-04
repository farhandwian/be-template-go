package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type IdentifikasiRisikoStrategisPemdaGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
	Periode   string
}

// Change the response type to use the DTO with KategoriRisikoName
type IdentifikasiRisikoStrategisPemdaGetAllUseCaseRes struct {
	IdentifikasiRisikoStrategisPemda []model.IdentifikasiRisikoStrategisPemda `json:"identifikasi_risiko_strategis_pemda"`
	Metadata                         *usecase.Metadata                        `json:"metadata"`
}

type IdentifikasiRisikoStrategisPemdaGetAllUseCase = core.ActionHandler[IdentifikasiRisikoStrategisPemdaGetAllUseCaseReq, IdentifikasiRisikoStrategisPemdaGetAllUseCaseRes]

func ImplIdentifikasiRisikoStrategisPemdaGetAllUseCase(getAllIdentifikasiRisikoStrategisPemdas gateway.IdentifikasiRisikoStrategisPemdaGetAll) IdentifikasiRisikoStrategisPemdaGetAllUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisPemdaGetAllUseCaseReq) (*IdentifikasiRisikoStrategisPemdaGetAllUseCaseRes, error) {

		// Fetch the results from the gateway (which already includes the mapped KategoriRisikoName)
		res, err := getAllIdentifikasiRisikoStrategisPemdas(ctx, gateway.IdentifikasiRisikoStrategisPemdaGetAllReq{
			Page: req.Page, Size: req.Size, Keyword: req.Keyword, Status: req.Status, SortBy: req.SortBy, SortOrder: req.SortOrder, Periode: req.Periode,
		})
		if err != nil {
			return nil, err
		}

		// Pagination calculation
		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / req.Size

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
