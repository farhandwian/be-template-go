package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type AssetGetByIDUseCaseReq struct {
	ID string `json:"id"`
}

type AssetGetByIDUseCaseRes struct {
	Asset model.Asset `json:"asset"`
}

type AssetGetByIDUseCase = core.ActionHandler[AssetGetByIDUseCaseReq, AssetGetByIDUseCaseRes]

func ImplAssetGetByIDUseCase(getAssetByID gateway.AssetGetByID) AssetGetByIDUseCase {
	return func(ctx context.Context, req AssetGetByIDUseCaseReq) (*AssetGetByIDUseCaseRes, error) {
		res, err := getAssetByID(ctx, gateway.AssetGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}
		return &AssetGetByIDUseCaseRes{Asset: res.Asset}, nil
	}
}
