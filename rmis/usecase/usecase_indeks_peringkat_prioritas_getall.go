package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type IndeksPeringkatPrioritasGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type IndeksPeringkatPrioritasGetAllUseCaseRes struct {
	IndeksPeringkatPrioritass []model.IndeksPeringkatPrioritas `json:"indeks_peringkat_prioritas"`
	Metadata                  *usecase.Metadata                `json:"metadata"`
}

type IndeksPeringkatPrioritasGetAllUseCase = core.ActionHandler[IndeksPeringkatPrioritasGetAllUseCaseReq, IndeksPeringkatPrioritasGetAllUseCaseRes]

func ImplIndeksPeringkatPrioritasGetAllUseCase(getAllIndeksPeringkatPrioritass gateway.IndeksPeringkatPrioritasGetAll) IndeksPeringkatPrioritasGetAllUseCase {
	return func(ctx context.Context, req IndeksPeringkatPrioritasGetAllUseCaseReq) (*IndeksPeringkatPrioritasGetAllUseCaseRes, error) {

		res, err := getAllIndeksPeringkatPrioritass(ctx, gateway.IndeksPeringkatPrioritasGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword, SortBy: req.SortBy, SortOrder: req.SortOrder})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &IndeksPeringkatPrioritasGetAllUseCaseRes{
			IndeksPeringkatPrioritass: res.IndeksPeringkatPrioritas,
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
