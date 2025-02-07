package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
	"shared/usecase"
)

type GetListJDIHReq struct {
	Keyword string
	Page    int
	Size    int
}

type GetListJDIHResp struct {
	ListJDIH []model.JDIH      `json:"jdih"`
	Metadata *usecase.Metadata `json:"metadata"`
}

type GetListDocumentAndLawUseCase = core.ActionHandler[GetListJDIHReq, GetListJDIHResp]

func ImplGetListDocumentAndLawUseCase(getListJDIH gateway.GetListJDIHGateway) GetListDocumentAndLawUseCase {
	return func(ctx context.Context, req GetListJDIHReq) (*GetListJDIHResp, error) {

		res, err := getListJDIH(ctx, gateway.GetListJDIHReq{Keyword: req.Keyword, Page: req.Page, Size: req.Size})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &GetListJDIHResp{
			ListJDIH: res.JDIH,
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
