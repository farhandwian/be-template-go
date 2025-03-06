package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type IdentifikasiRisikoStrategisOPDGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
	Status    string
}

type IdentifikasiRisikoStrategisOPDGetAllUseCaseRes struct {
	IdentifikasiRisikoStrategisOPDs []model.IdentifikasiRisikoStrategisOPDResponse `json:"identifkasi_risiko_strategis_opds"`
	Metadata                        *usecase.Metadata                              `json:"metadata"`
}

type IdentifikasiRisikoStrategisOPDGetAllUseCase = core.ActionHandler[IdentifikasiRisikoStrategisOPDGetAllUseCaseReq, IdentifikasiRisikoStrategisOPDGetAllUseCaseRes]

func ImplIdentifikasiRisikoStrategisOPDGetAllUseCase(getAllIdentifikasiRisikoStrategisOPDs gateway.IdentifikasiRisikoStrategisOPDGetAll) IdentifikasiRisikoStrategisOPDGetAllUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDGetAllUseCaseReq) (*IdentifikasiRisikoStrategisOPDGetAllUseCaseRes, error) {

		identifikasiRisikoStrategisOPDs, err := getAllIdentifikasiRisikoStrategisOPDs(ctx, gateway.IdentifikasiRisikoStrategisOPDGetAllReq{
			Page: req.Page, Size: req.Size, Keyword: req.Keyword, SortBy: req.SortBy, SortOrder: req.SortOrder, Status: req.Status,
		})
		if err != nil {
			return nil, err
		}

		totalItems := int(identifikasiRisikoStrategisOPDs.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)
		return &IdentifikasiRisikoStrategisOPDGetAllUseCaseRes{
			IdentifikasiRisikoStrategisOPDs: identifikasiRisikoStrategisOPDs.IdentifikasiRisikoStrategisOPD,
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
