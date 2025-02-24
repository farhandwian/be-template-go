package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type IKUGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type IKUGetAllUseCaseRes struct {
	IKU      []model.IKU       `json:"ikus"`
	Metadata *usecase.Metadata `json:"metadata"`
}

type IKUGetAllUseCase = core.ActionHandler[IKUGetAllUseCaseReq, IKUGetAllUseCaseRes]

func ImplIKUGetAllUseCase(getAllIKUs gateway.IKUGetAll) IKUGetAllUseCase {
	return func(ctx context.Context, req IKUGetAllUseCaseReq) (*IKUGetAllUseCaseRes, error) {

		res, err := getAllIKUs(ctx, gateway.IKUGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &IKUGetAllUseCaseRes{
			IKU: res.IKU,
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
