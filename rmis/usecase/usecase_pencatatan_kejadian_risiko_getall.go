package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type PencatatanKejadianRisikoGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type PencatatanKejadianRisikoGetAllUseCaseRes struct {
	PencatatanKejadianRisikos []model.PencatatanKejadianRisiko `json:"pencatatan_kejadian_risiko"`
	Metadata                  *usecase.Metadata                `json:"metadata"`
}

type PencatatanKejadianRisikoGetAllUseCase = core.ActionHandler[PencatatanKejadianRisikoGetAllUseCaseReq, PencatatanKejadianRisikoGetAllUseCaseRes]

func ImplPencatatanKejadianRisikoGetAllUseCase(getAllPencatatanKejadianRisikos gateway.PencatatanKejadianRisikoGetAll) PencatatanKejadianRisikoGetAllUseCase {
	return func(ctx context.Context, req PencatatanKejadianRisikoGetAllUseCaseReq) (*PencatatanKejadianRisikoGetAllUseCaseRes, error) {

		res, err := getAllPencatatanKejadianRisikos(ctx, gateway.PencatatanKejadianRisikoGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword, SortBy: req.SortBy, SortOrder: req.SortOrder})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &PencatatanKejadianRisikoGetAllUseCaseRes{
			PencatatanKejadianRisikos: res.PencatatanKejadianRisiko,
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
