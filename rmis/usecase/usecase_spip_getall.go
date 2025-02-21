package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type SpipGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type SpipGetAllUseCaseRes struct {
	Spip     []model.SPIP      `json:"spips"`
	Metadata *usecase.Metadata `json:"metadata"`
}

type SpipGetAllUseCase = core.ActionHandler[SpipGetAllUseCaseReq, SpipGetAllUseCaseRes]

func ImplSpipGetAllUseCase(getAllSpips gateway.SpipGetAll) SpipGetAllUseCase {
	return func(ctx context.Context, req SpipGetAllUseCaseReq) (*SpipGetAllUseCaseRes, error) {

		res, err := getAllSpips(ctx, gateway.SpipGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &SpipGetAllUseCaseRes{
			Spip: res.SPIP,
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
