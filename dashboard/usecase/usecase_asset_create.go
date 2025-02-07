package usecase

import (
	"context"
	"dashboard/gateway"
	"dashboard/model"
	"shared/core"
)

type AssetCreateUseCaseReq struct {
	Name     string `json:"name"`
	PIC      string `json:"pic"`
	Location string `json:"location"`
}

type AssetCreateUseCaseRes struct {
	ID string `json:"id"`
}

type AssetCreateUseCase = core.ActionHandler[AssetCreateUseCaseReq, AssetCreateUseCaseRes]

func ImplAssetCreateUseCase(
	generateId gateway.GenerateId,
	createAsset gateway.AssetSave,
) AssetCreateUseCase {
	return func(ctx context.Context, req AssetCreateUseCaseReq) (*AssetCreateUseCaseRes, error) {

		genObj, err := generateId(ctx, gateway.GenerateIdReq{})
		if err != nil {
			return nil, err
		}

		obj := model.Asset{
			ID:       genObj.RandomId,
			Name:     req.Name,
			PIC:      req.PIC,
			Location: req.Location,
			Status:   model.AssetStatusActive,
		}

		if _, err = createAsset(ctx, gateway.AssetSaveReq{Asset: obj}); err != nil {
			return nil, err
		}

		return &AssetCreateUseCaseRes{
			ID: genObj.RandomId,
		}, nil
	}
}
