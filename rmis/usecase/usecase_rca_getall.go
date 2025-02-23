package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type RcaGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type RcaGetAllUseCaseRes struct {
	Rca      []model.Rca       `json:"rca"`
	Metadata *usecase.Metadata `json:"metadata"`
}

type RcaGetAllUseCase = core.ActionHandler[RcaGetAllUseCaseReq, RcaGetAllUseCaseRes]

func ImplRcaGetAllUseCase(getAllRcas gateway.RcaGetAll) RcaGetAllUseCase {
	return func(ctx context.Context, req RcaGetAllUseCaseReq) (*RcaGetAllUseCaseRes, error) {

		res, err := getAllRcas(ctx, gateway.RcaGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &RcaGetAllUseCaseRes{
			Rca: res.Rca,
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
