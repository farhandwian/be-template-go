package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type DaftarRisikoPrioritasGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type DaftarRisikoPrioritasGetAllUseCaseRes struct {
	DaftarRisikoPrioritas []model.DaftarRisikoPrioritas `json:"hasil_analisis_risiko"`
	Metadata              *usecase.Metadata             `json:"metadata"`
}

type DaftarRisikoPrioritasGetAllUseCase = core.ActionHandler[DaftarRisikoPrioritasGetAllUseCaseReq, DaftarRisikoPrioritasGetAllUseCaseRes]

func ImplDaftarRisikoPrioritasGetAllUseCase(getAllDaftarRisikoPrioritass gateway.DaftarRisikoPrioritasGetAll) DaftarRisikoPrioritasGetAllUseCase {
	return func(ctx context.Context, req DaftarRisikoPrioritasGetAllUseCaseReq) (*DaftarRisikoPrioritasGetAllUseCaseRes, error) {

		res, err := getAllDaftarRisikoPrioritass(ctx, gateway.DaftarRisikoPrioritasGetAllReq{
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

		return &DaftarRisikoPrioritasGetAllUseCaseRes{
			DaftarRisikoPrioritas: res.DaftarRisikoPrioritas,
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
