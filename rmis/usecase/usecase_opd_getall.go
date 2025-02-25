package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type OPDGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type OPDGetAllUseCaseRes struct {
	OPDs     []model.OPD       `json:"opds"`
	Metadata *usecase.Metadata `json:"metadata"`
}

type OPDGetAllUseCase = core.ActionHandler[OPDGetAllUseCaseReq, OPDGetAllUseCaseRes]

func ImplOPDGetAllUseCase(getAllOPDs gateway.OPDGetAll) OPDGetAllUseCase {
	return func(ctx context.Context, req OPDGetAllUseCaseReq) (*OPDGetAllUseCaseRes, error) {

		res, err := getAllOPDs(ctx, gateway.OPDGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &OPDGetAllUseCaseRes{
			OPDs: res.OPD,
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
