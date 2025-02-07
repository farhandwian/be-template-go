package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type AssetUpdateUseCaseReq struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	PIC      string            `json:"pic"`
	Location string            `json:"location"`
	Status   model.AssetStatus `json:"status"`
}

type AssetUpdateUseCaseRes struct{}

type AssetUpdateUseCase = core.ActionHandler[AssetUpdateUseCaseReq, AssetUpdateUseCaseRes]

func ImplAssetUpdateUseCase(
	getAssetById gateway.AssetGetByID,
	updateAsset gateway.AssetSave,
) AssetUpdateUseCase {
	return func(ctx context.Context, req AssetUpdateUseCaseReq) (*AssetUpdateUseCaseRes, error) {

		res, err := getAssetById(ctx, gateway.AssetGetByIDReq{ID: req.ID})
		if err != nil {
			return nil, err
		}

		res.Asset.Location = req.Location
		res.Asset.Name = req.Name
		res.Asset.PIC = req.PIC
		res.Asset.Status = req.Status

		if _, err := updateAsset(ctx, gateway.AssetSaveReq{Asset: res.Asset}); err != nil {
			return nil, err
		}

		return &AssetUpdateUseCaseRes{}, nil
	}
}
