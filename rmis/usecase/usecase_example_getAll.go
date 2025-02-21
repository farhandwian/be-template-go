package usecase

import (
	"context"
	"rmis/gateway"
	"rmis/model"
	"shared/core"
	"shared/usecase"
)

type ExampleGetAllUseReq struct {
	TestString  string
	TestNumber  int
	TestBoolean bool
	SortOrder   string
	SortBy      string
	Page        int `json:"page"`
	Size        int `json:"size"`
}

type ExampleGetAllUseRes struct {
	Examples []model.Example   `json:"examples"`
	Metadata *usecase.Metadata `json:"metadata"`
}

type ExampleGetAllUseCase = core.ActionHandler[ExampleGetAllUseReq, ExampleGetAllUseRes]

func ImplExampleGetAllUseCase(
	exampleGateway gateway.ExampleGateway,
) ExampleGetAllUseCase {
	return func(ctx context.Context, req ExampleGetAllUseReq) (*ExampleGetAllUseRes, error) {
		res, err := exampleGateway(ctx, gateway.ExampleGetAllGatewayReq{Keyword: req.TestString, Page: req.Page, Size: req.Size})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)
		return &ExampleGetAllUseRes{
			Examples: res.Employees,
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
