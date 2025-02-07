package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
	"shared/usecase"
)

type AssetGetAllUseCaseReq struct {
	Keyword string
	Page    int
	Size    int
}

type AssetGetAllUseCaseRes struct {
	Assets   []model.Asset     `json:"assets"`
	Metadata *usecase.Metadata `json:"metadata"`
}

type AssetGetAllUseCase = core.ActionHandler[AssetGetAllUseCaseReq, AssetGetAllUseCaseRes]

func ImplAssetGetAllUseCase(getAllAssets gateway.AssetGetAll) AssetGetAllUseCase {
	return func(ctx context.Context, req AssetGetAllUseCaseReq) (*AssetGetAllUseCaseRes, error) {

		res, err := getAllAssets(ctx, gateway.AssetGetAllReq{Page: req.Page, Size: req.Size, Keyword: req.Keyword})
		if err != nil {
			return nil, err
		}

		totalItems := int(res.Count)
		totalPages := (totalItems + req.Size - 1) / (req.Size)

		return &AssetGetAllUseCaseRes{
			Assets: res.Assets,
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
