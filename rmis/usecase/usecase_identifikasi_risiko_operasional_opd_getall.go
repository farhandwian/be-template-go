package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type IdentifikasiRisikoOperasionalOPDGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type IdentifikasiRisikoOperasionalOPDGetAllUseCaseRes struct {
	IdentifikasiRisikoOperasionalOPDs []model.IdentifikasiRisikoOperasionalOPDGetRes `json:"identifkasi_risiko_operasional_opds"`
	Metadata                          *usecase.Metadata                              `json:"metadata"`
}

type IdentifikasiRisikoOperasionalOPDGetAllUseCase = core.ActionHandler[IdentifikasiRisikoOperasionalOPDGetAllUseCaseReq, IdentifikasiRisikoOperasionalOPDGetAllUseCaseRes]

func ImplIdentifikasiRisikoOperasionalOPDGetAllUseCase(getAllIdentifikasiRisikoOperasionalOPDs gateway.IdentifikasiRisikoOperasionalOPDGetAll, getOneOPD gateway.OPDGetByID) IdentifikasiRisikoOperasionalOPDGetAllUseCase {
	return func(ctx context.Context, req IdentifikasiRisikoOperasionalOPDGetAllUseCaseReq) (*IdentifikasiRisikoOperasionalOPDGetAllUseCaseRes, error) {

		identifikasiRisikoOperasionalOPDs, err := getAllIdentifikasiRisikoOperasionalOPDs(ctx, gateway.IdentifikasiRisikoOperasionalOPDGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(identifikasiRisikoOperasionalOPDs.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		identifikasiRisikoOperasionalOPDsRes := make([]model.IdentifikasiRisikoOperasionalOPDGetRes, len(identifikasiRisikoOperasionalOPDs.IdentifikasiRisikoOperasionalOPD))

		for i, identifikasiRisikoOperasionalOPD := range identifikasiRisikoOperasionalOPDs.IdentifikasiRisikoOperasionalOPD {
			opd, err := getOneOPD(ctx, gateway.OPDGetByIDReq{ID: *identifikasiRisikoOperasionalOPD.OPDID})
			if err != nil {
				return nil, err
			}
			identifikasiRisikoOperasionalOPDsRes[i] = model.IdentifikasiRisikoOperasionalOPDGetRes{
				IdentifikasiRisikoOperasionalOPD: identifikasiRisikoOperasionalOPD,
				OPD:                              opd.OPD,
			}
		}
		return &IdentifikasiRisikoOperasionalOPDGetAllUseCaseRes{
			IdentifikasiRisikoOperasionalOPDs: identifikasiRisikoOperasionalOPDsRes,
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
