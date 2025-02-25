package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type KriteriaDampakGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type KriteriaDampakGetAllUseCaseRes struct {
	KriteriaDampak []model.KriteriaDampak `json:"kriteria_dampaks"`
	Metadata       *usecase.Metadata      `json:"metadata"`
}

type KriteriaDampakGetAllUseCase = core.ActionHandler[KriteriaDampakGetAllUseCaseReq, KriteriaDampakGetAllUseCaseRes]

func ImplKriteriaDampakGetAllUseCase(getAllKriteriaDampaks gateway.KriteriaDampakGetAll) KriteriaDampakGetAllUseCase {
	return func(ctx context.Context, req KriteriaDampakGetAllUseCaseReq) (*KriteriaDampakGetAllUseCaseRes, error) {

		res, err := getAllKriteriaDampaks(ctx, gateway.KriteriaDampakGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &KriteriaDampakGetAllUseCaseRes{
			KriteriaDampak: res.KriteriaDampak,
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
