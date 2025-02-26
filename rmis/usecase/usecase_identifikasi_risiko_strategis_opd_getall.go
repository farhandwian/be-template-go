package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type IdentifikasiRisikoStrategisOPDGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type IdentifikasiRisikoStrategisOPDGetAllUseCaseRes struct {
	IdentifikasiRisikoStrategisOPD []model.IdentifikasiRisikoStrategisOPD `json:"identifkasi_risiko_strategis_opd"`
	Metadata                       *usecase.Metadata                      `json:"metadata"`
}

type IdentifikasiRisikoStrategisOPDGetAllUseCase = core.ActionHandler[IdentifikasiRisikoStrategisOPDGetAllUseCaseReq, IdentifikasiRisikoStrategisOPDGetAllUseCaseRes]

func ImplIdentifikasiRisikoStrategisOPDGetAllUseCase(getAllIdentifikasiRisikoStrategisOPDs gateway.IdentifikasiRisikoStrategisOPDGetAll) IdentifikasiRisikoStrategisOPDGetAllUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDGetAllUseCaseReq) (*IdentifikasiRisikoStrategisOPDGetAllUseCaseRes, error) {

		res, err := getAllIdentifikasiRisikoStrategisOPDs(ctx, gateway.IdentifikasiRisikoStrategisOPDGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &IdentifikasiRisikoStrategisOPDGetAllUseCaseRes{
			IdentifikasiRisikoStrategisOPD: res.IdentifikasiRisikoStrategisOPD,
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
