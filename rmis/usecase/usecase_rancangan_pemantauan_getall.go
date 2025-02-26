package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type RancanganPemantauanGetAllUseCaseReq struct {
	Keyword   string
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

type RancanganPemantauanGetAllUseCaseRes struct {
	RancanganPemantauan []model.RancanganPemantauan `json:"rancangan_pemantauan"`
	Metadata            *usecase.Metadata           `json:"metadata"`
}

type RancanganPemantauanGetAllUseCase = core.ActionHandler[RancanganPemantauanGetAllUseCaseReq, RancanganPemantauanGetAllUseCaseRes]

func ImplRancanganPemantauanGetAllUseCase(getAllRancanganPemantauans gateway.RancanganPemantauanGetAll) RancanganPemantauanGetAllUseCase {
	return func(ctx context.Context, req RancanganPemantauanGetAllUseCaseReq) (*RancanganPemantauanGetAllUseCaseRes, error) {

		res, err := getAllRancanganPemantauans(ctx, gateway.RancanganPemantauanGetAllReq{
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

		return &RancanganPemantauanGetAllUseCaseRes{
			RancanganPemantauan: res.RancanganPemantauan,
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
