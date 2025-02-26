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
	IdentifikasiRisikoStrategisOPDs []model.IdentifikasiRisikoStrategisOPDGetRes `json:"identifkasi_risiko_strategis_opds"`
	Metadata                        *usecase.Metadata                            `json:"metadata"`
}

type IdentifikasiRisikoStrategisOPDGetAllUseCase = core.ActionHandler[IdentifikasiRisikoStrategisOPDGetAllUseCaseReq, IdentifikasiRisikoStrategisOPDGetAllUseCaseRes]

func ImplIdentifikasiRisikoStrategisOPDGetAllUseCase(getAllIdentifikasiRisikoStrategisOPDs gateway.IdentifikasiRisikoStrategisOPDGetAll, getOneOPD gateway.OPDGetByID) IdentifikasiRisikoStrategisOPDGetAllUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoStrategisOPDGetAllUseCaseReq) (*IdentifikasiRisikoStrategisOPDGetAllUseCaseRes, error) {

		identifikasiRisikoStrategisOPDs, err := getAllIdentifikasiRisikoStrategisOPDs(ctx, gateway.IdentifikasiRisikoStrategisOPDGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(identifikasiRisikoStrategisOPDs.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		identifikasiRisikoStrategisOPDsRes := make([]model.IdentifikasiRisikoStrategisOPDGetRes, len(identifikasiRisikoStrategisOPDs.IdentifikasiRisikoStrategisOPD))

		for i, identifikasiRisikoStrategisOPD := range identifikasiRisikoStrategisOPDs.IdentifikasiRisikoStrategisOPD {
			opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *identifikasiRisikoStrategisOPD.OPDID})
			if err != nil {
				return nil, err
			}
			identifikasiRisikoStrategisOPDsRes[i] = model.IdentifikasiRisikoStrategisOPDGetRes{
				IdentifikasiRisikoStrategisOPD: identifikasiRisikoStrategisOPD,
				OPD:                            opd.OPD,
			}
		}
		return &IdentifikasiRisikoStrategisOPDGetAllUseCaseRes{
			IdentifikasiRisikoStrategisOPDs: identifikasiRisikoStrategisOPDsRes,
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
