package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type IdentifikasiRisikoOperasionalOPDGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type IdentifikasiRisikoOperasionalOPDGetAllUseCaseRes struct {
	IdentifikasiRisikoOperasionalOPDs []model.IdentifikasiRisikoOperasionalOPDResponse `json:"identifkasi_risiko_operasional_opds"`
	Metadata                          *usecase.Metadata                                `json:"metadata"`
}

type IdentifikasiRisikoOperasionalOPDGetAllUseCase = core.ActionHandler[IdentifikasiRisikoOperasionalOPDGetAllUseCaseReq, IdentifikasiRisikoOperasionalOPDGetAllUseCaseRes]

func ImplIdentifikasiRisikoOperasionalOPDGetAllUseCase(getAllIdentifikasiRisikoOperasionalOPDs gateway.IdentifikasiRisikoOperasionalOPDGetAll) IdentifikasiRisikoOperasionalOPDGetAllUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDGetAllUseCaseReq) (*IdentifikasiRisikoOperasionalOPDGetAllUseCaseRes, error) {

		identifikasiRisikoOperasionalOPDs, err := getAllIdentifikasiRisikoOperasionalOPDs(ctx, gateway.IdentifikasiRisikoOperasionalOPDGetAllReq{
			Page: req.Page, Size: req.Size, Keyword: req.Keyword, SortBy: req.SortBy, SortOrder: req.SortOrder, Status: req.Status,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(identifikasiRisikoOperasionalOPDs.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &IdentifikasiRisikoOperasionalOPDGetAllUseCaseRes{
			IdentifikasiRisikoOperasionalOPDs: identifikasiRisikoOperasionalOPDs.IdentifikasiRisikoOperasionalOPD,
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
