package usecase

import (
	"context"
	"dashboard/gateway"
	"shared/core"
)

type AssetDeleteUseCaseReq struct {
	ID string `json:"id"`
}

type AssetDeleteUseCaseRes struct{}

type AssetDeleteUseCase = core.ActionHandler[AssetDeleteUseCaseReq, AssetDeleteUseCaseRes]

func ImplAssetDeleteUseCase(deleteAsset gateway.AssetDelete) AssetDeleteUseCase {
	return func(ctx context.Context, req AssetDeleteUseCaseReq) (*AssetDeleteUseCaseRes, error) {

		if _, err := deleteAsset(ctx, gateway.AssetDeleteReq{ID: req.ID}); err != nil {
			return nil, err
		}

		return &AssetDeleteUseCaseRes{}, nil
	}
}
